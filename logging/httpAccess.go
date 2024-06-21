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
	Url      string                 `json:"url"`
	Cost     float64                `json:"cost"`
	Ip       string                 `json:"ip"`
	Body     map[string]interface{} `json:"body"`
	Response interface{}            `json:"response"`
	Header   http.Header            `json:"header"`
}

func (l HttpAccessLog) Log(ctx context.Context) {
	g.Log("access").Info(ctx, l)
}
