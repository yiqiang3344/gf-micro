package cfg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"reflect"
	"strings"
)

type CheckLevel string

const (
	MustInput        CheckLevel = "MustInput"
	MustInputZero    CheckLevel = "MustInputZero"
	MustInputNotZero CheckLevel = "MustInputNotZero"
	ProposalNotZero  CheckLevel = "ProposalNotZero"
	OptionalInput    CheckLevel = "OptionalInput"
	ProposalZero     CheckLevel = "ProposalZero"
)

var Levels = []CheckLevel{
	MustInput,
	MustInputZero,
	MustInputNotZero,
	ProposalNotZero,
	OptionalInput,
	ProposalZero,
}

var Kinds = []reflect.Kind{
	reflect.Bool,
	reflect.Int64,
	reflect.Float64,
	reflect.String,
	reflect.Map,
	reflect.Slice,
}

type ExtraKey string

const (
	EQ ExtraKey = "EQ"
	NE ExtraKey = "NE"
	GT ExtraKey = "GT"
	GE ExtraKey = "GE"
	LT ExtraKey = "LT"
	LE ExtraKey = "LE"
	IN ExtraKey = "IN"
	NI ExtraKey = "NI"
	CO ExtraKey = "CO"
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
	ProposalNotZero:  {"level": "建议", "msg": "非零值"},
	OptionalInput:    {"level": "提醒", "msg": "选填"},
	ProposalZero:     {"level": "提醒", "msg": "为零值"},
}

func Check(ctx context.Context, parser *gcmd.Parser, apolloIp string) {
	var (
		errs  []string
		env   = DEV
		rules []checkOpt
	)

	//环境配置
	envCfgV := parser.GetOpt("env")
	if !envCfgV.IsNil() {
		env = ENV(envCfgV.String())
	}
	if !contains([]ENV{PROD, DEV}, env) {
		panic("env参数错误")
	}

	//判断是否检查apollo配置
	if !parser.GetOpt("all").IsNil() || !parser.GetOpt("apollo").IsNil() {
		errs = append(errs, checkRules(ctx, gcfg.Instance(), env, apolloRules)...)
	}

	//先判断是否有apollo配置，有的话接入apollo
	if !gcfg.Instance().MustGet(ctx, "apollo").IsNil() {
		adapter, err := gcfg_apollo.CreateAdapterApollo(ctx, apolloIp)
		if err != nil {
			panic(err)
		}
		gcfg.Instance().SetAdapter(adapter)
	}

	//准备通用规则
	rules = append(rules, commonRules...)
	ruleMap := getRuleListMap()
	//是否有指定要检查的配置
	noAppoint := true
	for k, _ := range parser.GetOptAll() {
		if !contains([]string{"all", "env", "help"}, k) {
			noAppoint = false
			break
		}
	}
	if !parser.GetOpt("all").IsNil() {
		//准备所有规则
		ruleMap.Iterator(func(key, value interface{}) bool {
			rules = append(rules, value.([]checkOpt)...)
			return true
		})
	} else if noAppoint {
		//对已有配置准备对应规则
		ruleMap.Iterator(func(key, value interface{}) bool {
			if !gcfg.Instance().MustGet(ctx, key.(string)).IsNil() {
				rules = append(rules, value.([]checkOpt)...)
			}
			return true
		})
	} else {
		//准备指定规则
		for k, _ := range parser.GetOptAll() {
			if r := ruleMap.Get(k); r != nil {
				rules = append(rules, r.([]checkOpt)...)
			}
		}
	}

	//按规则检查配置
	errs = append(errs, checkRules(ctx, gcfg.Instance(), env, rules)...)

	//打印结果
	printErr(errs)
	return
}

func checkRules(ctx context.Context, gcfg *gcfg.Config, env ENV, rules []checkOpt) (errs []string) {
	for _, opt := range rules {
		if ok, ret := checkRule(gcfg.MustGet(ctx, opt.Pattern).Interface(), env, opt); !ok {
			errs = append(errs, ret)
		}
	}
	return
}

