package t1userV1

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
	"testing"
	v1 "web/api/user/v1"
	"web/test"
)

var (
	testDataFile = "./testdata/user_test.xlsx"
)

func TestUserCreate(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx)
		err           error
		testData      [][]string
		caseName      string
		assertType    string
		expect        string
		needDelete    bool
		testCaseValid bool
	)

	//1.数据准备
	//获取excel对象
	f, err := excelize.OpenFile(testDataFile)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//2.测试用例
	testData, err = f.GetRows("userCreate")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for i, row := range testData {
		if i == 0 {
			//第一行是行头，不是数据
			continue
		}
		caseName = row[0]               //用例名称
		testCaseValid = row[1] == "yes" //测试用例是否开启
		needDelete = row[2] == "yes"    //是否需要删除测试数据
		assertType = row[3]             //测试用例断言方式
		expect = row[4]                 //测试用例期望结果

		if !testCaseValid {
			continue
		}
		data := v1.UserCreateReq{
			Nickname: row[5],
			Password: row[6],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/user/create", data)
			if assertType == "eq" {
				test.Assert(caseName, ret, expect)
			} else if assertType == "status" {
				if j, err := gjson.DecodeToJson(ret); err != nil {
					t.Errorf(`%+v json解析失败:%+v`, caseName, err.Error())
				} else {
					test.Assert(caseName, j.Get("status").String(), expect)
				}
			} else {
				t.Errorf(`%+v 异常的断言类型:%+v`, caseName, assertType)
			}
		})
		//是否需要删除测试数据
		if needDelete {
			//删除测试数据
		}
	}

	//3.清除准备数据
}

func TestUserLogin(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx)
		err           error
		testData      [][]string
		caseName      string
		assertType    string
		expect        string
		needDelete    bool
		testCaseValid bool
	)

	//1.数据准备
	//获取excel对象
	f, err := excelize.OpenFile(testDataFile)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//2.测试用例
	testData, err = f.GetRows("userLogin")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for i, row := range testData {
		if i == 0 {
			//第一行是行头，不是数据
			continue
		}
		caseName = row[0]               //用例名称
		testCaseValid = row[1] == "yes" //测试用例是否开启
		needDelete = row[2] == "yes"    //是否需要删除测试数据
		assertType = row[3]             //测试用例断言方式
		expect = row[4]                 //测试用例期望结果

		if !testCaseValid {
			continue
		}
		data := v1.UserLoginReq{
			Nickname: row[5],
			Password: row[6],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/user/login", data)
			if assertType == "eq" {
				test.Assert(caseName, ret, expect)
			} else if assertType == "code" {
				if j, err := gjson.DecodeToJson(ret); err != nil {
					t.Errorf(`%+v json解析失败:%+v`, caseName, err.Error())
				} else {
					test.Assert(caseName, j.Get("code").String(), expect)
				}
			} else {
				t.Errorf(`%+v 异常的断言类型:%+v`, caseName, assertType)
			}
		})
		//是否需要删除测试数据
		if needDelete {
			//删除测试数据
		}
	}

	//3.清除准备数据
}

func TestUserDetail(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx)
		err           error
		testData      [][]string
		caseName      string
		assertType    string
		expect        string
		needDelete    bool
		testCaseValid bool
	)

	//1.数据准备
	//获取excel对象
	f, err := excelize.OpenFile(testDataFile)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//2.测试用例
	testData, err = f.GetRows("userDetail")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for i, row := range testData {
		if i == 0 {
			//第一行是行头，不是数据
			continue
		}
		caseName = row[0]               //用例名称
		testCaseValid = row[1] == "yes" //测试用例是否开启
		needDelete = row[2] == "yes"    //是否需要删除测试数据
		assertType = row[3]             //测试用例断言方式
		expect = row[4]                 //测试用例期望结果

		if !testCaseValid {
			continue
		}
		data := v1.UserDetailReq{}
		gtest.C(t, func(t *gtest.T) {
			//是否登录
			if row[5] == "yes" {
				err := test.Login(ctx, row[6], row[7], client)
				if err != nil {
					t.Errorf(`%+v 登录失败:%+v`, caseName, err.Error())
					return
				}
			}

			ret := client.PostContent(ctx, "/user/detail", data)
			if assertType == "eq" {
				test.Assert(caseName, ret, expect)
			} else if assertType == "code" {
				if j, err := gjson.DecodeToJson(ret); err != nil {
					t.Errorf(`%+v json解析失败:%+v`, caseName, err.Error())
					return
				} else {
					test.Assert(caseName, j.Get("code").String(), expect)
				}
			} else {
				t.Errorf(`%+v 异常的断言类型:%+v`, caseName, assertType)
				return
			}
			client.SetHeader("token", "")
		})
		//是否需要删除测试数据
		if needDelete {
			//删除测试数据
		}
	}

	//3.清除准备数据
}

func TestUserLogout(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx)
		err           error
		testData      [][]string
		caseName      string
		assertType    string
		expect        string
		needDelete    bool
		testCaseValid bool
	)

	//1.数据准备
	//获取excel对象
	f, err := excelize.OpenFile(testDataFile)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//2.测试用例
	testData, err = f.GetRows("userLogout")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	for i, row := range testData {
		if i == 0 {
			//第一行是行头，不是数据
			continue
		}
		caseName = row[0]               //用例名称
		testCaseValid = row[1] == "yes" //测试用例是否开启
		needDelete = row[2] == "yes"    //是否需要删除测试数据
		assertType = row[3]             //测试用例断言方式
		expect = row[4]                 //测试用例期望结果

		if !testCaseValid {
			continue
		}
		data := v1.UserDetailReq{}
		gtest.C(t, func(t *gtest.T) {
			//是否登录
			if row[5] == "yes" {
				err := test.Login(ctx, row[6], row[7], client)
				if err != nil {
					t.Errorf(`%+v 登录失败:%+v`, caseName, err.Error())
					return
				}
			}

			ret := client.PostContent(ctx, "/user/logout", data)
			if assertType == "eq" {
				test.Assert(caseName, ret, expect)
			} else if assertType == "code" {
				if j, err := gjson.DecodeToJson(ret); err != nil {
					t.Errorf(`%+v json解析失败:%+v`, caseName, err.Error())
					return
				} else {
					test.Assert(caseName, j.Get("code").String(), expect)
				}
			} else {
				t.Errorf(`%+v 异常的断言类型:%+v`, caseName, assertType)
				return
			}
			client.SetHeader("token", "")
		})
		//是否需要删除测试数据
		if needDelete {
			//删除测试数据
		}
	}

	//3.清除准备数据
}
