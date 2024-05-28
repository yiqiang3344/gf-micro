package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type BizLog struct {
	Tag     string      `json:"tag"`
	Message interface{} `json:"message"`
}

func (l BizLog) Log(ctx context.Context) {
	g.Log("biz").Info(ctx, l)
}
