package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type ErrorLog struct {
	HttpCode    int     `json:"http_code"`
	Scheme      string  `json:"scheme"`
	Method      string  `json:"method"`
	Host        string  `json:"host"`
	Url         string  `json:"url"`
	Cost        float64 `json:"cost"`
	Ip          string  `json:"ip"`
	ErrorCode   int     `json:"error_code"`
	ErrorMsg    string  `json:"error_msg"`
	ErrorDetail string  `json:"error_detail"`
}

func (l ErrorLog) Log(ctx context.Context) {
	g.Log("error").Info(ctx, l)
}
