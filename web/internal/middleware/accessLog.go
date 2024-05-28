package middleware

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"web/internal/logging"
)

func MiddlewareHandlerAccessLog(r *ghttp.Request) {
	var (
		scheme   = "http"
		proto    = r.Header.Get("X-Forwarded-Proto")
		body     = r.GetRequestMap()
		response interface{}
	)

	r.Middleware.Next()

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}

	response = r.Response.BufferString()
	if j, err := gjson.DecodeToJson(r.Response.BufferString()); err == nil {
		_r := new(DefaultHandlerResponse)
		if err = j.Scan(_r); err == nil {
			response = _r
		}
	}

	logging.AccessLog{
		HttpCode: r.Response.Status,
		Scheme:   scheme,
		Method:   r.Method,
		Host:     r.Host,
		Url:      r.URL.Path,
		Cost:     float64(gtime.TimestampMilli()-r.EnterTime.TimestampMilli()) / 1000,
		Ip:       r.GetClientIp(),
		Body:     body,
		Response: response,
		Header:   r.Header,
	}.Log(r.GetCtx())
}
