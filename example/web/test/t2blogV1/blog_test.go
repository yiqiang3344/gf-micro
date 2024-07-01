package t2blogV1

import (
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"golang.org/x/net/context"
	"testing"
	"web/test"
)

var (
	testDataFile = "./testdata/blog_test.xlsx"
)

func TestBlogCreate(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/create.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/create", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogEdit(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/edit.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/edit", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogDetail(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/detail.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/detail", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogList(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/list.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/list", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogDelete(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/delete.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/delete", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogBatDelete(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/batDelete.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/bat-delete", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}

func TestBlogGetBatDeleteStatus(t *testing.T) {
	o, err := testWithExcel.New(t, "./testdata/getBatDeleteStatus.xlsx",
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			client := test.GetClient(ctx)
			if caseInfo.Login["needLogin"] == "yes" {
				if err = test.Login(ctx, caseInfo.Login["nickname"], caseInfo.Login["password"], client); err != nil {
					return
				}
			}
			ret = client.PostContent(ctx, "/blog/get-bat-delete-status", caseInfo.Body)
			return
		}),
	)
	if err != nil {
		panic(err)
	}
	o.Run(context.Background())
}
