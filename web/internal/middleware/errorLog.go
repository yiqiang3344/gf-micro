package middleware

import (
	"web/internal/logging"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

func MiddlewareHandlerErrorLog(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	if err == nil {
		return
	}

	var (
		code   = gerror.Code(err)
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
	)

	// 过滤业务错误码，大于999或等于-1的为业务错误码
	if code.Code() > 999 || code.Code() == -1 {
		return
	}

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	l := g.Log("error")
	l.SetStack(false)
	logging.ErrorLog{
		HttpCode:    r.Response.Status,
		Scheme:      scheme,
		Method:      r.Method,
		Host:        r.Host,
		Url:         r.URL.String(),
		Cost:        float64(gtime.TimestampMilli()-r.EnterTime.TimestampMilli()) / 1000,
		Ip:          r.GetClientIp(),
		ErrorCode:   code.Code(),
		ErrorMsg:    code.Message(),
		ErrorDetail: err.Error(),
	}.Log(r.Context())
}
