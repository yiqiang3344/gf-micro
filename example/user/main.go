package main

import (
	_ "github.com/yiqiang3344/gf-micro/example/user/internal/logic"

	"github.com/yiqiang3344/gf-micro/example/user/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
