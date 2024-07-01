package testWithExcel

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"regexp"
	"strings"
)

const (
	VarRegPatten = `\$\{(\w+)\}`
)

var (
	varMap = map[string]string{}
)

func GetVarMap() map[string]string {
	return varMap
}

func IsVar(key string) bool {
	reg, _ := regexp.Compile(VarRegPatten)
	return reg.MatchString(key)
}

func GetVar(key string) string {
	reg, _ := regexp.Compile(VarRegPatten)
	ret := reg.FindAllStringSubmatch(key, 1)
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
	reg, _ := regexp.Compile(VarRegPatten)
	ret := reg.ReplaceAllStringFunc(str, func(s string) string {
		return GetVar(s)
	})
	return ret
}

func SetVarByAssignVarPattern(jsonStr string, assignVar string) error {
	if strings.Trim(assignVar, " ") == "" {
		return nil
	}
	for _, v := range strings.Split(assignVar, ",") {
		r := strings.Split(v, ":")
		if len(r) != 2 {
			return fmt.Errorf("变量赋值标识异常:%s", assignVar)
		}
		if j, err := gjson.DecodeToJson(jsonStr); err == nil {
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
