package testWithExcel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"strings"
	"testing"
)

const (
	CaseSheetName = "case"      //用例的sheet名
	LoginPrefix   = "l:"        //登录信息的前缀
	BodyPrefix    = "b:"        //body信息的前缀
	DescKey       = "desc"      //用例描述的key
	AssignVarKey  = "assignVar" //变量设置的key
)

type TestWithExcel interface {
	Run(ctx context.Context)
}

type defaultTestWithExcel struct {
	T              *testing.T     `v:""`
	Failfast       bool           `v:""` //只要有用例失败就停止，默认true
	CaseData       CaseData       `v:"required#测试用例数据CaseData不能为空"`
	PrepareData    PrepareData    `v:""` //测试准备数据
	BeforeFunc     BeforeFunc     `v:""` //前期处理
	CaseHandleFunc CaseHandleFunc `v:"required#用例的处理方法CaseHandleFunc不能为空"`
	AfterCaseFunc  AfterCaseFunc  `v:""` //单个用例结束后的处理
	AfterFunc      AfterFunc      `v:""` //测试结束后的处理
}

type CaseInfo struct {
	Cfg       *CaseInfoCfg      //配置信息
	Body      map[string]string //body信息
	Login     map[string]string //登录信息
	Extend    map[string]string //自定义扩展参数
	Desc      string            //描述信息
	AssignVar string            //变量设置
}
type CaseInfoCfg struct {
	Name       string //用例名称
	IsOpen     bool   //是否开启
	AssertType string //用例断言类型
	Expect     string //用例期望结果
}
type CaseData []CaseInfo
type PrepareData map[string][]map[string]string
type OptionsFunc func(o *defaultTestWithExcel)
type BeforeFunc func(ctx context.Context, prepareData PrepareData)
type CaseHandleFunc func(ctx context.Context, t *testing.T, caseInfo CaseInfo) (ret interface{}, err error)
type AfterCaseFunc func(ctx context.Context, caseInfo CaseInfo, caseRet interface{}, isCasePass bool)
type AfterFunc func(ctx context.Context, prepareData PrepareData, caseData CaseData, failCase *CaseInfo)

func WithFailfast(f bool) OptionsFunc {
	return func(o *defaultTestWithExcel) {
		o.Failfast = f
	}
}
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
		Failfast:    true,
		CaseData:    caseData,
		PrepareData: prepareData,
	}

	//参数赋值
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
	var (
		cfgIndexMap = map[int]string{
			0: "name",
			1: "isOpen",
			2: "assertType",
			3: "expect",
		}
		loginIndexMap, bodyIndexMap, extendIndexMap = map[int]string{}, map[int]string{}, map[int]string{}
		descIndex, AssignVarIndex                   int
	)
	d, err := f.GetRows(CaseSheetName)
	if err != nil {
		return
	}
	//第一列为key，且前几个key是固定的，顺序也不能错
	if len(d) == 0 {
		return nil, fmt.Errorf("sheet%s的用例数据不能为空", CaseSheetName)
	}
	//检查配置的字段定义
	for k, v := range cfgIndexMap {
		if len(d[0])-1 < k || d[0][k] != v {
			return nil, fmt.Errorf("sheet%s未在第%d列定义配置字段%s", CaseSheetName, k, v)
		}
	}
	//获取各种字段的key索引
	for i := len(cfgIndexMap); i < len(d[0]); i++ {
		switch {
		case strings.HasPrefix(d[0][i], LoginPrefix):
			loginIndexMap[i], _ = strings.CutPrefix(d[0][i], LoginPrefix)
		case strings.HasPrefix(d[0][i], BodyPrefix):
			bodyIndexMap[i], _ = strings.CutPrefix(d[0][i], BodyPrefix)
		case d[0][i] == DescKey:
			descIndex = i
		case d[0][i] == AssignVarKey:
			AssignVarIndex = i
		default:
			extendIndexMap[i] = d[0][i]
		}
	}

	for i := 1; i < len(d); i++ {
		c := CaseInfo{
			Cfg:    &CaseInfoCfg{},
			Body:   map[string]string{},
			Login:  map[string]string{},
			Extend: map[string]string{},
		}
		for k, v := range d[i] {
			if k1, ok := cfgIndexMap[k]; ok {
				switch k1 {
				case "name":
					c.Cfg.Name = v
				case "isOpen":
					c.Cfg.IsOpen = v == "yes" || v == "on" || v == "true"
				case "assertType":
					c.Cfg.AssertType = v
				case "expect":
					c.Cfg.Expect = v
				}
			} else if k1, ok = loginIndexMap[k]; ok {
				c.Login[k1] = v
			} else if k1, ok = bodyIndexMap[k]; ok {
				c.Body[k1] = v
			} else if k == descIndex {
				c.Desc = v
			} else if k == AssignVarIndex {
				c.AssignVar = v
			} else if k1, ok = extendIndexMap[k]; ok {
				c.Extend[k1] = v
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
	var failCase *CaseInfo
	for _, v := range s.CaseData {
		if !v.Cfg.IsOpen {
			continue
		}
		//变量替换
		HandleVar(&v.Body)
		HandleVar(&v.Login)
		HandleVar(&v.Extend)
		failCase = nil
		gtest.C(s.T, func(t *gtest.T) {
			isCasePass := false
			//调用自定义处理方法
			ret, err := s.CaseHandleFunc(ctx, s.T, v)
			defer func() {
				//单个用例结束处理
				if s.AfterCaseFunc != nil {
					s.AfterCaseFunc(ctx, v, ret, isCasePass)
				}
				//是否失败
				if !isCasePass {
					failCase = &v
				}
			}()
			if err != nil {
				t.Errorf(`用例[%s]处理异常:%v`, v.Cfg.Name, err)
				return
			}
			//变量采集
			err = SetVarByAssignVarPattern(gconv.String(ret), v.AssignVar)
			if err != nil {
				t.Errorf(`用例[%s]处理异常:%v`, v.Cfg.Name, err)
				return
			}
			//断言
			AssertByType(v.Cfg.AssertType, v.Cfg.Name, ret, ReplayVar(v.Cfg.Expect))
			isCasePass = true
		})
		if s.Failfast && failCase != nil {
			break
		}
	}

	//3.测试后处理
	if s.AfterFunc != nil {
		s.AfterFunc(ctx, s.PrepareData, s.CaseData, failCase)
	}
}
