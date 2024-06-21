package cmd

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	_ "github.com/yiqiang3344/gf-micro/example/user/internal/logic"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cmd"
	"github.com/yiqiang3344/gf-micro/example/user/internal/controller/user"
)

var (
	// Main is the main command.
	Main = cmd.GenMain(
		gcmd.Command{
			Name:  "main",
			Usage: "main",
			Brief: "start user server",
			Func: cmd.GetGrpcCmdFunc(
				cmd.GetGrpcMiddleware(),
				func(server *grpcx.GrpcServer) {
					user.Register(server)
				},
			),
		},
	)
)
