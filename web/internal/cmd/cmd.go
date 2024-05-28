package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/os/gcfg"
	"web/internal/controller"
	"web/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
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
