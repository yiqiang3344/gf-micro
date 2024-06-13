package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
	"os"
	"yijunqiang/gf-micro/blog/internal/service"
)

var (
	batDeleteBlogConsumer = &gcmd.Command{
		Name:  "batDeleteBlogConsumer",
		Usage: "./main batDeleteBlogConsumer",
		Brief: "批量删除博客的消费者",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			shutdown := initMiddleware(ctx)
			defer shutdown()

			stopFunc, err := service.Blog().BatDeleteConsumer(ctx)
			if err != nil {
				g.Log("debug").Errorf(ctx, "batDeleteBlogConsumer异常:%+v", err)
				panic(err)
			}

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				stopFunc()
			})
			gproc.Listen()
			return
		},
	}
)
