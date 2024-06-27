package testFrame

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestTest(t *testing.T) {
	tI, err := New(t, "./demo_test.xlsx", WithCaseHandleFunc(func(ctx context.Context, caseInfo CaseInfo) (ret interface{}, err error) {
		g.Dump(caseInfo)
		return
	}))
	if err != nil {
		panic(err)
	}
	tI.Run(context.Background())
}
