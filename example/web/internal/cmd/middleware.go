package cmd

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yiqiang3344/gf-micro/response"
	"github.com/yiqiang3344/gf-micro/utility"
	"net/http"
	"reflect"
)

var notShowMsgCode = []gcode.Code{
	gcode.CodeInternalError,
	gcode.CodeDbOperationError,
	gcode.CodeInvalidOperation,
	gcode.CodeInvalidConfiguration,
	gcode.CodeMissingConfiguration,
	gcode.CodeNotImplemented,
	gcode.CodeOperationFailed,
	gcode.CodeSecurityReason,
	gcode.CodeServerBusy,
	gcode.CodeNecessaryPackageNotImport,
	gcode.CodeInternalPanic,
}

// HttpResponseMiddleware 部分业务状态码不显示错误信息
func HttpResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	if gstr.InArray(response.WhiteList, r.URL.Path) {
		return
	}

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		r.Response.ClearBuffer()
		r.Response.WriteJson(response.DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: gcode.CodeInternalError.Message(),
			Data:    g.Map{},
		})
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	//转化业务状态码
	codeTmp := code.Code()
	if code.Code() == 0 {
		msg = "success"
	}

	//部分状态码不显示错误明细
	if utility.InArray(notShowMsgCode, code) {
		msg = "内部错误"
	}

	//判断res是否为nil，为nil则转改空对象
	v := reflect.ValueOf(res)
	if v.Kind().String() == "invalid" || v.IsNil() {
		res = g.Map{}
	}

	r.Response.WriteJson(response.DefaultHandlerResponse{
		Code:    codeTmp,
		Message: msg,
		Data:    res,
	})
}
