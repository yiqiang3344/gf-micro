package cmd

import (
	_ "web/internal/logic"
	_ "web/internal/packed"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/auth"
	"github.com/yiqiang3344/gf-micro/cmd"
	"time"
	v1 "web/api/user/v1"
	"web/internal/controller"
	"web/internal/service"
)

var (
	authWhitePaths = []string{
		"/user/create",
		"/user/login",
	}
	authLoginPaths = []string{
		"/user/login",
	}
	authLogoutPaths = []string{
		"/user/logout",
	}
	getLoginTokenFunc = func(r *ghttp.Request) (token string) {
		token = r.GetHandlerResponse().(*v1.UserLoginRes).Token
		return
	}
	cookieExpire = 24 * time.Hour
	Main         = cmd.GenMain(
		gcmd.Command{
			Name:  "main",
			Usage: "main",
			Brief: "start http server",
			Func: cmd.GetHttpCmdFunc(
				cmd.GetHttpMiddleware(),
				[]ghttp.HandlerFunc{},
				[]cmd.ServerGroup{
					{
						Prefix: "/",
						Middlewares: []ghttp.HandlerFunc{
							auth.GetHttpMiddleware(
								authWhitePaths,
								authLoginPaths,
								authLogoutPaths,
								service.User().LoginByToken,
								getLoginTokenFunc,
								cookieExpire,
							),
						},
						Controllers: []interface{}{
							controller.User,
							controller.Blog,
						},
					},
				},
			),
		},
	)
)
