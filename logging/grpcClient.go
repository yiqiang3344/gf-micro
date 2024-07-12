package logging

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type GrpcClientLog struct {
	Path  string      `json:"path"`
	Cost  string      `json:"cost"`
	Req   interface{} `json:"req"`
	Res   interface{} `json:"res"`
	Error errorI      `json:"error"`
}

func (l GrpcClientLog) Log(ctx context.Context, err ...error) {
	if len(err) > 0 && err[0] != nil {
		code := gerror.Code(err[0])
		e := errorI{
			Code:    code.Code(),
			Message: code.Message(),
			Detail:  code.Detail(),
		}
		if code.Detail() == nil {
			e.Detail = err[0].Error()
		}
		l.Error = e
	}
	g.Log("webclient").Info(ctx, l)
}