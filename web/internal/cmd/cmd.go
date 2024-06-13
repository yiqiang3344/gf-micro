package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"web/internal/controller"
	"web/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			shutdown := initMiddleware(ctx)
			defer shutdown()

			s := g.Server()
			s.Use(
				middleware.MiddlewareHandlerJson,
				middleware.MiddlewareHandlerAccessLog,
				middleware.MiddlewareHandlerErrorLog,
				middleware.MiddlewareHandlerResponse,
			)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.MiddlewareAuth,
				)
				group.Bind(
					controller.User,
					controller.Blog,
				)
			})
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
