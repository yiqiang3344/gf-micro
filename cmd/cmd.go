package cmd

import "github.com/gogf/gf/v2/os/gcmd"

// GenMain 标准化生成应用主命令
func GenMain(mainSrv gcmd.Command, subSrv ...*gcmd.Command) gcmd.Command {
	err := mainSrv.AddCommand(subSrv...)
	if err != nil {
		panic(err)
	}
	return mainSrv
}
