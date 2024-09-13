package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/logging"
	"github.com/yiqiang3344/gf-micro/response"
)

type ServerGroup struct {
	Prefix      string //前缀，默认"/"
	Middlewares []ghttp.HandlerFunc
	Controllers []interface{}
}

type HttpCmdOption struct {
	LogFormatJsonMiddleware func(r *ghttp.Request)
	AccessLogMiddleware     func(r *ghttp.Request)
	ErrorLogMiddleware      func(r *ghttp.Request)
	ResponseMiddleware      func(r *ghttp.Request)
	BindHttpHealthApi       func(group *ghttp.RouterGroup)
}

type HttpCmdOptionApply func(option *HttpCmdOption)

func WithHttpCmdLogFormatJsonMiddleware(v func(r *ghttp.Request)) HttpCmdOptionApply {
	return func(option *HttpCmdOption) {
		option.LogFormatJsonMiddleware = v
	}
}
func WithHttpCmdAccessLogMiddleware(v func(r *ghttp.Request)) HttpCmdOptionApply {
	return func(option *HttpCmdOption) {
		option.AccessLogMiddleware = v
	}
}
func WithHttpCmdErrorLogMiddleware(v func(r *ghttp.Request)) HttpCmdOptionApply {
	return func(option *HttpCmdOption) {
		option.ErrorLogMiddleware = v
	}
}
func WithHttpCmdResponseMiddleware(v func(r *ghttp.Request)) HttpCmdOptionApply {
	return func(option *HttpCmdOption) {
		option.ResponseMiddleware = v
	}
}
func WithHttpCmdBindHttpHealthApi(v func(group *ghttp.RouterGroup)) HttpCmdOptionApply {
	return func(option *HttpCmdOption) {
		option.BindHttpHealthApi = v
	}
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
func GetHttpCmdFunc(middleware MiddlewareForCmd, useMiddlewares []ghttp.HandlerFunc, serverGroups []ServerGroup, optionApplyFuncs ...HttpCmdOptionApply) func(ctx context.Context, parser *gcmd.Parser) error {
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		shutdown := middleware(ctx, parser)
		defer shutdown()

		s := g.Server(g.Cfg().MustGet(ctx, cfg.APPNAME).String())
		s.SetAddr(GetCommonArguments(ctx, parser, Port).String())
		option := &HttpCmdOption{
			LogFormatJsonMiddleware: logging.HttpLogFormatJsonMiddleware,
			AccessLogMiddleware:     logging.HttpAccessLogMiddleware,
			ErrorLogMiddleware:      logging.HttpErrorLogMiddleware,
			ResponseMiddleware:      response.HttpResponseMiddleware,
			BindHttpHealthApi:       BindHttpHealthApi,
		}
		for _, optionApply := range optionApplyFuncs {
			optionApply(option)
		}
		var middlewares []ghttp.HandlerFunc
		if option.LogFormatJsonMiddleware != nil {
			middlewares = append(middlewares, option.LogFormatJsonMiddleware)
		}
		if option.AccessLogMiddleware != nil {
			middlewares = append(middlewares, option.AccessLogMiddleware)
		}
		if option.ErrorLogMiddleware != nil {
			middlewares = append(middlewares, option.ErrorLogMiddleware)
		}
		if option.ResponseMiddleware != nil {
			middlewares = append(middlewares, option.ResponseMiddleware)
		}
		middlewares = append(middlewares, useMiddlewares...)
		s.Use(middlewares...)
		for _, v := range serverGroups {
			s.Group(v.Prefix, func(group *ghttp.RouterGroup) {
				group.Middleware(v.Middlewares...)
				if option.BindHttpHealthApi != nil {
					//绑定健康检查接口
					option.BindHttpHealthApi(group)
				}
				group.Bind(v.Controllers...)
			})
		}
		s.Run()
		return
	}
}

// BindHttpHealthApi 绑定健康检查api
func BindHttpHealthApi(group *ghttp.RouterGroup) {
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
}
