package main

import (
	_ "svc1/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"svc1/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
