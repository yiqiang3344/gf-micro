package logging

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc/codes"
)

type ErrorLog struct {
	Method      string      `json:"method"`
	Cost        string      `json:"cost"`
	ErrorCode   codes.Code  `json:"error_code"`
	ErrorMsg    string      `json:"error_msg"`
	Req         interface{} `json:"req"`
	Res         interface{} `json:"res"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	ErrorDetail string      `json:"error_detail"`
	Stack       string      `json:"stack"`
}

func (l ErrorLog) Log(ctx context.Context) {
	g.Log("error").Info(ctx, l)
}
