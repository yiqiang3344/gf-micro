package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type HttpClientLog struct {
	HttpCode int                    `json:"http_code"`
	Scheme   string                 `json:"scheme"`
	Method   string                 `json:"method"`
	Host     string                 `json:"host"`
	Path     string                 `json:"path"`
	Cost     string                 `json:"cost"`
	Req      map[string]interface{} `json:"req"`
	Res      interface{}            `json:"res"`
	Header   http.Header            `json:"header"`
}

func (l HttpClientLog) Log(ctx context.Context) {
	g.Log("webclient").Info(ctx, l)
}
