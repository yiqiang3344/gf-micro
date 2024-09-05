package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cfg"
)

var CheckCfgCmd = &gcmd.Command{
	Name:        "checkCfg",
	Usage:       "./main checkCfg",
	Brief:       "check config",
	Description: "检查配置文件是否符合规范",
	Arguments: append(CommonArguments, []gcmd.Argument{
		{
			Name:   "env",
			Short:  "e",
			Brief:  "要检查的环境，可选值: prod, dev。默认dev",
			IsArg:  false,
			Orphan: false,
		},
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
			Brief:  "强制检查服务注册发现redis配置",
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
	}...,
	),
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		cfg.Check(ctx, parser, GetCommonArguments(ctx, parser, Port).String())
		return
	},
}
