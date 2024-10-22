package cmd

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cmd"
	"github.com/yiqiang3344/gf-micro/example/user/internal/controller/user"
	_ "github.com/yiqiang3344/gf-micro/example/user/internal/logic"
)

var (
	// Main is the main command.
	Main = cmd.GenMain(
		gcmd.Command{
			Name:      "main",
			Usage:     "main",
			Brief:     "start user grpc server",
			Arguments: append(cmd.CommonArguments, []gcmd.Argument{}...),
			Func: cmd.GetGrpcCmdFunc(
				cmd.GetGrpcMiddleware(),
				func(server *grpcx.GrpcServer) {
					user.Register(server)
				},
			),
		},
		cmd.GenHttpForGrpc,
		httpForGrpc,
	)
)
