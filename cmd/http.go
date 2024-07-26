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
			logging.HttpLogFormatJsonMiddleware,
			logging.HttpAccessLogMiddleware,
			logging.HttpErrorLogMiddleware,
			response.HttpResponseMiddleware,
		}, useMiddlewares...)
		s.Use(middlewares...)
		for _, v := range serverGroups {
			s.Group(v.Prefix, func(group *ghttp.RouterGroup) {
				group.Middleware(v.Middlewares...)
				//绑定健康检查
				group.GET("/health", func(r *ghttp.Request) {
					var err error
					defer func() {
						// 异常则响应500,及异常内容
						if err != nil {
							r.Response.Status = 500
							r.Response.Write(err.Error())
						}
					}()

					//如果有数据库配置，则检查数据库是否能正常连接
					if !g.Config().MustGet(r.GetCtx(), "database.default").IsNil() {
						_, err = g.DB().Ctx(r.GetCtx()).Raw("show tables").All()
						if err != nil {
							return
						}
					}

					//如果有redis配置，则检查redis是否能正常连接
					if !g.Config().MustGet(r.GetCtx(), "redis.default").IsNil() {
						_, err = g.Redis().Get(r.GetCtx(), "health")
						if err != nil {
							return
						}
					}

					//正常响应200
					r.Response.Status = 200
					r.Response.Write("success")
				})
				group.Bind(v.Controllers...)
			})
		}
		s.Run()
		return
	}
}
