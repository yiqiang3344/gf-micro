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
)

type checkOpt struct {
	Pattern string                   //匹配路由
	Level   CheckLevel               //检查级别
	Kind    reflect.Kind             //类型
	Extra   map[ExtraKey]interface{} //额外检查,如 eq 指定值
}

var checkMsgMap = map[CheckLevel]string{
	MustInput:        "必填项",
	MustInputZero:    "必须为零值",
	MustInputNotZero: "必须非零值",
	ProposalNotEmpty: "建议非空",
	OptionalInput:    "选填项",
	ProposalEmpty:    "建议为空",
}

func checkRules(ctx context.Context, gcfg *gcfg.Config, rules []checkOpt) (errs []string) {
	for _, opt := range rules {
		v := gcfg.MustGet(ctx, opt.Pattern).Interface()
		errs = append(errs, checkLevel(v, opt.Pattern, opt.Level)...)
		errs = append(errs, checkKind(v, opt.Pattern, opt.Kind)...)
		errs = append(errs, checkExtra(v, opt.Pattern, opt.Kind, opt.Extra)...)
	}
	return
}

func checkExtra(v interface{}, pattern string, kind reflect.Kind, extra map[ExtraKey]interface{}) (errs []string) {
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
		switch k1 {
		case EQ:
			if v1 != v {
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "等于", v1))
			}
		case NE:
			if v1 == v {
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "不等于", v1))
			}
		case GT:
			switch {
			case kind == reflect.String && gconv.String(v) > gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) > gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) > gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) > gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) > gconv.Int(v1):
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "大于", v1))
			}
		case GE:
			switch {
			case kind == reflect.String && gconv.String(v) >= gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) >= gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) >= gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) >= gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) >= gconv.Int(v1):
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "大于等于", v1))
			}
		case LT:
			switch {
			case kind == reflect.String && gconv.String(v) < gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) < gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) < gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) < gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) < gconv.Int(v1):
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "小于", v1))
			}
		case LE:
			switch {
			case kind == reflect.String && gconv.String(v) <= gconv.String(v1):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && gconv.Int64(v) <= gconv.Int64(v1):
			case contains(floatKins, kind) && gconv.Float64(v) <= gconv.Float64(v1):
			case kind == reflect.Map && len(gconv.Map(v)) <= gconv.Int(v1):
			case kind == reflect.Slice && len(gconv.SliceAny(v)) <= gconv.Int(v1):
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "小于等于", v1))
			}
		case IN:
			switch {
			case kind == reflect.String && contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && contains(gconv.Int64s(v1), gconv.Int64(v)):
			case contains(floatKins, kind) && contains(gconv.Float64s(v1), gconv.Float64(v)):
			case kind == reflect.Map:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "被包含于", v1))
			}
		case NI:
			switch {
			case kind == reflect.String && !contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Bool:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case (contains(intKins, kind) || contains(uintKins, kind)) && !contains(gconv.Int64s(v1), gconv.Int64(v)):
			case contains(floatKins, kind) && !contains(gconv.Float64s(v1), gconv.Float64(v)):
			case kind == reflect.Map:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
			default:
				errs = append(errs, fmt.Sprintf("[错误]%s:值应该%s%v", pattern, "不被包含于", v1))
			}
		default:
			panic(fmt.Errorf("不支持的额外规则%v:%v", k1, v1))
		}
	}
	return
}

func checkKind(v interface{}, pattern string, kind reflect.Kind) (errs []string) {
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
	switch {
	case kind == reflect.String && k != reflect.String:
		errs = append(errs, fmt.Sprintf("[错误]%s应该为String类型,目前是%v", pattern, k))
	case kind == reflect.Bool && k != reflect.Bool:
		errs = append(errs, fmt.Sprintf("[错误]%s应该为bool类型,目前是%v", pattern, k))
	case contains(intKins, kind) && !contains(intKins, k):
		errs = append(errs, fmt.Sprintf("[错误]%s应该为Int类型,目前是%v", pattern, k))
	case contains(uintKins, kind) && !contains(uintKins, k):
		errs = append(errs, fmt.Sprintf("[错误]%s应该为Uint类型,目前是%v", pattern, k))
	case contains(floatKins, kind) && !contains(floatKins, k):
		errs = append(errs, fmt.Sprintf("[错误]%s应该为Float类型,目前是%v", pattern, k))
	case kind == reflect.Map && k != reflect.Map:
		errs = append(errs, fmt.Sprintf("[错误]%s应该为Map类型,目前是%v", pattern, k))
	case kind == reflect.Slice && k != reflect.Slice:
		errs = append(errs, fmt.Sprintf("[错误]%s应该为Slice类型,目前是%v", pattern, k))
	}
	return
}

func checkLevel(v interface{}, pattern string, level CheckLevel) (errs []string) {
	switch level {
	case MustInput:
		if v == nil {
			errs = append(errs, fmt.Sprintf("[错误]%s:%s", pattern, checkMsgMap[level]))
		}
	case MustInputZero:
		if v == nil || !reflect.ValueOf(v).IsZero() {
			errs = append(errs, fmt.Sprintf("[错误]%s:%s", pattern, checkMsgMap[level]))
		}
	case MustInputNotZero:
		vZero := reflect.ValueOf(v)
		vZero.SetZero()
		if v == nil || reflect.ValueOf(v).IsZero() {
			errs = append(errs, fmt.Sprintf("[错误]%s:%s", pattern, checkMsgMap[level]))
		}
	case ProposalNotEmpty:
		if v == nil || reflect.ValueOf(v).IsZero() {
			errs = append(errs, fmt.Sprintf("[建议]%s:%s", pattern, checkMsgMap[level]))
		}
	case OptionalInput:
	case ProposalEmpty:
		if v != nil && !reflect.ValueOf(v).IsZero() {
			errs = append(errs, fmt.Sprintf("[提醒]%s:%s", pattern, checkMsgMap[level]))
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
