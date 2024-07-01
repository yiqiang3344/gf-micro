package test

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"testing"
	"time"
)

func TestFullFlow(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/fullFlow.xlsx",
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
	o.Run(context.Background())
}
