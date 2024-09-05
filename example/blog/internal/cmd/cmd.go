package cmd

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/xxl-job/xxl-job-executor-go"
	_ "github.com/yiqiang3344/gf-micro/example/blog/internal/logic"

	"github.com/yiqiang3344/gf-micro/cmd"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/service"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/controller/blog"
)

var (
	Main = cmd.GenMain(
		gcmd.Command{
			Name:        "main",
			Usage:       "./main",
			Brief:       "博客服务管理工具",
			Description: `功能包括：微服务启动(默认)，消费者启动，xxljob定时任务等`,
			Arguments:   append(cmd.CommonArguments, []gcmd.Argument{}...),
			Func: cmd.GetGrpcCmdFunc(
				cmd.GetGrpcMiddleware(),
				func(server *grpcx.GrpcServer) {
					blog.Register(server)
				},
			),
		},
		&gcmd.Command{
			Name:  "batDeleteBlogConsumer",
			Usage: "./main batDeleteBlogConsumer",
			Brief: "批量删除博客的消费者",
			Func: cmd.GetConsumeCmdFunc(
				"batDeleteBlogConsumer",
				cmd.GetGrpcMiddleware(),
				service.Blog().BatDeleteConsumer,
			),
		},
		&gcmd.Command{
			Name:  "xxljob",
			Usage: "./main xxljob",
			Brief: "定时任务",
			Func: cmd.GetXxljobCmdFunc(
				cmd.GetGrpcMiddleware(),
				map[string]xxl.TaskFunc{
					"stats": service.Blog().Stats,
				},
			),
		},
	)
)
