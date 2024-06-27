package testFrame

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/xuri/excelize/v2"
	"strings"
	"testing"
	"web/test"
)

const (
	CaseSheetName = "case" //用例的sheet名
	LoginPrefix   = "l:"   //登录信息的前缀
	BodyPrefix    = "b:"   //body信息的前缀
	DescKey       = "desc" //用例描述的key
)

type TestWithExcel interface {
	Run(ctx context.Context)
}

type defaultTestWithExcel struct {
	T              *testing.T     `v:""`
	CaseData       CaseData       `v:"required#测试用例数据不能为空"`
	PrepareData    PrepareData    `v:""` //测试准备数据
	BeforeFunc     BeforeFunc     `v:""` //前期处理
	CaseHandleFunc CaseHandleFunc `v:"required#用例的处理方法不能为空"`
	AfterCaseFunc  AfterCaseFunc  `v:""` //单个用例结束后的处理
	AfterFunc      AfterFunc      `v:""` //测试结束后的处理
}

type CaseInfo struct {
	Cfg   map[string]string //配置信息
	Body  map[string]string //body信息
	Login map[string]string //登录信息
	Desc  string            //描述信息
}
type CaseData []CaseInfo
type PrepareData map[string][]map[string]string
type OptionsFunc func(o *defaultTestWithExcel)
type BeforeFunc func(ctx context.Context, prepareData PrepareData)
type CaseHandleFunc func(ctx context.Context, caseInfo CaseInfo) (ret interface{}, err error)
type AfterCaseFunc func(ctx context.Context, caseInfo CaseInfo, caseRet interface{}, isCasePass bool)
type AfterFunc func(ctx context.Context, prepareData PrepareData, caseData CaseData)

func WithBeforeFunc(f BeforeFunc) OptionsFunc {
	return func(o *defaultTestWithExcel) {
		o.BeforeFunc = f
	}
}
func WithCaseHandleFunc(f CaseHandleFunc) OptionsFunc {
	return func(o *defaultTestWithExcel) {
		o.CaseHandleFunc = f
	}
}
func WithAfterCaseFunc(f AfterCaseFunc) OptionsFunc {
	return func(o *defaultTestWithExcel) {
		o.AfterCaseFunc = f
	}
}
func WithAfterFunc(f AfterFunc) OptionsFunc {
	return func(o *defaultTestWithExcel) {
		o.AfterFunc = f
	}
}

func New(t *testing.T, testDataFile string, funcs ...OptionsFunc) (ret TestWithExcel, err error) {
	var ctx = context.Background()

	if t == nil {
		return nil, fmt.Errorf("t不能为空")
	}

	//获取excel对象
	f, err := excelize.OpenFile(testDataFile)
	if err != nil {
		return
	}
	//处理excel数据，提取测试用例数据和测试准备数据 todo
	caseData, err := parseCaseData(f)
	if err != nil {
		return
	}
	prepareData, err := parsePrepareData(f)
	if err != nil {
		return
	}

	d := &defaultTestWithExcel{
		T:           t,
		CaseData:    caseData,
		PrepareData: prepareData,
	}

	//方法赋值
	for _, v := range funcs {
		v(d)
	}

	//检查必填参数
	if err = g.Validator().Data(d).Run(ctx); err != nil {
		return
	}

	return d, nil
}

