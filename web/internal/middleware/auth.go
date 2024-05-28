package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	//// 登陆时不需要校验
	//whiteRouteList := []string{
	//}
	//
	//if gstr.InArray(whiteRouteList, r.URL.Path) {
	//	r.Middleware.Next()
	//	return
	//}
	//
	//token := r.GetHeader("token")
	//if token == "" {
	//	r.SetError(gerror.NewCode(code.WithCode(code.CodeLoginFailed, "请先登陆")))
	//	return
	//}
	//
	//userEntity, err := user.New().LoginByToken(r.GetCtx(), token)
	//if err != nil {
	//	r.SetError(err)
	//	return
	//}
	//
	//service.Context().InitRUser(r, userEntity)

	r.Middleware.Next()
}
