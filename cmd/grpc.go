package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/flowColor"
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
				flowColor.GrpcServerUnary,
				logging.GrpcServerLoggerUnary,
				grpcx.Server.UnaryValidate,
			)}...,
		)
		s := grpcx.Server.New(c)
		if flowColor.IsOpen() {
			//设置服务的流量染色标识元数据，会跟随服务注册到注册中心去
			s.Service(&gsvc.LocalService{
				Name: c.Name,
				Metadata: gsvc.Metadata{
					flowColor.FlowColor: *flowColor.GetLocalFlowColor(),
				},
			})
		}
		//注册服务接口
		grpcRegFunc(s)
		s.Run()
		return
	}
}
