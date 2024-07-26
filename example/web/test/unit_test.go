package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yiqiang3344/gf-micro/flowColor"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"strings"
	"testing"
	"time"
)

type testInfo struct {
	filePath string
	route    string
}

// TestUnit 单元测试，可通过-args传参测试某些接口，也可以按顺序全量测试
// 跑之前需要先清空数据库和缓存，不能和全流程测试一起跑，会有数据冲突
// 跑之前需要启动web,user,blog,blogConsume几个服务
func TestUnit(t *testing.T) {
	funcListMap := gmap.NewListMap()

	funcListMap.Set("UserCreate", testInfo{filePath: "./userV1/create.xlsx", route: "/user/create"})
	funcListMap.Set("UserLogin", testInfo{filePath: "./userV1/login.xlsx", route: "/user/login"})
	funcListMap.Set("UserDetail", testInfo{filePath: "./userV1/detail.xlsx", route: "/user/detail"})
	funcListMap.Set("UserLogout", testInfo{filePath: "./userV1/logout.xlsx", route: "/user/logout"})
	funcListMap.Set("BlogCreate", testInfo{filePath: "./blogV1/create.xlsx", route: "/blog/create"})
	funcListMap.Set("BlogEdit", testInfo{filePath: "./blogV1/edit.xlsx", route: "/blog/edit"})
	funcListMap.Set("BlogDetail", testInfo{filePath: "./blogV1/detail.xlsx", route: "/blog/detail"})
	funcListMap.Set("BlogList", testInfo{filePath: "./blogV1/list.xlsx", route: "/blog/list"})
	funcListMap.Set("BlogDelete", testInfo{filePath: "./blogV1/delete.xlsx", route: "/blog/delete"})
	funcListMap.Set("BlogBatDelete", testInfo{filePath: "./blogV1/batDelete.xlsx", route: "/blog/bat-delete"})
	funcListMap.Set("BlogGetBatDeleteStatus", testInfo{filePath: "./blogV1/getBatDeleteStatus.xlsx", route: "/blog/get-bat-delete-status"})

	if !flag.Parsed() {
		flag.Parse()
	}
	//读取要运行的方法列表，参数逗号分割
	funcsStr := flag.Arg(0)
	var funcs []string
	if funcsStr != "" {
		funcs = strings.Split(flag.Arg(0), ",")
	}
	if len(funcs) > 0 {
		for _, v := range funcs {
			do(t, funcListMap.Get(v).(testInfo))
		}
	} else {
		funcListMap.Iterator(func(k, v interface{}) bool {
			do(t, v.(testInfo))
			return true
		})
	}
}

func do(t *testing.T, info testInfo) {
	ctx := flowColor.SetCtxFlowColor(context.Background(), "dev")
	o, err := testWithExcel.New(t, info.filePath,
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			if caseInfo.Extend["delayTime"] != "" {
				time.Sleep(gconv.Duration(caseInfo.Extend["delayTime"]) * time.Second)
			}

			client := GetClient(ctx)

			if caseInfo.Login["needLogin"] == "yes" {
				if caseInfo.Login["token"] != "" {
					client.SetHeader("token", caseInfo.Login["token"])
				} else if err = Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}

			ret = client.PostContent(ctx, info.route, caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(fmt.Errorf("接口%s运行异常:%v", info.route, err))
	}
	o.Run(ctx)
}
