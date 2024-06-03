package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gcfg"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	_ "yijunqiang/gf-micro/user/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"yijunqiang/gf-micro/user/internal/cmd"
)

func main() {
	ctx := gctx.New()

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
