package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/logging"
	"google.golang.org/grpc"
)

// GrpcRegFunc 注册服务接口方法
type GrpcRegFunc func(server *grpcx.GrpcServer)

// GetGrpcCmdFunc 标准化grpc服务方法
//
//	示例:
//	gcmd.Command{
//		Name:  "main",
//		Usage: "./main",
//		Brief: "主命令",
//		Func: cmd.GetGrpcCmdFunc(cmd.GetGrpcMiddleware(), func(server *grpcx.GrpcServer){
//			demoController.Register(server)
//		}),
//	}
func GetGrpcCmdFunc(middleware MiddlewareForCmd, grpcRegFunc GrpcRegFunc) func(ctx context.Context, parser *gcmd.Parser) error {
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		shutdown := middleware(ctx)
		defer shutdown()

		c := grpcx.Server.NewConfig()
		c.Options = append(c.Options, []grpc.ServerOption{
			grpcx.Server.ChainUnary(
				logging.UnarySLogger,
				grpcx.Server.UnaryValidate,
			)}...,
		)
		s := grpcx.Server.New(c)
		//注册服务接口
		grpcRegFunc(s)
		s.Run()
		return
	}
}
