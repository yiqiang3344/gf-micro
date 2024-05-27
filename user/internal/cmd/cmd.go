package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"

	"yijunqiang/gf-micro/user/internal/controller/user"
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
				)}...,
			)
			s := grpcx.Server.New(c)
			user.Register(s)
			s.Run()
			return nil
		},
	}
)
