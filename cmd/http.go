package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/logging"
	"github.com/yiqiang3344/gf-micro/response"
)

type ServerGroup struct {
	Prefix      string //前缀，默认"/"
	Middlewares []ghttp.HandlerFunc
	Controllers []interface{}
}

// GetHttpCmdFunc 标准化http服务方法
//
//	示例:
//	gcmd.Command{
//		Name:  "main",
//		Usage: "main",
//		Brief: "start http server",
//		Func: cmd.GetHttpCmdFunc(
//			cmd.GetHttpMiddleware(),
//			[]ghttp.HandlerFunc{},
//			[]cmd.ServerGroup{
//				{
//					Prefix: "/",
//					Middlewares: []ghttp.HandlerFunc{},
//					Controllers: []interface{}{
//						controller.Demo,
//					},
//				},
//			},
//		),
//	}
func GetHttpCmdFunc(middleware MiddlewareForCmd, useMiddlewares []ghttp.HandlerFunc, serverGroups []ServerGroup) func(ctx context.Context, parser *gcmd.Parser) error {
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		shutdown := middleware(ctx)
		defer shutdown()

		s := g.Server()
		middlewares := append([]ghttp.HandlerFunc{
			logging.MiddlewareLogFormatJson,
			logging.MiddlewareHandlerAccessLog,
			logging.MiddlewareHandlerErrorLog,
			response.MiddlewareHandlerResponse,
		}, useMiddlewares...)
		s.Use(middlewares...)
		for _, v := range serverGroups {
			s.Group(v.Prefix, func(group *ghttp.RouterGroup) {
				group.Middleware(v.Middlewares...)
				group.Bind(v.Controllers...)
			})
		}
		s.Run()
		return
	}
}
