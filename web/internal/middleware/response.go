package middleware

import (
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
	"reflect"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

var whiteList = []string{
	"/api.json",
}

func MiddlewareHandlerResponse(r *ghttp.Request) {
	//r.SetCtx(gi18n.WithLanguage(r.Context(), "zh-CN")) //相应设置为中文

	r.Middleware.Next()

	if gstr.InArray(whiteList, r.URL.Path) {
		return
	}

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		r.Response.ClearBuffer()
		r.Response.WriteJson(DefaultHandlerResponse{
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

	//判断res是否为nil，为nil则转改空对象
	v := reflect.ValueOf(res)
	if v.Kind().String() == "invalid" || v.IsNil() {
		res = g.Map{}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    codeTmp,
		Message: msg,
		Data:    res,
	})
}
