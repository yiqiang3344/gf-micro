package main

import (
	_ "web/internal/logic"

	_ "web/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"web/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