func checkRule(v interface{}, env ENV, opt checkOpt) (bool, string) {
	//判断规则是否合理
	if !contains(Kinds, opt.Kind) {
		//配置文件只需要用到几种类型
		panic(fmt.Sprintf("规则配置异常，不支持的Kind:%v", opt.Kind))
	}
	if !contains(Levels, opt.Level) {
		panic(fmt.Sprintf("规则配置异常，不支持的Level:%v", opt.Level))
	}

	if !contains(opt.Env, env) {
		return true, ""
	}
	//如果是json.Number则转int64或float64
	if v1, ok := v.(json.Number); ok {
		if strings.Contains(v1.String(), ".") {
			v, _ = v1.Float64()
		} else {
			v, _ = v1.Int64()
		}
	}
	vStr := "nil"
	if v != nil {
		vStr = fmt.Sprintf("%T%v", v, v)
	}
	if ok, k := checkKind(v, opt.Kind); !ok {
		return false, fmt.Sprintf("[错误]%s类型应该为%s,目前是%T", opt.Pattern, k, v)
	} else if ok := checkLevel(v, opt.Level, opt.Kind); !ok {
		return false, fmt.Sprintf("[%s]%s值为%s,期望%s", checkMsgMap[opt.Level]["level"], opt.Pattern, vStr, checkMsgMap[opt.Level]["msg"])
	} else if ok, k := checkExtra(v, opt.Kind, opt.Extra); !ok {
		return false, fmt.Sprintf("[错误]%s:值%s应该%s", opt.Pattern, vStr, k)
	}
	return true, ""
}

func checkLevel(v interface{}, level CheckLevel, kind reflect.Kind) bool {
	if !contains(Kinds, kind) {
		//配置文件只需要用到几种类型
		panic(fmt.Sprintf("规则配置异常，不支持的Kind:%v", kind))
	}
	if !contains(Levels, level) {
		panic(fmt.Sprintf("规则配置异常，不支持的Level:%v", level))
	}

	ok := true
	switch level {
	case MustInput:
		if v == nil {
			ok = false
		}
	case MustInputZero, ProposalZero:
		switch kind {
		case reflect.Map, reflect.Slice:
			if v != nil && reflect.ValueOf(v).Len() != 0 {
				ok = false
			}
		default:
			if v == nil || !reflect.ValueOf(v).IsZero() {
				ok = false
			}
		}
	case MustInputNotZero, ProposalNotZero:
		switch kind {
		case reflect.Map, reflect.Slice:
			if v == nil || reflect.ValueOf(v).Len() == 0 {
				ok = false
			}
		default:
			if v == nil || reflect.ValueOf(v).IsZero() {
				ok = false
			}
		}
	case OptionalInput:
	}
	return ok
}

