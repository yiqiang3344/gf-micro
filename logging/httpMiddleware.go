package logging

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yiqiang3344/gf-micro/response"
	"io"
	"net/http"
	"time"
)

func HttpLogFormatJsonMiddleware(r *ghttp.Request) {
	glog.SetDefaultHandler(HandlerJson)

	r.Middleware.Next()
}

func HttpAccessLogMiddleware(r *ghttp.Request) {
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
		Path:     r.URL.Path,
		Cost:     float64(gtime.TimestampMilli()-r.EnterTime.TimestampMilli()) / 1000,
		Req:      body,
		Res:      res,
		Ip:       r.GetClientIp(),
		Header:   r.Header,
	}.Log(r.GetCtx())
}

func HttpErrorLogMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	if err == nil {
		return
	}

	ErrorLog{
		Method: r.Method,
		Req:    r.GetRequestMap(),
	}.Log(r.Context(), err)
}

func HttpClientLogMiddleware(c *gclient.Client, r *http.Request) (response *gclient.Response, err error) {
	var (
		ctx    = r.Context()
		start  = time.Now()
		scheme = "http"
		proto  = r.Header.Get("X-Forwarded-Proto")
		req    = map[string]interface{}{}
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}

	switch r.Method {
	case http.MethodGet, http.MethodOptions, http.MethodHead:
		req, _ = gstr.Parse(r.URL.Query().Encode())
	default:
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		req, _ = gstr.Parse(string(bodyBytes))
	}

	response, err = c.Next(r)
	duration := time.Since(start)
	resBytes, _ := io.ReadAll(response.Body)
	response.Body = io.NopCloser(bytes.NewBuffer(resBytes))
	HttpClientLog{
		HttpCode: response.StatusCode,
		Scheme:   scheme,
		Method:   r.Method,
		Host:     r.Host,
		Path:     r.URL.Path,
		Cost:     fmt.Sprintf("%.3fms", float64(duration)/1e6),
		Req:      req,
		Res:      string(resBytes),
		Header:   r.Header,
	}.Log(ctx)
	return
}
