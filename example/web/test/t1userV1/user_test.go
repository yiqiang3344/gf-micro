package t1userV1

import (
	"context"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"testing"
	v1 "web/api/user/v1"
	"web/test"
)

func TestUserCreate(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/create.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			data := v1.UserCreateReq{
				Nickname: caseInfo.Body["nickname"],
				Password: caseInfo.Body["password"],
			}
			ret = client.PostContent(ctx, "/user/create", data)
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
			data := v1.UserLoginReq{
				Nickname: caseInfo.Body["nickname"],
				Password: caseInfo.Body["password"],
			}
			ret = client.PostContent(ctx, "/user/login", data)
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

			data := v1.UserDetailReq{}
			ret = client.PostContent(ctx, "/user/detail", data)
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

			data := v1.UserLogoutReq{}
			ret = client.PostContent(ctx, "/user/logout", data)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}
