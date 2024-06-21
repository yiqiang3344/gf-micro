package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/yiqiang3344/gf-micro/logging"
	"os"
)

type ConsumeFunc func(ctx context.Context, parser *gcmd.Parser) (stopFunc func(), err error)

// GetConsumeCmdFunc 标准化消费者服务方法
//
//	示例:
//	gcmd.Command{
//		Name:  "demoConsumer",
//		Usage: "./main demoConsumer",
//		Brief: "批量删除博客的消费者",
//		Func: cmd.GetConsumeCmdFunc("demoConsumer", cmd.GetGrpcMiddleware(), demoConsumeFunc),
//	}
func GetConsumeCmdFunc(name string, middleware MiddlewareForCmd, consumeFunc ConsumeFunc) func(ctx context.Context, parser *gcmd.Parser) error {
	return func(ctx context.Context, parser *gcmd.Parser) (err error) {
		shutdown := middleware(ctx)
		defer shutdown()

		stopFunc, err := consumeFunc(ctx, parser)
		if err != nil {
			reqs, _ := parser.MarshalJSON()
			logging.ErrorLog{
				Method: name,
				Req:    string(reqs),
			}.Log(ctx, err)
			panic(err)
		}

		gproc.AddSigHandlerShutdown(func(sig os.Signal) {
			stopFunc()
		})
		gproc.Listen()
		return
	}
}
