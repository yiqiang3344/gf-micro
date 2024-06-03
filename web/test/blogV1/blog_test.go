package blogV1

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
	"testing"
	v1 "web/api/blog/v1"
	"web/test"
)

var (
	testDataFile = "./testdata/blog_test.xlsx"
)

func TestBlogCreate(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogCreate")
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
		data := v1.BlogCreateReq{
			Title:   row[5],
			Content: row[6],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/create", data)
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

func TestBlogEdit(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogEdit")
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
		data := v1.BlogEditReq{
			Id:      row[5],
			Title:   row[6],
			Content: row[7],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/edit", data)
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

func TestBlogDetail(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogDetail")
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
		data := v1.BlogDetailReq{
			Id: row[5],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/detail", data)
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

func TestBlogList(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogList")
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
		data := v1.BlogDetailReq{}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/list", data)
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

func TestBlogDelete(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogDelete")
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
		data := v1.BlogDeleteReq{
			Id: row[5],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/delete", data)
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

func TestBlogBatDelete(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogBatDelete")
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
		data := v1.BlogBatDeleteReq{
			Ids: row[5],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/bat-delete", data)
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

func TestBlogGetBatDeleteStatus(t *testing.T) {
	var (
		ctx           = context.Background()
		client        = test.GetClient(ctx, "")
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
	testData, err = f.GetRows("blogGetBatDeleteStatus")
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
		data := v1.BlogGetBatDeleteStatusReq{
			BatNo: row[5],
		}
		gtest.C(t, func(t *gtest.T) {
			ret := client.PostContent(ctx, "/blog/get-bat-delete-status", data)
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