func parseCaseData(f *excelize.File) (ret CaseData, err error) {
	d, err := f.GetRows(CaseSheetName)
	if err != nil {
		return
	}
	//第一列为key，且前几个key是固定的，顺序也不能错
	if len(d) == 0 {
		return nil, fmt.Errorf("sheet%s的用例数据不能为空", CaseSheetName)
	}
	//配置的字段定义
	cfg := map[int]string{
		0: "name",
		1: "isOpen",
		2: "needDelete",
		3: "assertType",
		4: "expect",
	}
	for k, v := range cfg {
		if len(d[0])-1 < k || d[0][k] != v {
			return nil, fmt.Errorf("sheet%s未在第%d列定义配置字段%s", CaseSheetName, k, v)
		}
	}
	if d[0][len(d[0])-1] != DescKey {
		return nil, fmt.Errorf("sheet%s未在最后一列定义用例描述字段%s", CaseSheetName, DescKey)
	}
	login, body, descIndex := map[int]string{}, map[int]string{}, len(d[0])-1
	for i := len(cfg); i < len(d[0]); i++ {
		if strings.HasPrefix(d[0][i], LoginPrefix) {
			login[i], _ = strings.CutPrefix(d[0][i], LoginPrefix)
		}
		if strings.HasPrefix(d[0][i], BodyPrefix) {
			body[i], _ = strings.CutPrefix(d[0][i], BodyPrefix)
		}
	}

	for i := 1; i < len(d); i++ {
		c := CaseInfo{
			Cfg:   map[string]string{},
			Body:  map[string]string{},
			Login: map[string]string{},
		}
		for k, v := range d[i] {
			if k1, ok := cfg[k]; ok {
				c.Cfg[k1] = v
			}
			if k1, ok := login[k]; ok {
				c.Login[k1] = v
			}
			if k1, ok := body[k]; ok {
				c.Body[k1] = v
			}
			if k == descIndex {
				c.Desc = v
			}
		}
		ret = append(ret, c)
	}
	return
}

func parsePrepareData(f *excelize.File) (ret PrepareData, err error) {
	ret = make(PrepareData)
	l := f.GetSheetList()
	for _, v := range l {
		if v == CaseSheetName {
			continue
		}
		r, err1 := f.GetRows(v)
		if err1 != nil {
			return nil, err1
		}
		if len(r) == 0 {
			continue
		}
		header := map[int]string{}
		for k1, v1 := range r[0] {
			header[k1] = v1
		}
		var sheetRows []map[string]string
		for i := 1; i < len(r); i++ {
			d := map[string]string{}
			for k1, v1 := range r[i] {
				if v2, ok := header[k1]; ok {
					d[v2] = v1
				}
			}
			sheetRows = append(sheetRows, d)
		}
		ret[v] = sheetRows
	}
	return
}

func (s *defaultTestWithExcel) Run(ctx context.Context) {
	//1.前期准备
	if s.BeforeFunc != nil {
		s.BeforeFunc(ctx, s.PrepareData)
	}

	//2.处理用例
	for _, v := range s.CaseData {
		caseName := v.Cfg["name"]              //用例名称
		caseIsOpen := v.Cfg["isOpen"] == "yes" //用例是否开启
		assertType := v.Cfg["assertType"]      //用例断言方式
		expect := v.Cfg["expect"]              //用例期望结果

		if !caseIsOpen {
			continue
		}
		gtest.C(s.T, func(t *gtest.T) {
			isCasePass := false
			ret, err := s.CaseHandleFunc(ctx, v)
			defer func() {
				//单个用例结束处理
				if s.AfterCaseFunc != nil {
					s.AfterCaseFunc(ctx, v, ret, isCasePass)
				}
			}()
			if err != nil {
				t.Errorf(`用例[%s]处理异常:%v`, caseName, err)
				return
			}
			//待完善结果处理 todo
			if assertType == "eq" {
				test.Assert(caseName, ret, expect)
			} else if assertType == "status" {
				if j, err1 := gjson.DecodeToJson(ret); err1 != nil {
					t.Errorf(`用例[%s]json解析失败:%v`, caseName, err.Error())
					return
				} else {
					test.Assert(caseName, j.Get("status").String(), expect)
				}
			} else {
				t.Errorf(`用例[%s]异常的断言类型:%s`, caseName, assertType)
				return
			}
			isCasePass = true
		})
	}

	//3.测试后处理
	if s.AfterFunc != nil {
		s.AfterFunc(ctx, s.PrepareData, s.CaseData)
	}
}
