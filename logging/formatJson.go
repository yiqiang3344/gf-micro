package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yiqiang3344/gf-micro/cfg"
	"strings"
)

type HandlerOutputJson struct {
	Time       string `json:"time"`               // Formatted time string, like "2016-01-09 12:00:00".
	TraceId    string `json:"trace_id,omitempty"` // Trace id, only available if tracing is enabled.
	Category   string `json:"category"`
	CtxStr     string `json:"ctx_str,omitempty"`     // The retrieved context value string from context, only available if Config.CtxKeys configured.
	Level      string `json:"level"`                 // Formatted level string, like "DEBU", "ERRO", etc. Eg: ERRO
	CallerFunc string `json:"caller_func,omitempty"` // The source function name that calls logging, only available if F_CALLER_FN set.
	CallerPath string `json:"caller_path,omitempty"` // The source file path and its line number that calls logging, only available if F_FILE_SHORT or F_FILE_LONG set.
	Prefix     string `json:"prefix,omitempty"`      // Custom prefix string for logging content.
	Content    any    `json:"content"`               // Content is the main logging content, containing error stack string produced by logger.
	Stack      string `json:"stack,omitempty"`       // Stack string produced by logger, only available if Config.StStatus configured.
}

func HandlerJson(ctx context.Context, in *glog.HandlerInput) {
	//根据日志文件判断日志类别
	logCategory := strings.Split(in.Logger.GetConfig().File, ".")[0]
	//在配置中的才json做转换
	if gstr.InArray(g.Config().MustGet(ctx, cfg.JSONFORMATLOGS).Strings(), logCategory) {
		output := HandlerOutputJson{
			Time:       in.TimeFormat,
			TraceId:    in.TraceId,
			Category:   logCategory,
			CtxStr:     in.CtxStr,
			Level:      in.LevelFormat,
			CallerFunc: in.CallerFunc,
			CallerPath: in.CallerPath,
			Prefix:     in.Prefix,
			Content:    in.Values[0],
			Stack:      in.Stack,
		}
		bf := bytes.NewBuffer([]byte{})
		enc := json.NewEncoder(bf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(output)
		if err != nil {
			panic(err)
		}
		in.Buffer.Write(bf.Bytes())
	}
	in.Next(ctx)
}
