package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"time"
)

// GetHttpMiddleware 获取http鉴权中间件
//
//	参数说明：
//		whitePaths 不需要校验的接口path
//		loginPaths 登入的接口path
//		logoutPaths 登出的接口path
//		loginByTokenFunc 根据token获取鉴权用户信息的方法
//		getLoginTokenFunc 从登录类接口响应信息中获取token的方法，如 r.GetHandlerResponse().(*v1.UserLoginRes).Token
//		cookieExpire cookie有效期
func GetHttpMiddleware(
	whitePaths,
	loginPaths,
	logoutPaths []string,
	loginByTokenFunc func(ctx context.Context, token string) (*User, error),
	getLoginTokenFunc func(r *ghttp.Request) string,
	cookieExpire time.Duration,
) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		defer func() {
			//如果登录成功，则设置cookie中的token，有效期24h
			if err := r.GetError(); err == nil && gstr.InArray(loginPaths, r.URL.Path) {
				r.Cookie.SetCookie("token", getLoginTokenFunc(r), r.Server.GetCookieDomain(), r.Server.GetCookiePath(), cookieExpire)
			}
			//如果成功登出，则清除cookie中的token
			if err := r.GetError(); err == nil && gstr.InArray(logoutPaths, r.URL.Path) {
				r.Cookie.Remove("token")
			}
		}()
		//强行加入健康检查path
		whitePaths = append(whitePaths, "/health")
		if gstr.InArray(whitePaths, r.URL.Path) {
			r.Middleware.Next()
			return
		}
		user, err := loginByTokenFunc(r.GetCtx(), getToken(r))
		if err != nil {
			r.SetError(err)
			return
		}
		GetStandardAuth().InitRUser(r, user)

		r.Middleware.Next()
	}
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