func checkKind(v interface{}, kind reflect.Kind) (ok bool, expectKind string) {
	if !contains(Kinds, kind) {
		//配置文件只需要用到几种类型
		panic(fmt.Sprintf("规则配置异常，不支持的Kind:%v", kind))
	}

	ok = true
	if v == nil {
		return
	}
	k := reflect.ValueOf(v).Kind()
	intKins := []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
	floatKins := []reflect.Kind{reflect.Float32, reflect.Float64}
	ok = false
	switch {
	case kind == reflect.String && k != reflect.String:
		expectKind = "String"
	case kind == reflect.Bool && k != reflect.Bool:
		expectKind = "Bool"
	case kind == reflect.Int64 && !contains(intKins, k):
		expectKind = "Int"
	case kind == reflect.Float64 && !contains(floatKins, k) && k != reflect.Int64: //整数float64会转为int64
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
	if !contains(Kinds, kind) {
		//配置文件只需要用到几种类型
		panic(fmt.Sprintf("规则配置异常，不支持的Kind:%v", kind))
	}

	ok = true
	if len(extra) == 0 {
		return
	}

	//先检查类型
	for k1, v1 := range extra {
		v1Str := fmt.Sprintf("%T%v", v1, v1)
		switch k1 {
		case EQ:
			switch {
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if (v == nil && v1 != nil) || (v != nil && v1 == nil) {
				ok = false
				expect = fmt.Sprintf("%s%s", "等于", v1Str)
				return
			}

			switch {
			case kind == reflect.String && gconv.String(v) == gconv.String(v1):
			case kind == reflect.Bool && gconv.Bool(v) == gconv.Bool(v1):
			case kind == reflect.Int64 && gconv.Int64(v) == gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) == gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "等于", v1Str)
				return
			}
		case NE:
			switch {
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil && v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "不等于", v1Str)
				return
			} else if v != nil && v1 == nil {
				return
			}

			switch {
			case kind == reflect.String && gconv.String(v) != gconv.String(v1):
			case kind == reflect.Bool && gconv.Bool(v) != gconv.Bool(v1):
			case kind == reflect.Int64 && gconv.Int64(v) != gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) != gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "不等于", v1Str)
				return
			}
		case GT:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "大于", v1Str)
				return
			}

			switch {
			case kind == reflect.String && gconv.String(v) > gconv.String(v1):
			case kind == reflect.Int64 && gconv.Int64(v) > gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) > gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "大于", v1Str)
				return
			}
		case GE:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "大于等于", v1Str)
				return
			}

			switch {
			case kind == reflect.String && gconv.String(v) >= gconv.String(v1):
			case kind == reflect.Int64 && gconv.Int64(v) >= gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) >= gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "大于等于", v1Str)
				return
			}
		case LT:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "小于", v1Str)
				return
			}
			switch {
			case kind == reflect.String && gconv.String(v) < gconv.String(v1):
			case kind == reflect.Int64 && gconv.Int64(v) < gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) < gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "小于", v1Str)
				return
			}
		case LE:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "小于等于", v1Str)
				return
			}
			switch {
			case kind == reflect.String && gconv.String(v) <= gconv.String(v1):
			case kind == reflect.Int64 && gconv.Int64(v) <= gconv.Int64(v1):
			case kind == reflect.Float64 && gconv.Float64(v) <= gconv.Float64(v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "小于等于", v1Str)
				return
			}
		case IN:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "被包含于", v1Str)
				return
			}
			switch {
			case kind == reflect.String && contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Int64 && contains(gconv.Int64s(v1), gconv.Int64(v)):
			case kind == reflect.Float64 && contains(gconv.Float64s(v1), gconv.Float64(v)):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "被包含于", v1Str)
				return
			}
		case NI:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Map:
				panic(fmt.Errorf("map类型不支持ExtraKey:%v", k1))
			case kind == reflect.Slice:
				panic(fmt.Errorf("slice类型不支持ExtraKey:%v", k1))
			}

			if v == nil || v1 == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "不被包含于", v1Str)
				return
			}
			switch {
			case kind == reflect.String && !contains(gconv.Strings(v1), gconv.String(v)):
			case kind == reflect.Int64 && !contains(gconv.Int64s(v1), gconv.Int64(v)):
			case kind == reflect.Float64 && !contains(gconv.Float64s(v1), gconv.Float64(v)):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "不被包含于", v1Str)
				return
			}
		case CO:
			switch {
			case kind == reflect.Bool:
				panic(fmt.Errorf("bool类型不支持ExtraKey:%v", k1))
			case kind == reflect.Int64:
				panic(fmt.Errorf("int64类型不支持ExtraKey:%v", k1))
			case kind == reflect.Float64:
				panic(fmt.Errorf("float64类型不支持ExtraKey:%v", k1))
			}

			if v == nil {
				ok = false
				expect = fmt.Sprintf("%s%s", "包含", v1Str)
				if reflect.Map == kind {
					expect = fmt.Sprintf("%s%s", "包含key:", gconv.String(v1))
				}
				return
			}
			switch {
			case kind == reflect.String && strings.Contains(gconv.String(v), gconv.String(v1)):
			case kind == reflect.Map && containsKey(v.(map[string]interface{}), gconv.String(v1)):
			case kind == reflect.Slice && contains(v.([]interface{}), v1):
			default:
				ok = false
				expect = fmt.Sprintf("%s%s", "包含", v1Str)
				if reflect.Map == kind {
					expect = fmt.Sprintf("%s%s", "包含key:", gconv.String(v1))
				}
				return
			}
		default:
			panic(fmt.Errorf("不支持的ExtraKey:%v", k1))
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
