package testWithExcel

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"github.com/yiqiang3344/gf-micro/utility"
	"strings"
)

type Property struct {
	Name     string `json:"name"`
	Format   string `json:"format"`
	Desc     string `json:"desc"`
	Required bool   `json:"required"`
}

type Case struct {
	Name       string            `json:"name"`
	IsOpen     string            `json:"isOpen"`
	AssertType string            `json:"assertType"`
	Expect     string            `json:"expect"`
	LoginAttrs []string          `json:"loginAttrs"`
	ReqAttrs   []string          `json:"reqAttrs"`
	ReqVals    map[string]string `json:"reqVals"`
	AssignVar  string            `json:"assignVar"`
	Desc       string            `json:"desc"`
	Method     string            `json:"method"`
	Path       string            `json:"path"`
}

type Path struct {
	Path    string      `json:"path"`
	Method  string      `json:"method"`
	Summary string      `json:"summary"`
	Req     []*Property `json:"req"`
}

type excelHead struct {
	Name  string  `json:"name"`
	Width float64 `json:"width"`
}

type FilterType string

const (
	FilterTypeInclude FilterType = "include" //包含
	FilterTypeExpel   FilterType = "expel"   //排出
)

type Filter struct {
	Path   string
	Method string
	Type   FilterType
}

