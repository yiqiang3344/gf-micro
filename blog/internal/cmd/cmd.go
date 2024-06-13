package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"google.golang.org/grpc"
	"yijunqiang/gf-micro/blog/internal/logging"

	"yijunqiang/gf-micro/blog/internal/controller/blog"
)

var (
	Main = gcmd.Command{
		Name:        "main",
		Usage:       "main",
		Brief:       "博客服务管理工具",
		Description: `功能包括：微服务启动，消费者启动等`,
	}
	microServer = &gcmd.Command{
		Name:  "start",
		Usage: "./main start",
		Brief: "启动微服务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			shutdown := initMiddleware(ctx)
			defer shutdown()

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
					logging.UnaryLogger,
				)}...,
			)
			s := grpcx.Server.New(c)
			blog.Register(s)
			s.Run()
			return nil
		},
	}
)

func initMiddleware(ctx context.Context) func() {
	var shutdownArr []func()

	//配置中心
	if !gcfg.Instance().MustGet(ctx, "apollo").IsNil() {
		adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
		if err != nil {
			panic(err)
		}
		gcfg.Instance().SetAdapter(adapter)
	}

	//日志json化
	glog.SetDefaultHandler(logging.HandlerJson)

	// 链路追踪初始化
	if !gcfg.Instance().MustGet(ctx, "otlp").IsNil() {
		shutdown, err := otlpgrpc.Init(
			gcfg.Instance().MustGet(ctx, "appName").String(),
			gcfg.Instance().MustGet(ctx, "otlp.endpoint").String(),
			gcfg.Instance().MustGet(ctx, "otlp.traceToken").String(),
		)
		if err != nil {
			g.Log().Debugf(ctx, "otlp初始化失败:%v\n", err)
		} else {
			shutdownArr = append(shutdownArr, shutdown)
		}
	}

	// grpc服务注册发现
	grpcx.Resolver.Register(etcd.New(gcfg.Instance().MustGet(ctx, "registry.etcd").String()))

	return func() {
		for _, v := range shutdownArr {
			v()
		}
		return
	}
}

func init() {
	err := Main.AddCommand(microServer, batDeleteBlogConsumer)
	if err != nil {
		panic(err)
	}
}
