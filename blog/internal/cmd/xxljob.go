package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/xxl-job/xxl-job-executor-go"
	"yijunqiang/gf-micro/blog/internal/logging"
	"yijunqiang/gf-micro/blog/internal/service"
)

var (
	xxljob = &gcmd.Command{
		Name:  "xxljob",
		Usage: "./main xxljob",
		Brief: "对接xxljob的定时任务",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			shutdown := initMiddleware(ctx)
			defer shutdown()

			exec := xxl.NewExecutor(
				xxl.ServerAddr(g.Cfg().MustGet(ctx, "xxljob.serverAddr").String()),
				xxl.AccessToken(g.Cfg().MustGet(ctx, "xxljob.token").String()),
				xxl.RegistryKey(g.Cfg().MustGet(ctx, "appName").String()), //执行器名称
				xxl.SetLogger(&logger{}),
			)
			exec.Init()

			//注册任务handler
			exec.RegTask("stats", service.Blog().Stats)

			err = exec.Run()
			if err != nil {
				logging.ErrorLog{
					Method: "xxljob",
				}.Log(ctx, err)
			}
			return
		},
	}
)

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
