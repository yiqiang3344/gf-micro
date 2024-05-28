package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"yijunqiang/gf-micro/user/internal/controller/user"
	"yijunqiang/gf-micro/user/internal/logging"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start user micro server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 链路追踪初始化
			shutdown, err := otlpgrpc.Init(
				gcfg.Instance().MustGet(ctx, "appName").String(),
				gcfg.Instance().MustGet(ctx, "otlp.endpoint").String(),
				gcfg.Instance().MustGet(ctx, "otlp.traceToken").String(),
			)
			if err != nil {
				g.Log().Fatal(ctx, err)
			}
			defer shutdown()

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
					UnaryLogger, //自定义请求及异常日志
				)}...,
			)
			s := grpcx.Server.New(c)
			user.Register(s)
			s.Run()
			return nil
		},
	}
)

// UnaryLogger is the default unary interceptor for logging purpose.
func UnaryLogger(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	var (
		start    = time.Now()
		res, err = handler(ctx, req)
		duration = time.Since(start)
	)
	handleAccessLog(ctx, err, duration, info, req, res)
	handleErrorLog(ctx, err, duration, info, req, res)
	return res, err
}

// handleAccessLog handles the access logging for server.
func handleAccessLog(
	ctx context.Context, err error, duration time.Duration, info *grpc.UnaryServerInfo, req, res interface{},
) {
	logging.AccessLog{
		Method: info.FullMethod,
		Cost:   fmt.Sprintf("%.3fms", float64(duration)/1e6),
		Req:    req,
		Res:    res,
	}.Log(ctx)
}

// handleErrorLog handles the error logging for server.
func handleErrorLog(
	ctx context.Context, err error, duration time.Duration, info *grpc.UnaryServerInfo, req, res interface{},
) {
	if err == nil {
		return
	}
	var (
		code          = gerror.Code(err)
		codeDetail    = code.Detail()
		codeDetailStr string
		grpcCode      codes.Code
		grpcMessage   string
	)
	if grpcStatus, ok := status.FromError(err); ok {
		grpcCode = grpcStatus.Code()
		grpcMessage = grpcStatus.Message()
	}
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	stackStr := ""
	if stack := gerror.Stack(err); stack != "" {
		stackStr += stack
	} else {
		stackStr += ", " + err.Error()
	}
	logging.ErrorLog{
		Method:      info.FullMethod,
		Cost:        fmt.Sprintf("%.3fms", float64(duration)/1e6),
		ErrorCode:   grpcCode,
		ErrorMsg:    grpcMessage,
		Req:         req,
		Res:         res,
		Code:        code.Code(),
		Message:     code.Message(),
		ErrorDetail: codeDetailStr,
		Stack:       stackStr,
	}.Log(ctx)
}
