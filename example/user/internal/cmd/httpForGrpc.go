package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/cmd"
	controller "github.com/yiqiang3344/gf-micro/example/user/internal/controller/http"
	"github.com/yiqiang3344/gf-micro/logging"
	"github.com/yiqiang3344/gf-micro/response"
)

var (
	httpForGrpc = &gcmd.Command{
		Name:      "httpForGrpc",
		Usage:     "./main httpForGrpc",
		Brief:     "start http for grpc server",
		Arguments: append(cmd.CommonArguments, []gcmd.Argument{}...),
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(g.Cfg().MustGet(ctx, cfg.APPNAME).String())
			s.SetAddr(cmd.GetCommonArguments(ctx, parser, cmd.Port).String())
			s.Use(
				logging.HttpLogFormatJsonMiddleware,
				logging.HttpAccessLogMiddleware,
				logging.HttpErrorLogMiddleware,
				response.HttpForGrpcResponseMiddleware,
			)

			s.BindObject("/{.struct}/{.method}", new(controller.User))

			s.Run()
			return
		},
	}
)
