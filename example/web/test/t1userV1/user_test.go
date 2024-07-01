package t1userV1

import (
	"context"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"testing"
	"web/test"
)

func TestUserCreate(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/create.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			ret = client.PostContent(ctx, "/user/create", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestUserLogin(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/login.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			ret = client.PostContent(ctx, "/user/login", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestUserDetail(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/detail.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)

			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}

			ret = client.PostContent(ctx, "/user/detail", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestUserLogout(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/logout.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)

			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}

			ret = client.PostContent(ctx, "/user/logout", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}
