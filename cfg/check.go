package cfg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"strings"
)

type CheckLevel int

const (
	MustInput        CheckLevel = iota
	MustInputZero    CheckLevel = iota
	MustInputNotZero CheckLevel = iota
	ProposalNotEmpty CheckLevel = iota
	OptionalInput    CheckLevel = iota
	ProposalEmpty    CheckLevel = iota
)

type ExtraKey int

const (
	EQ ExtraKey = iota
	NE ExtraKey = iota
	GT ExtraKey = iota
	GE ExtraKey = iota
	LT ExtraKey = iota
	LE ExtraKey = iota
	IN ExtraKey = iota
	NI ExtraKey = iota
	CO ExtraKey = iota
)

type ENV string

const (
	PROD ENV = "prod"
	DEV  ENV = "dev"
)

type checkOpt struct {
	Pattern string                   //匹配路由
	Env     []ENV                    //环境
	Level   CheckLevel               //检查级别
	Kind    reflect.Kind             //类型
	Extra   map[ExtraKey]interface{} //额外检查,如 eq 指定值
}

var checkMsgMap = map[CheckLevel]map[string]string{
	MustInput:        {"level": "错误", "msg": "必填"},
	MustInputZero:    {"level": "错误", "msg": "必须为零值"},
	MustInputNotZero: {"level": "错误", "msg": "必须非零值"},
	ProposalNotEmpty: {"level": "建议", "msg": "非空"},
	OptionalInput:    {"level": "提醒", "msg": "选填"},
	ProposalEmpty:    {"level": "提醒", "msg": "为空"},
}

func checkRules(ctx context.Context, gcfg *gcfg.Config, env ENV, rules []checkOpt) (errs []string) {
	for _, opt := range rules {
		if !contains(opt.Env, env) {
			continue
		}
		if ok, ret := checkRule(gcfg.MustGet(ctx, opt.Pattern).Interface(), opt); !ok {
			errs = append(errs, ret)
		}
	}
	return
}

func checkRule(v interface{}, opt checkOpt) (bool, string) {
	vStr := "nil"
	if v != nil {
		vStr = fmt.Sprintf("%T%v", v, v)
	}
	if ok := checkLevel(v, opt.Level); !ok {
		return false, fmt.Sprintf("[%s]%s值为%s,期望%s", checkMsgMap[opt.Level]["level"], opt.Pattern, vStr, checkMsgMap[opt.Level]["msg"])
	} else if ok, k := checkKind(v, opt.Kind); !ok {
		return false, fmt.Sprintf("[错误]%s类型应该为%s,目前是%T", opt.Pattern, k, v)
	} else if ok, k := checkExtra(v, opt.Kind, opt.Extra); !ok {
		return false, fmt.Sprintf("[错误]%s:值%s应该%s", opt.Pattern, vStr, k)
	}
	return true, ""
}

func checkLevel(v interface{}, level CheckLevel) bool {
	ok := true
	switch level {
	case MustInput:
		if v == nil {
			ok = false
		}
	case MustInputZero:
		if v == nil || !reflect.ValueOf(v).IsZero() {
			ok = false
		}
	case MustInputNotZero:
		if v == nil || reflect.ValueOf(v).IsZero() {
			ok = false
		}
	case ProposalNotEmpty:
		if v == nil || reflect.ValueOf(v).IsZero() {
			ok = false
		}
	case OptionalInput:
	case ProposalEmpty:
		if v != nil && !reflect.ValueOf(v).IsZero() {
			ok = false
		}
	}
	return ok
}

func checkKind(v interface{}, kind reflect.Kind) (ok bool, expectKind string) {
	ok = true
	if v == nil {
		return
	}
	//先检查类型
	kt := reflect.TypeOf(v).String()
	if kt == "json.Number" {
		if strings.Contains(v.(json.Number).String(), ".") {
			v, _ = v.(json.Number).Float64()
		} else {
			v, _ = v.(json.Number).Int64()
		}
	}
	k := reflect.ValueOf(v).Kind()
	intKins := []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	uintKins := []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}
	floatKins := []reflect.Kind{reflect.Float32, reflect.Float64}
	ok = false
	switch {
	case kind == reflect.String && k != reflect.String:
		expectKind = "String"
	case kind == reflect.Bool && k != reflect.Bool:
		expectKind = "Bool"
	case contains(intKins, kind) && !contains(intKins, k):
		expectKind = "Int"
	case contains(uintKins, kind) && (!contains(intKins, k) || gconv.Int(v) < 0):
		expectKind = "Uint"
	case contains(floatKins, kind) && !contains(floatKins, k):
		expectKind = "Float"
	case kind == reflect.Map && k != reflect.Map:
		expectKind = "Map"
	case kind == reflect.Slice && k != reflect.Slice:
		expectKind = "Slice"
	default:
		ok = true
	}
	return
}