// GenTestCaseExcelByOpenApiJson 根据openApi的json文件生成接口测试用例excel
func GenTestCaseExcelByOpenApiJson(openApiJsonStr string, outputDir string, filter ...[]*Filter) error {
	//初始化过滤条件
	filterIncludeMap := map[string]int{}
	filterExpelMap := map[string]int{}
	if len(filter) > 0 {
		for _, v := range filter[0] {
			switch v.Type {
			case FilterTypeInclude:
				filterIncludeMap[v.Method+"-"+v.Path] = 1
			case FilterTypeExpel:
				filterExpelMap[v.Method+"-"+v.Path] = 1
			}
		}
	}

	//根据openApi的json文件生成接口测试用例excel
	j, err := gjson.DecodeToJson(openApiJsonStr)
	if err != nil {
		return err
	}

	var paths []*Path
	//遍历paths
	for k, v := range j.Get("paths").Map() {
		schemas := j.Get("components.schemas").Map()
		for k1, v1 := range v.(map[string]interface{}) {
			path := new(Path)
			path.Path = k
			path.Method = k1
			path.Summary = v1.(map[string]interface{})["summary"].(string)

			//按过滤条件过滤
			if len(filterIncludeMap) > 0 {
				if _, ok := filterIncludeMap[path.Method+"-"+path.Path]; !ok {
					g.Dump("不包含", path.Method+"-"+path.Path)
					continue
				}
			}
			if len(filterExpelMap) > 0 {
				if _, ok := filterExpelMap[path.Method+"-"+path.Path]; ok {
					g.Dump("过滤", path.Method+"-"+path.Path)
					continue
				}
			}

			switch path.Method {
			case "post", "put", "patch":
				if v1.(map[string]interface{})["requestBody"] != nil {
					rS, _ := strings.CutPrefix(v1.(map[string]interface{})["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})["schema"].(map[string]interface{})["$ref"].(string), "#/components/schemas/")
					schema := schemas[rS].(map[string]interface{})
					requiredArr := gconv.Strings(schema["required"])
					for k2, v2 := range schema["properties"].(map[string]interface{}) {
						path.Req = append(path.Req, &Property{
							Name:     k2,
							Format:   v2.(map[string]interface{})["type"].(string),
							Desc:     v2.(map[string]interface{})["description"].(string),
							Required: gstr.InArray(requiredArr, k2),
						})
					}
				}
			case "get", "delete":
				if v1.(map[string]interface{})["parameters"] != nil {
					for _, v2 := range v1.(map[string]interface{})["parameters"].([]interface{}) {
						path.Req = append(path.Req, &Property{
							Name:     v2.(map[string]interface{})["name"].(string),
							Format:   v2.(map[string]interface{})["schema"].(map[string]interface{})["format"].(string),
							Desc:     v2.(map[string]interface{})["description"].(string),
							Required: v2.(map[string]interface{})["required"].(bool),
						})
					}
				}
			}
			// 对参数排序，起码每次都是相同的
			var arr []*Property
			for _, v2 := range garray.NewSortedArrayFrom(gconv.Interfaces(path.Req), func(a, b interface{}) int {
				if a.(*Property).Name > b.(*Property).Name {
					return 1
				} else {
					return -1
				}
			}).Slice() {
				arr = append(arr, v2.(*Property))
			}
			path.Req = arr
			paths = append(paths, path)
		}
	}
	//根据paths生成cases
	caseMap := gmap.NewListMap()
	for _, path := range paths {
		var cases []*Case
		case1 := new(Case)
		case1.Name = path.Summary
		case1.IsOpen = "yes"
		case1.AssertType = AssertTypeEQ
		case1.Expect = "请根据实际情况输入"
		case1.LoginAttrs = []string{"needLogin", "token"}
		case1.Desc = path.Summary
		case1.Method = path.Method
		case1.Path = path.Path
		var requiredArr []string
		for _, v := range path.Req {
			case1.ReqAttrs = append(case1.ReqAttrs, v.Name)
			if v.Required {
				requiredArr = append(requiredArr, v.Name)
			}
		}
		//各种可能的组合
		var combinations [][]string
		generateCombinations(requiredArr, 0, []string{}, &combinations)
		for k, combination := range combinations {
			case1.Name = fmt.Sprintf("%s%d", path.Summary, k+1)
			if len(combination) == 0 {
				case1.Desc = "参数为空"
			} else {
				case1.Desc = fmt.Sprintf("传参:%s", strings.Join(combination, ","))
				case1.ReqVals = map[string]string{}
				for _, v := range combination {
					case1.ReqVals[v] = "请根据实际情况输入"
				}
			}
			c := *case1
			cases = append(cases, &c)
		}
		caseMap.Set(path.Summary+"-"+path.Method, cases)
	}
	//生成excel文件
	for k, v := range caseMap.Map() {
		if len(v.([]*Case)) == 0 {
			continue
		}
		f := excelize.NewFile()
		// 创建一个工作表
		excelFilepath := fmt.Sprintf("%s/%s.xlsx", outputDir, strings.ReplaceAll(k.(string), "/", "-"))
		sheet1Head := []excelHead{
			{Name: "name", Width: 15},
			{Name: "isOpen", Width: 6},
			{Name: "assertType", Width: 9},
			{Name: "expect", Width: 16},
		}
		for _, v1 := range v.([]*Case)[0].LoginAttrs {
			sheet1Head = append(sheet1Head, excelHead{Name: "l:" + v1, Width: float64(len(v1)) + 1})
		}
		for _, v1 := range v.([]*Case)[0].ReqAttrs {
			sheet1Head = append(sheet1Head, excelHead{Name: "b:" + v1, Width: 16})
		}
		sheet1Head = append(sheet1Head, []excelHead{
			{Name: "desc", Width: 20},
			{Name: "assignVar", Width: 9},
			{Name: "method", Width: 7},
			{Name: "path", Width: 15},
		}...)
		sheet1Name := "case"
		sheet1Index, err1 := f.NewSheet(sheet1Name)
		if err1 != nil {
			return err1
		}
		// 设置头部及列宽
		for k1, v1 := range sheet1Head {
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(1, k1+1), v1.Name)
			f.SetColWidth(sheet1Name, utility.ConvertNumToExcelCol(k1+1), utility.ConvertNumToExcelCol(k1+1), v1.Width)
		}
		// 设置值
		n := 2
		for _, v1 := range v.([]*Case) {
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 1), v1.Name)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 2), v1.IsOpen)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 3), v1.AssertType)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 4), v1.Expect)
			for k2, _ := range v1.LoginAttrs {
				f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 5+k2), "")
			}
			for k2, v2 := range v1.ReqAttrs {
				vT := ""
				if v3, ok := v1.ReqVals[v2]; ok {
					vT = v3
				}
				f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, 5+len(v1.LoginAttrs)+k2), vT)
			}
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, len(v1.LoginAttrs)+len(v1.ReqAttrs)+5), v1.Desc)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, len(v1.LoginAttrs)+len(v1.ReqAttrs)+6), v1.AssignVar)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, len(v1.LoginAttrs)+len(v1.ReqAttrs)+7), v1.Method)
			f.SetCellValue(sheet1Name, utility.ConvertToExcelPos(n, len(v1.LoginAttrs)+len(v1.ReqAttrs)+8), v1.Path)
			n++
		}
		// 设置工作簿的默认工作表
		f.SetActiveSheet(sheet1Index)
		f.DeleteSheet("Sheet1")
		// 根据指定路径保存文件
		if err2 := f.SaveAs(excelFilepath); err2 != nil {
			return err2
		}
	}
	return nil
}

// generateCombinations 生成字符串切片的所有组合
func generateCombinations(input []string, index int, current []string, result *[][]string) {
	// 递归终止条件：到达输入切片的末尾
	if index == len(input) {
		combination := make([]string, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	// 不选择当前元素，直接递归到下一个元素
	generateCombinations(input, index+1, current, result)

	// 选择当前元素，加入到当前组合中，然后递归到下一个元素
	next := make([]string, len(current)+1)
	copy(next, current)
	next[len(current)] = input[index]
	generateCombinations(input, index+1, next, result)
}
