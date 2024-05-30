package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type AccessLog struct {
	Method string      `json:"method"`
	Cost   string      `json:"cost"`
	Req    interface{} `json:"req"`
	Res    interface{} `json:"res"`
	Error  error       `json:"error"`
}

func (l AccessLog) Log(ctx context.Context) {
	g.Log("access").Info(ctx, l)
}
