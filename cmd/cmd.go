package cmd

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
)

const (
	ApolloIP = "apolloIp"
	Port     = "port"
)

var (
	// CommonArguments 命令行公共参数
	CommonArguments = []gcmd.Argument{
		{
			Name:   Port,
			Short:  "",
			Brief:  "端口,参考格式：`:8080`，默认使用配置文件的`server.address`",
			IsArg:  false,
			Orphan: false,
		},
		{
			Name:   ApolloIP,
			Short:  "",
			Brief:  "apollo连接地址，默认使用配置文件的`apollo.IP`或无",
			IsArg:  false,
			Orphan: false,
		},
	}
)

// GenMain 标准化生成应用主命令
func GenMain(mainSrv gcmd.Command, subSrv ...*gcmd.Command) gcmd.Command {
	err := mainSrv.AddCommand(subSrv...)
	if err != nil {
		panic(err)
	}
	return mainSrv
}

func GetCommonArguments(ctx context.Context, parser *gcmd.Parser, pattern string, isGrpc ...bool) (ret *gvar.Var) {
	switch pattern {
	case Port:
		if parser.GetOpt(Port) != nil {
			return parser.GetOpt(Port)
		}
		if len(isGrpc) > 0 && isGrpc[0] {
			return gcfg.Instance().MustGet(ctx, "grpc.address")
		} else {
			return gcfg.Instance().MustGet(ctx, "server.address")
		}
	case ApolloIP:
		if parser.GetOpt(ApolloIP) != nil {
			return parser.GetOpt(ApolloIP)
		}
		return gcfg.Instance().MustGet(ctx, "apollo.IP")
	}
	return
}
