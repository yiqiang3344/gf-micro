package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/logging"
)

// GetXxljobCmdFunc 标准化xxljob服务方法
//
//	示例:
//	gcmd.Command{
//		Name:  "xxlJob",
//		Usage: "./main xxlJob",
//		Brief: "定时任务",
//		Func: cmd.GetConsumeCmdFunc(cmd.GetGrpcMiddleware(), map[string]xxl.TaskFunc{
//			"task1": func(cxt context.Context, param *RunReq) string{return ""},
//			"task2": func(cxt context.Context, param *RunReq) string{return ""}
//		}]),
//	}
func GetXxljobCmdFunc(middleware MiddlewareForCmd, taskMap map[string]xxl.TaskFunc) func(ctx context.Context, parser *gcmd.Parser) error {
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		exec := xxl.NewExecutor(
			xxl.ServerAddr(g.Cfg().MustGet(ctx, "xxljob.serverAddr").String()),
			xxl.AccessToken(g.Cfg().MustGet(ctx, "xxljob.token").String()),
			xxl.RegistryKey(g.Cfg().MustGet(ctx, cfg.APPNAME).String()), //执行器名称
			xxl.SetLogger(&logger{}),
		)
		exec.Init()

		//注册任务handler
		for k, v := range taskMap {
			exec.RegTask(k, func(cxt context.Context, param *xxl.RunReq) string {
				shutdown := middleware(ctx)
				defer shutdown()

				//生成链路主span
				ctx1, span := gtrace.NewSpan(ctx, fmt.Sprintf("xxljob/%s", k))
				defer func() {
					span.End()
				}()

				return v(ctx1, param)
			})
		}

		err = exec.Run()
		if err != nil {
			logging.ErrorLog{
				Method: "xxlJob",
			}.Log(ctx, err)
		}
		return
	}
}

// xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	g.Log("xxljob").Infof(context.Background(), format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	logging.ErrorLog{
		Method: "xxljob",
	}.Log(context.Background(), gerror.Newf(format, a...))
}
