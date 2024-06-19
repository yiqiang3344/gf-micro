package middleware

import (
	"web/internal/logging"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareHandlerErrorLog(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	if err == nil {
		return
	}

	logging.ErrorLog{
		Method: r.Method,
	}.Log(r.Context(), err)
}
