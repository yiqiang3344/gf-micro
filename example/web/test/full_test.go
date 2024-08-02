package test

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yiqiang3344/gf-micro/flowColor"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"testing"
	"time"
)

// TestFull 全流程测试
// 跑之前需要先清空数据库和缓存，不能和单元测试一起跑，会有数据冲突
// 跑之前需要启动web,user,blog,blogConsume几个服务
func TestFull(t *testing.T) {
	ctx := flowColor.SetCtxFlowColor(context.Background(), "local")
	o, err := testWithExcel.New(t, "./full.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			if caseInfo.Extend["delayTime"] != "" {
				time.Sleep(gconv.Duration(caseInfo.Extend["delayTime"]) * time.Second)
			}

			client := GetClient(ctx)

			if caseInfo.Login["needLogin"] == "yes" {
				client.SetHeader("token", caseInfo.Login["token"])
			}

			ret = client.PostContent(ctx, caseInfo.Extend["path"], caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(ctx)
}
