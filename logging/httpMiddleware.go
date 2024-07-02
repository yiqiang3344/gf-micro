package logging

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yiqiang3344/gf-micro/response"
)

func MiddlewareLogFormatJson(r *ghttp.Request) {
	glog.SetDefaultHandler(HandlerJson)

	r.Middleware.Next()
}

func MiddlewareHandlerAccessLog(r *ghttp.Request) {
	var (
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
		body   = r.GetRequestMap()
		res    interface{}
	)

	r.Middleware.Next()

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}

	res = r.Response.BufferString()
	if j, err := gjson.DecodeToJson(r.Response.BufferString()); err == nil {
		rTmp := new(response.DefaultHandlerResponse)
		if err = j.Scan(rTmp); err == nil {
			res = rTmp
		}
	}

	HttpAccessLog{
		HttpCode: r.Response.Status,
		Scheme:   scheme,
		Method:   r.Method,
		Host:     r.Host,
		Url:      r.URL.Path,
		Cost:     float64(gtime.TimestampMilli()-r.EnterTime.TimestampMilli()) / 1000,
		Req:      body,
		Res:      res,
		Ip:       r.GetClientIp(),
		Header:   r.Header,
	}.Log(r.GetCtx())
}

func MiddlewareHandlerErrorLog(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	if err == nil {
		return
	}

	ErrorLog{
		Method: r.Method,
	}.Log(r.Context(), err)
}
