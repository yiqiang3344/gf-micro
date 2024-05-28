package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/glog"
	"yijunqiang/gf-micro/user/internal/logging"

	_ "yijunqiang/gf-micro/user/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"yijunqiang/gf-micro/user/internal/cmd"
)

func main() {
	//日志json化
	glog.SetDefaultHandler(logging.HandlerJson)

	cmd.Main.Run(gctx.New())
}
