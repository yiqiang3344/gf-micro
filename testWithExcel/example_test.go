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

func TestGen(t *testing.T) {
	err := GenTestCaseExcelByOpenApiJson(
		`{"openapi":"3.0.0","components":{"schemas":{"web.api.blog.v1.BlogBatDeleteReq":{"properties":{"ids":{"description":"ID列表,逗号分割","format":"string","properties":{},"type":"string"}},"required":["ids"],"type":"object"},"web.api.blog.v1.BlogBatDeleteRes":{"properties":{"batNo":{"description":"批次号","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.blog.v1.BlogCreateReq":{"properties":{"title":{"description":"标题","format":"string","properties":{},"type":"string"},"content":{"description":"内容","format":"string","properties":{},"type":"string"}},"required":["title","content"],"type":"object"},"web.api.blog.v1.BlogCreateRes":{"properties":{},"type":"object"},"web.api.blog.v1.BlogDeleteReq":{"properties":{"id":{"description":"ID","format":"string","properties":{},"type":"string"}},"required":["id"],"type":"object"},"web.api.blog.v1.BlogDeleteRes":{"properties":{},"type":"object"},"web.api.blog.v1.BlogDetailReq":{"properties":{"id":{"description":"ID","format":"string","properties":{},"type":"string"}},"required":["id"],"type":"object"},"web.api.blog.v1.BlogDetailRes":{"properties":{"id":{"description":"ID","format":"uint32","properties":{},"type":"integer"},"title":{"description":"标题","format":"string","properties":{},"type":"string"},"content":{"description":"内容","format":"string","properties":{},"type":"string"},"nickname":{"description":"作者","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.blog.v1.BlogEditReq":{"properties":{"id":{"description":"ID","format":"string","properties":{},"type":"string"},"title":{"description":"标题","format":"string","properties":{},"type":"string"},"content":{"description":"内容","format":"string","properties":{},"type":"string"}},"required":["id","title","content"],"type":"object"},"web.api.blog.v1.BlogEditRes":{"properties":{},"type":"object"},"web.api.blog.v1.BlogGetBatDeleteStatusReq":{"properties":{"batNo":{"description":"批次号","format":"string","properties":{},"type":"string"}},"required":["batNo"],"type":"object"},"web.api.blog.v1.BlogGetBatDeleteStatusRes":{"properties":{"status":{"description":"状态","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.blog.v1.BlogListReq":{"properties":{},"type":"object"},"web.api.blog.v1.BlogListRes":{"properties":{"list":{"description":"博客列表","format":"[]*model.BlogDetailOutput","items":{"$ref":"#/components/schemas/web.internal.model.BlogDetailOutput"},"properties":{},"type":"array"}},"type":"object"},"web.internal.model.BlogDetailOutput":{"properties":{"id":{"description":"ID","format":"uint32","properties":{},"type":"integer"},"title":{"description":"标题","format":"string","properties":{},"type":"string"},"content":{"description":"内容","format":"string","properties":{},"type":"string"},"nickname":{"description":"作者","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.user.v1.UserCreateReq":{"properties":{"nickname":{"description":"用户名","format":"string","properties":{},"type":"string"},"password":{"description":"密码","format":"string","properties":{},"type":"string"}},"required":["nickname","password"],"type":"object"},"web.api.user.v1.UserCreateRes":{"properties":{},"type":"object"},"web.api.user.v1.UserDetailReq":{"properties":{},"type":"object"},"web.api.user.v1.UserDetailRes":{"properties":{"id":{"description":"ID","format":"uint32","properties":{},"type":"integer"},"nickname":{"description":"昵称","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.user.v1.UserLoginReq":{"properties":{"nickname":{"description":"用户名","format":"string","properties":{},"type":"string"},"password":{"description":"密码","format":"string","properties":{},"type":"string"}},"required":["nickname","password"],"type":"object"},"web.api.user.v1.UserLoginRes":{"properties":{"token":{"description":"token","format":"string","properties":{},"type":"string"}},"type":"object"},"web.api.user.v1.UserLogoutReq":{"properties":{},"type":"object"},"web.api.user.v1.UserLogoutRes":{"properties":{},"type":"object"}}},"info":{"title":"","version":""},"paths":{"/blog/bat-delete":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogBatDeleteReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogBatDeleteRes"}}},"description":""}},"summary":"博客/批量删除","tags":["Blog"]}},"/blog/create":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogCreateReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogCreateRes"}}},"description":""}},"summary":"博客/创建","tags":["Blog"]}},"/blog/delete":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogDeleteReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogDeleteRes"}}},"description":""}},"summary":"博客/删除","tags":["Blog"]}},"/blog/detail":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogDetailReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogDetailRes"}}},"description":""}},"summary":"博客/详情","tags":["Blog"]}},"/blog/edit":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogEditReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogEditRes"}}},"description":""}},"summary":"博客/编辑","tags":["Blog"]}},"/blog/get-bat-delete-status":{"get":{"parameters":[{"description":"批次号","in":"query","name":"batNo","required":true,"schema":{"description":"批次号","format":"string","properties":{},"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogGetBatDeleteStatusRes"}}},"description":""}},"summary":"博客/获取批量删除状态","tags":["Blog"]},"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogGetBatDeleteStatusReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogGetBatDeleteStatusRes"}}},"description":""}},"summary":"博客/获取批量删除状态","tags":["Blog"]}},"/blog/list":{"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogListReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.blog.v1.BlogListRes"}}},"description":""}},"summary":"博客/列表","tags":["Blog"]}},"/user/create":{"get":{"parameters":[{"description":"用户名","in":"query","name":"nickname","required":true,"schema":{"description":"用户名","format":"string","properties":{},"type":"string"}},{"description":"密码","in":"query","name":"password","required":true,"schema":{"description":"密码","format":"string","properties":{},"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserCreateRes"}}},"description":""}},"summary":"用户/注册","tags":["User"]},"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserCreateReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserCreateRes"}}},"description":""}},"summary":"用户/注册","tags":["User"]},"put":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserCreateReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserCreateRes"}}},"description":""}},"summary":"用户/注册","tags":["User"]}},"/user/detail":{"get":{"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserDetailRes"}}},"description":""}},"summary":"用户/详情","tags":["User"]},"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserDetailReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserDetailRes"}}},"description":""}},"summary":"用户/详情","tags":["User"]}},"/user/login":{"delete":{"parameters":[{"description":"用户名","in":"query","name":"nickname","required":true,"schema":{"description":"用户名","format":"string","properties":{},"type":"string"}},{"description":"密码","in":"query","name":"password","required":true,"schema":{"description":"密码","format":"string","properties":{},"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginRes"}}},"description":""}},"summary":"用户/登录","tags":["User"]},"get":{"parameters":[{"description":"用户名","in":"query","name":"nickname","required":true,"schema":{"description":"用户名","format":"string","properties":{},"type":"string"}},{"description":"密码","in":"query","name":"password","required":true,"schema":{"description":"密码","format":"string","properties":{},"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginRes"}}},"description":""}},"summary":"用户/登录","tags":["User"]},"patch":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginRes"}}},"description":""}},"summary":"用户/登录","tags":["User"]},"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginRes"}}},"description":""}},"summary":"用户/登录","tags":["User"]},"put":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLoginRes"}}},"description":""}},"summary":"用户/登录","tags":["User"]}},"/user/logout":{"get":{"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLogoutRes"}}},"description":""}},"summary":"用户/登出","tags":["User"]},"post":{"requestBody":{"required":true,"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLogoutReq"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/web.api.user.v1.UserLogoutRes"}}},"description":""}},"summary":"用户/登出","tags":["User"]}}}}`,
		"./output",
		[]*Filter{
			{
				Method: "post",
				Type:   FilterTypeInclude,
				Path:   "/user/login",
			},
			{
				Method: "post",
				Type:   FilterTypeInclude,
				Path:   "/user/detail",
			},
		},
	)
	if err != nil {
		panic(err)
	}
}
