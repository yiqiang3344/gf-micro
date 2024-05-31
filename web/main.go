package main

import (
	"github.com/gogf/gf/v2/os/gcfg"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	_ "web/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"web/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()

	//配置中心
	if gcfg.Instance().MustGet(ctx, "apollo") != nil {
		adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
		if err != nil {
			panic(err)
		}
		gcfg.Instance().SetAdapter(adapter)
	}

	cmd.Main.Run(ctx)
}
