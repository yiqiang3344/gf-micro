package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "yijunqiang/gf-micro/blog/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"yijunqiang/gf-micro/blog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
