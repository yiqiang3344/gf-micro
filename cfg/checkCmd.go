package cfg

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
)

var CheckCmd = &gcmd.Command{
	Name:        "checkCfg",
	Usage:       "./main checkCfg",
	Brief:       "check config",
	Description: "检查配置文件是否符合规范",
	Arguments: []gcmd.Argument{
		{
			Name:   "all",
			Short:  "a",
			Brief:  "强制检查下方所有配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "grpc",
			Short:  "g",
			Brief:  "强制检查grpc服务配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "server",
			Short:  "s",
			Brief:  "强制检查server服务配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "rocketmq",
			Short:  "m",
			Brief:  "强制检查rocketmq配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "xxljob",
			Short:  "j",
			Brief:  "强制检查xxlJob配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "apollo",
			Short:  "c",
			Brief:  "强制检查apollo配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "otlp",
			Short:  "o",
			Brief:  "强制检查otlp链路追踪配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "registry",
			Short:  "y",
			Brief:  "强制检查服务注册发现registry配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "redis",
			Short:  "r",
			Brief:  "强制检查服务注册发现registry配置",
			IsArg:  false,
			Orphan: true,
		},
		{
			Name:   "database",
			Short:  "d",
			Brief:  "强制检查数据库配置",
			IsArg:  false,
			Orphan: true,
		},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		var (
			errs []string
		)

		//判断是否检查apollo配置
		if !parser.GetOpt("all").IsNil() || !parser.GetOpt("apollo").IsNil() {
			errs = append(errs, checkRules(ctx, gcfg.Instance(), apolloRules)...)
		}

		//先判断是否有apollo配置，有的话接入apollo
		if !gcfg.Instance().MustGet(ctx, "apollo").IsNil() {
			adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
			if err != nil {
				panic(err)
			}
			gcfg.Instance().SetAdapter(adapter)
		}

		if !parser.GetOpt("grpc").IsNil() {
			errs = append(errs, checkRules(ctx, gcfg.Instance(), grpcRules)...)
		}

		//打印结果
		printErr(errs)
		return
	},
}
