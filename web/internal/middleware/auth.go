package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"time"
	v1 "web/api/user/v1"
	"web/internal/service"
)

func MiddlewareAuth(r *ghttp.Request) {
	// 不需要校验的path
	whiteRouteList := []string{
		"/user/create",
		"/user/login",
	}
	// 登入的path
	loginRouteList := []string{
		"/user/login",
	}
	// 登出的path
	logoutRouteList := []string{
		"/user/logout",
	}

	defer func() {
		//如果登录成功，则设置cookie中的token，有效期24h
		if err := r.GetError(); err == nil && gstr.InArray(loginRouteList, r.URL.Path) {
			r.Cookie.SetCookie("token", r.GetHandlerResponse().(*v1.UserLoginRes).Token, r.Server.GetCookieDomain(), r.Server.GetCookiePath(), 24*time.Hour)
		}
		//如果成功登出，则清除cookie中的token
		if err := r.GetError(); err == nil && gstr.InArray(logoutRouteList, r.URL.Path) {
			r.Cookie.Remove("token")
		}
	}()
	if gstr.InArray(whiteRouteList, r.URL.Path) {
		r.Middleware.Next()
		return
	}
	user, err := service.User().LoginByToken(r.GetCtx(), getToken(r))
	if err != nil {
		r.SetError(err)
		return
	}
	service.Context().InitRUser(r, user)

	r.Middleware.Next()
}

func getToken(r *ghttp.Request) (token string) {
	//从header.token获取
	token = r.GetHeader("token")
	if !g.IsEmpty(token) {
		return
	}

	//从url.query参数获取
	token = r.GetQuery("token").String()
	if !g.IsEmpty(token) {
		return
	}

	//从cookie获取
	token = r.Cookie.Get("token").String()
	if !g.IsEmpty(token) {
		return
	}

	return
}
