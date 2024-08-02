package flowColor

import (
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

// HttpServerMiddleware http服务端流量染色中间件
func HttpServerMiddleware(r *ghttp.Request) {
	if !IsOpen() {
		r.Middleware.Next()
		return
	}

	//没有流量染色标识的话，默认访问基准服务
	color := ColorBase
	if v := r.Header.Get(FlowColor); v != "" {
		color = v
	}

	r.SetCtx(SetCtxFlowColor(r.Context(), color))

	r.Middleware.Next()
}

// HttpClientMiddleware http客户端流量染色中间件
func HttpClientMiddleware(c *gclient.Client, r *http.Request) (response *gclient.Response, err error) {
	ctx := SetCtxFlowColor(r.Context(), ColorBase)
	r.Header.Set(FlowColor, *GetCtxFlowColor(ctx))
	return c.Next(r)
}
