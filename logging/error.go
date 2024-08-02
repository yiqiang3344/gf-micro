package logging

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"strconv"
)

type ErrorLog struct {
	Method  string      `json:"method"`  //异常的方法
	Req     interface{} `json:"req"`     //异常方法的入参
	Code    int         `json:"code"`    //异常错误码
	Message string      `json:"message"` //异常信息
	Detail  interface{} `json:"detail"`  //异常详情
}

var ignoreErrCodeStrs = []string{
	strconv.Itoa(gcode.CodeNotAuthorized.Code()),
	strconv.Itoa(gcode.CodeValidationFailed.Code()),
	strconv.Itoa(gcode.CodeInvalidParameter.Code()),
	strconv.Itoa(gcode.CodeMissingParameter.Code()),
	strconv.Itoa(gcode.CodeNotFound.Code()),
	strconv.Itoa(gcode.CodeBusinessValidationFailed.Code()),
}

func (l ErrorLog) Log(ctx context.Context, err ...error) {
	stack := ""
	if len(err) > 0 {
		e := err[0]
		code := gerror.Code(e)
		l.Code = code.Code()
		l.Message = code.Message()
		l.Detail = code.Detail()
		if l.Detail == nil {
			l.Detail = e.Error()
		}
		if gerror.HasStack(e) {
			stack = "\n" + gerror.Stack(e)
		}
	}
	//过滤用户操作检查类异常，比如参数校验，信息检查之类的
	if gstr.InArray(ignoreErrCodeStrs, strconv.Itoa(l.Code)) || l.Code > 300 {
		return
	}
	log := g.Log("error").Stack(true, 1)
	vs := []interface{}{l}
	if stack != "" {
		//首先使用自定义的stack信息，没有则使用默认堆栈信息
		log.SetStack(false)
		vs = append(vs, stack)
	}
	log.Error(ctx, vs...)
}
