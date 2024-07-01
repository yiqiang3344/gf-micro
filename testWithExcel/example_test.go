package testWithExcel

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	tExample, err := New(t, "./example.xlsx",
		WithFailfast(true),
		WithBeforeFunc(func(ctx context.Context, prepareData PrepareData) {
			fmt.Printf("处理测试准备数据:%+v\n", prepareData)
			return
		}),
		WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo CaseInfo) (ret interface{}, err error) {
			fmt.Printf("开始处理测试用例%s，用例配置数据%+v,body数据%+v,登录数据%+v,扩展数据%+v,变量设置%s,变量数据%+v,用例描述:%s\n", caseInfo.Cfg.Name, caseInfo.Cfg, caseInfo.Body, caseInfo.Login, caseInfo.Extend, caseInfo.AssignVar, GetVarMap(), caseInfo.Desc)
			ret = caseInfo.Cfg.Name
			if strings.Contains(caseInfo.Cfg.Name, ":") || caseInfo.Cfg.Name == "cash_eq_assignVar" {
				ret = `{"status":1,"message":"success","data":{"test":1,"list":[{"id":1},{"id":2}]}}`
			}

			switch caseInfo.Cfg.Name {
			case "cash_gt", "cash_ge", "cash_lt", "cash_le", "cash_eq_useVar":
				ret = "1"
			case "case_none":
				Assert(caseInfo.Cfg.Name, ret, caseInfo.Cfg.Expect)
			}
			return
		}),
		WithAfterCaseFunc(func(ctx context.Context, caseInfo CaseInfo, caseRet interface{}, isCasePass bool) {
			fmt.Printf("用例%s结束,处理返回结果%+v,用例结果%v,变量数据%+v\n", caseInfo.Cfg.Name, caseRet, isCasePass, GetVarMap())
		}),
		WithAfterFunc(func(ctx context.Context, prepareData PrepareData, caseData CaseData, failCase *CaseInfo) {
			fmt.Printf("测试结束啦\n")
		}),
	)
	if err != nil {
		panic(err)
	}
	tExample.Run(context.Background())
}
