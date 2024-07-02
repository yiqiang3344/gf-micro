package testWithExcel

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"regexp"
	"strings"
)

const (
	VarRegPattern = `\$\{(\w+)\}`
)

var (
	varMap = map[string]string{}
)

func GetVarMap() map[string]string {
	return varMap
}

// IsVar 检查给定的字符串是否符合变量的命名规则。
func IsVar(key string) bool {
	reg, _ := regexp.Compile(VarRegPattern)
	return reg.MatchString(key)
}

func GetVar(varStr string) string {
	reg, _ := regexp.Compile(VarRegPattern)
	ret := reg.FindAllStringSubmatch(varStr, 1)
	if len(ret) == 0 {
		return ""
	}
	if v, ok := varMap[ret[0][1]]; ok {
		return v
	}
	return ""
}

func SetVar(key string, value string) {
	varMap[key] = value
}

func ReplayVar(str string) string {
	reg, _ := regexp.Compile(VarRegPattern)
	ret := reg.ReplaceAllStringFunc(str, func(s string) string {
		return GetVar(s)
	})
	return ret
}

// SetVarByAssignVarPattern 根据指定的变量赋值模式，解析JSON字符串并设置变量。
// jsonStr 是待解析的JSON字符串。
// assignVar 是变量赋值模式，格式为"变量名1:JSON路径1,变量名2:JSON路径2"。
// 返回错误如果赋值模式不正确或JSON解析失败。
func SetVarByAssignVarPattern(jsonStr string, assignVar string) error {
	// 移除赋值模式中的空格并检查是否为空，如果为空则无需处理。
	if strings.Trim(assignVar, " ") == "" {
		return nil
	}
	// 分割赋值模式为变量名和JSON路径的组合。
	for _, v := range strings.Split(assignVar, ",") {
		r := strings.Split(v, ":")
		// 检查每个变量名和JSON路径是否正确分割，确保格式正确。
		if len(r) != 2 {
			return fmt.Errorf("变量赋值标识异常:%s", assignVar)
		}
		// 尝试将JSON字符串解析为gjson格式，以便提取指定路径的值。
		if j, err := gjson.DecodeToJson(jsonStr); err == nil {
			// 使用解析的JSON值设置变量。
			SetVar(r[0], j.Get(r[1]).String())
		} else {
			return fmt.Errorf("json[%s]解析异常:%v", jsonStr, err)
		}
	}
	return nil
}

func HandleVar(m *map[string]string) {
	m1 := *m
	for k, v := range m1 {
		if !IsVar(v) {
			continue
		}
		m1[k] = ReplayVar(v)
	}
	m = &m1
}
