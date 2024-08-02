package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
)

type HttpAccessLog struct {
	HttpCode int                    `json:"http_code"`
	Scheme   string                 `json:"scheme"`
	Method   string                 `json:"method"`
	Host     string                 `json:"host"`
	Path     string                 `json:"path"`
	Cost     float64                `json:"cost"`
	Req      map[string]interface{} `json:"req"`
	Res      interface{}            `json:"res"`
	Ip       string                 `json:"ip"`
	Header   http.Header            `json:"header"`
}

func (l HttpAccessLog) Log(ctx context.Context) {
	g.Log("access").Info(ctx, l)
}
