package testWithExcel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"strings"
)

var HttpCmd = &gcmd.Command{
	Name:        "genUnitTestExcel",
	Usage:       "./main genUnitTestExcel",
	Brief:       "Gen Unit Test Excel",
	Description: "注意：执行之前请确保openApi开启，且本地http服务已启动，且配置文件中指定了服务端口，脚本会根据配置文件自动获取openApi的json文件并生成单元测试用例的excel模版",
	Arguments: []gcmd.Argument{
		{
			Name:   "output",
			Short:  "o",
			Brief:  "输出的目录,为空则输出到`test/unitExcel/`",
			IsArg:  false,
			Orphan: false,
		},
		{
			Name:   "include",
			Short:  "i",
			Brief:  "过滤条件，要包含的接口，格式为method|path，多个则逗号分割，如: get|/user/login,post|/user/create",
			IsArg:  false,
			Orphan: false,
		},
		{
			Name:   "expel",
			Short:  "e",
			Brief:  "过滤条件，要排出的接口，格式为method|path，多个则逗号分割，如: get|/user/login,post|/user/create",
			IsArg:  false,
			Orphan: false,
		},
	},
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		//接入配置中心
		if gcfg.Instance().MustGet(ctx, "apollo") != nil {
			adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
			if err != nil {
				panic(err)
			}
			gcfg.Instance().SetAdapter(adapter)
		}

		//获取端口
		address := g.Cfg().MustGet(ctx, "server.address").String()
		if !strings.Contains(address, ":") {
			panic(fmt.Errorf("配置server.address格式错误: %s", address))
		}
		port := gstr.Split(g.Cfg().MustGet(ctx, "server.address").String(), ":")[1]
		if gconv.Int(port) == 0 {
			panic(fmt.Errorf("配置server.address中端口格式错误: %s", port))
		}

		//获取openApi的json的Path
		openapiPath := g.Cfg().MustGet(ctx, "server.openapiPath").String()
		if openapiPath == "" {
			panic(fmt.Errorf("配置server.openapiPath未配置"))
		}

		//输出目录
		output := parser.GetOpt("output").String()
		if output == "" {
			//默认目录
			output = "test/unitExcel"
		}
		//目录不存在则创建
		if !gfile.IsDir(output) {
			err = gfile.Mkdir(output)
			if err != nil {
				panic(err)
			}
		}

		//过滤条件
		var filters []*Filter
		include := strings.Trim(parser.GetOpt("include").String(), " ")
		if include != "" {
			for _, v := range strings.Split(include, ",") {
				if !strings.Contains(v, "|") {
					panic(fmt.Errorf("参数%s格式不正确", "include"))
				}
				arr := strings.Split(v, "|")
				filters = append(filters, &Filter{
					Path:   arr[1],
					Method: arr[0],
					Type:   FilterTypeInclude,
				})
			}
		}
		expel := strings.Trim(parser.GetOpt("expel").String(), " ")
		if expel != "" {
			for _, v := range strings.Split(expel, ",") {
				if !strings.Contains(v, "|") {
					panic(fmt.Errorf("参数%s格式不正确", "expel"))
				}
				arr := strings.Split(v, "|")
				filters = append(filters, &Filter{
					Path:   arr[1],
					Method: arr[0],
					Type:   FilterTypeExpel,
				})
			}
		}

		//初始化客户端
		prefix := fmt.Sprintf("http://127.0.0.1:%s", port)
		client := g.Client()
		client.SetPrefix(prefix)

		//查询openApi的json
		jsonContent := client.GetContent(ctx, openapiPath)

		if err = GenTestCaseExcelByOpenApiJson(jsonContent, output, filters); err != nil {
			panic(err)
		}

		return
	},
}