func checkExtra(v interface{}, kind reflect.Kind, extra map[ExtraKey]interface{}) (ok bool, expect string) {
	ok = true
	if v == nil || len(extra) == 0 {
		return
	}
	//先检查类型
	kt := reflect.TypeOf(v).String()
	if kt == "json.Number" {
		if strings.Contains(v.(json.Number).String(), ".") {
			v, _ = v.(json.Number).Float64()
		} else {
			v, _ = v.(json.Number).Int64()
		}
	}
	intKins := []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	uintKins := []reflect.Kind{reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}
	floatKins := []reflect.Kind{reflect.Float32, reflect.Float64}
	for k1, v1 := range extra {
		v1Str := fmt.Sprintf("%T%v", v1, v1)
		switch k1 {
		case EQ:
			switch {
			case kind == reflect.String && gconv.String(v) == gconv.String(v1):
			case kind == reflect.Bool && gconv.Bool(v) == gconv.Bool(v1):
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) == gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) == gconv.Float64(v1):
			case kind == reflect.Map && reflect.DeepEqual(v, v1):
			case kind == reflect.Slice && reflect.DeepEqual(v, v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "等于", v1Str)
				return
			}
		case NE:
			switch {
			case kind == reflect.String && gconv.String(v) != gconv.String(v1):
			case kind == reflect.Bool && gconv.Bool(v) != gconv.Bool(v1):
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) != gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) != gconv.Float64(v1):
			case kind == reflect.Map && !reflect.DeepEqual(v, v1):
			case kind == reflect.Slice && !reflect.DeepEqual(v, v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "不等于", v1Str)
				return
			}
		case GT:
			switch {
			case kind == reflect.String && gconv.String(v) > gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) > gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) > gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) > gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) > gconv.Int(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "大于", v1Str)
				return
			}
		case GE:
			switch {
			case kind == reflect.String && gconv.String(v) >= gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) >= gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) >= gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) >= gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) >= gconv.Int(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "大于等于", v1Str)
				return
			}
		case LT:
			switch {
			case kind == reflect.String && gconv.String(v) < gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) < gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) < gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) < gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) < gconv.Int(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "小于", v1Str)
				return
			}
		case LE:
			switch {
			case kind == reflect.String && gconv.String(v) <= gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) <= gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) <= gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) <= gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) <= gconv.Int(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "小于等于", v1Str)
				return
			}
		case IN:
			switch {
			case kind == reflect.String && contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && contains(gconv.Int64s(v1), gconv.Int64(v)):
			case contains(floatKins, kind) && contains(gconv.Float64s(v1), gconv.Float64(v)):
			case kind == reflect.Map:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case kind == reflect.Slice:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "被包含于", v1Str)
				return
			}
		case NI:
			switch {
			case kind == reflect.String && !contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case (contains(intKins, kind) || contains(uintKins, kind)) && !contains(gconv.Int64s(v1), gconv.Int64(v)):
			case contains(floatKins, kind) && !contains(gconv.Float64s(v1), gconv.Float64(v)):
			case kind == reflect.Map:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case kind == reflect.Slice:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "不被包含于", v1Str)
				return
			}
		case CO:
			switch {
			case kind == reflect.String:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case contains(intKins, kind) || contains(uintKins, kind):
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case contains(floatKins, kind):
				panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
			case kind == reflect.Map && containsKey(v.(map[string]interface{}), v1.(string)):
			case kind == reflect.Slice && contains(v.([]interface{}), v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "包含", v1Str)
				return
			}
		default:
			panic(fmt.Errorf("不支持的额外规则%v:%s", k1, v1Str))
		}
	}
	return
}

// 输出错误信息
func printErr(errs []string) {
	if len(errs) == 0 {
		fmt.Println("恭喜没有问题~")
	} else {
		fmt.Println("检查结果如下:")
	}
	for _, err := range errs {
		fmt.Println(err)
	}
}

// 使用范型写一个判断是否存在某个元素的方法
func contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// 写一个判断map中是否包含指定key的函数
func containsKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
