package logging

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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
	AccessLog{
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
	ErrorLog{
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
