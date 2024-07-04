package testWithExcel

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
	"strings"
)

const (
	AssertTypeNone           = "none"   //不使用通用的断言，处理方法中自行处理
	AssertTypeEQ             = "eq"     //值相等期望值
	AssertTypeEQWithValue    = "eq:"    //值为json,提取指定字段值判断是否等于期望值,eq:为前缀,eq:后面为json的pattern,如“a.b”,下面的同理
	AssertTypeNE             = "ne"     //值不相等期望值
	AssertTypeNEWithValue    = "ne:"    //值为json,提取指定字段值判断是否不等于期望值,ne:为前缀
	AssertTypeGT             = "gt"     //值大于期望值
	AssertTypeGTWithValue    = "gt:"    //值为json,提取指定字段值判断是否大于期望值,gt:为前缀
	AssertTypeGE             = "ge"     //值大于等于期望值
	AssertTypeGEWithValue    = "ge:"    //值为json,提取指定字段值判断是否大于等于期望值,ge:为前缀
	AssertTypeLT             = "lt"     //值小于期望值
	AssertTypeLTWithValue    = "lt:"    //值为json,提取指定字段值判断是否小于期望值,lt:为前缀
	AssertTypeLE             = "le"     //值小于等于期望值
	AssertTypeLEWithValue    = "le:"    //值为json,提取指定字段值判断是否小于等于期望值,le:为前缀
	AssertTypeIN             = "in"     //值属于期望值，期望值为json数组
	AssertTypeINWithValue    = "in:"    //值为json,提取指定字段值判断属于期望值,in:为前缀,期望值为json数组
	AssertTypeNI             = "ni"     //值不属于期望值，期望值为json数组
	AssertTypeNIWithValue    = "ni:"    //值为json,提取指定字段值判断不属于期望值,ni:为前缀,期望值为json数组
	AssertTypeCntEqWithValue = "cntEq:" //值为json,提取指定字段值判断长度等于期望值,cntEq:为前缀
	AssertTypeCntLtWithValue = "cntLt:" //值为json,提取指定字段值判断长度小于期望值,cntLt:为前缀
	AssertTypeCntLeWithValue = "cntLe:" //值为json,提取指定字段值判断长度等于等于期望值,cntLe:为前缀
	AssertTypeCntGtWithValue = "cntGt:" //值为json,提取指定字段值判断长度大于期望值,cntGt:为前缀
	AssertTypeCntGeWithValue = "cntGe:" //值为json,提取指定字段值判断长度大于等于期望值,cntGe:为前缀
)

// AssertByType checks `value` and `expect` by AssertType.
func AssertByType(type_ string, caseName string, value interface{}, expect string) {
	typeNew, pattern, c2 := strings.Cut(type_, ":")
	valueNew := value
	if c2 && gstr.InArray([]string{AssertTypeEQWithValue, AssertTypeNEWithValue, AssertTypeGTWithValue, AssertTypeGEWithValue, AssertTypeLTWithValue, AssertTypeLEWithValue, AssertTypeINWithValue, AssertTypeNIWithValue, AssertTypeCntEqWithValue, AssertTypeCntLtWithValue, AssertTypeCntLeWithValue, AssertTypeCntGtWithValue, AssertTypeCntGeWithValue}, typeNew+":") {
		if strings.Trim(pattern, "") == "" {
			panic(fmt.Sprintf("[ASSERT] %v AssertType %v json pattern为空", caseName, type_))
		}
		if j, err := gjson.DecodeToJson(value); err != nil {
			panic(fmt.Sprintf("[ASSERT] %v value %v json解析失败:%v", caseName, value, err))
		} else {
			valueNew = j.Get(pattern).Interface()
		}
		if gstr.InArray([]string{AssertTypeCntEqWithValue, AssertTypeCntLtWithValue, AssertTypeCntLeWithValue, AssertTypeCntGtWithValue, AssertTypeCntGeWithValue}, typeNew+":") {
			typeNew = typeNew + ":"
			l, err := lenForValue(valueNew)
			if err != nil {
				panic(fmt.Sprintf("[ASSERT] %v EXPECT %v 获取值长度失败:%v", caseName, valueNew, err))
			}
			valueNew = l
		}
	}
	switch typeNew {
	case AssertTypeNone:
		//忽略不做处理
		return
	case AssertTypeEQ:
		Assert(caseName, valueNew, expect)
	case AssertTypeNE:
		AssertNE(caseName, valueNew, expect)
	case AssertTypeGT:
		AssertGT(caseName, valueNew, expect)
	case AssertTypeGE:
		AssertGE(caseName, valueNew, expect)
	case AssertTypeLT:
		AssertLT(caseName, valueNew, expect)
	case AssertTypeLE:
		AssertLE(caseName, valueNew, expect)
	case AssertTypeIN:
		expectN, err := jsonToStrArr(expect)
		if err != nil {
			panic(fmt.Sprintf("[ASSERT] %v EXPECT %v json解析失败:%v", caseName, expect, err))
		}
		AssertIN(caseName, valueNew, expectN)
	case AssertTypeNI:
		expectN, err := jsonToStrArr(expect)
		if err != nil {
			panic(fmt.Sprintf("[ASSERT] %v EXPECT %v json解析失败:%v", caseName, expect, err))
		}
		AssertNI(caseName, valueNew, expectN)
	case AssertTypeCntEqWithValue:
		Assert(caseName, valueNew, expect)
	case AssertTypeCntLtWithValue:
		AssertLT(caseName, valueNew, expect)
	case AssertTypeCntLeWithValue:
		AssertLE(caseName, valueNew, expect)
	case AssertTypeCntGtWithValue:
		AssertGT(caseName, valueNew, expect)
	case AssertTypeCntGeWithValue:
		AssertGE(caseName, valueNew, expect)
	default:
		panic(fmt.Sprintf("不合法的断言类型:%s", typeNew))
	}
}

func jsonToStrArr(jsonStr string) (ret []any, err error) {
	err = gjson.Unmarshal([]byte(jsonStr), &ret)
	return
}

func lenForValue(value interface{}) (ret int, err error) {
	if value == nil {
		return 0, nil
	}
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		ret = reflect.ValueOf(value).Len()
	default:
		err = fmt.Errorf("类型不合法")
	}
	return
}

// Assert checks `value` and `expect` EQUAL.
func Assert(caseName string, value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	var (
		strValue  = gconv.String(value)
		strExpect = gconv.String(expect)
	)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v == %v`, caseName, strValue, strExpect))
	}
}

// AssertEQ checks `value` and `expect` EQUAL, including their TYPES.
func AssertEQ(caseName string, value, expect interface{}) {
	// Value assert.
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err != nil {
			panic(err)
		}
		return
	}
	strValue := gconv.String(value)
	strExpect := gconv.String(expect)
	if strValue != strExpect {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v == %v`, caseName, strValue, strExpect))
	}
	// Type assert.
	t1 := reflect.TypeOf(value)
	t2 := reflect.TypeOf(expect)
	if t1 != t2 {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT TYPE %v[%v] == %v[%v]`, caseName, strValue, t1, strExpect, t2))
	}
}

// AssertNE checks `value` and `expect` NOT EQUAL.
func AssertNE(caseName string, value, expect interface{}) {
	rvExpect := reflect.ValueOf(expect)
	if isNil(value) {
		value = nil
	}
	if rvExpect.Kind() == reflect.Map {
		if err := compareMap(value, expect); err == nil {
			panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v != %v`, caseName, value, expect))
		}
		return
	}
	var (
		strValue  = gconv.String(value)
		strExpect = gconv.String(expect)
	)
	if strValue == strExpect {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v != %v`, caseName, strValue, strExpect))
	}
}

// AssertNQ checks `value` and `expect` NOT EQUAL, including their TYPES.
func AssertNQ(caseName string, value, expect interface{}) {
	// Type assert.
	t1 := reflect.TypeOf(value)
	t2 := reflect.TypeOf(expect)
	if t1 == t2 {
		panic(
			fmt.Sprintf(
				`[ASSERT] %v EXPECT TYPE %v[%v] != %v[%v]`,
				caseName, gconv.String(value), t1, gconv.String(expect), t2,
			),
		)
	}
	// Value assert.
	AssertNE(caseName, value, expect)
}

// AssertGT checks `value` is GREATER THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGT,
// others are invalid.
func AssertGT(caseName string, value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) > gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) > gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) > gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) > gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v > %v`, caseName, value, expect))
	}
}

// AssertGE checks `value` is GREATER OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertGTE,
// others are invalid.
func AssertGE(caseName string, value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) >= gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int64(value) >= gconv.Int64(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint64(value) >= gconv.Uint64(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) >= gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(
			`[ASSERT] %v EXPECT %v(%v) >= %v(%v)`,
			caseName,
			value, reflect.ValueOf(value).Kind(),
			expect, reflect.ValueOf(expect).Kind(),
		))
	}
}

// AssertLT checks `value` is LESS EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLT,
// others are invalid.
func AssertLT(caseName string, value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) < gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) < gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) < gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) < gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v < %v`, caseName, value, expect))
	}
}

// AssertLE checks `value` is LESS OR EQUAL THAN `expect`.
// Notice that, only string, integer and float types can be compared by AssertLTE,
// others are invalid.
func AssertLE(caseName string, value, expect interface{}) {
	passed := false
	switch reflect.ValueOf(expect).Kind() {
	case reflect.String:
		passed = gconv.String(value) <= gconv.String(expect)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		passed = gconv.Int(value) <= gconv.Int(expect)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		passed = gconv.Uint(value) <= gconv.Uint(expect)

	case reflect.Float32, reflect.Float64:
		passed = gconv.Float64(value) <= gconv.Float64(expect)
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v <= %v`, caseName, value, expect))
	}
}

// AssertIN checks `value` is IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.
func AssertIN(caseName string, value, expect interface{}) {
	var (
		passed     = true
		expectKind = reflect.ValueOf(expect).Kind()
	)
	switch expectKind {
	case reflect.Slice, reflect.Array:
		expectSlice := gconv.Strings(expect)
		if len(gconv.Strings(value)) == 0 {
			passed = false
		}
		for _, v1 := range gconv.Strings(value) {
			result := false
			for _, v2 := range expectSlice {
				if v1 == v2 {
					result = true
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	default:
		panic(fmt.Sprintf(`[ASSERT] %v INVALID EXPECT VALUE TYPE: %v`, caseName, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v IN %v`, caseName, value, expect))
	}
}

// AssertNI checks `value` is NOT IN `expect`.
// The `expect` should be a slice,
// but the `value` can be a slice or a basic type variable.
func AssertNI(caseName string, value, expect interface{}) {
	var (
		passed     = true
		expectKind = reflect.ValueOf(expect).Kind()
	)
	switch expectKind {
	case reflect.Slice, reflect.Array:
		if len(gconv.Strings(value)) == 0 {
			passed = false
		}
		for _, v1 := range gconv.Strings(value) {
			result := true
			for _, v2 := range gconv.Strings(expect) {
				if v1 == v2 {
					result = false
					break
				}
			}
			if !result {
				passed = false
				break
			}
		}
	default:
		panic(fmt.Sprintf(`[ASSERT] %v INVALID EXPECT VALUE TYPE: %v`, caseName, expectKind))
	}
	if !passed {
		panic(fmt.Sprintf(`[ASSERT] %v EXPECT %v NOT IN %v`, caseName, value, expect))
	}
}

func isNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Interface,
		reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()

	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}
	}
	return false
}

func compareMap(value, expect interface{}) error {
	var (
		rvValue  = reflect.ValueOf(value)
		rvExpect = reflect.ValueOf(expect)
	)
	if rvExpect.Kind() == reflect.Map {
		if rvValue.Kind() == reflect.Map {
			if rvExpect.Len() == rvValue.Len() {
				// Turn two interface maps to the same type for comparison.
				// Direct use of rvValue.MapIndex(key).Interface() will panic
				// when the key types are inconsistent.
				mValue := make(map[string]string)
				mExpect := make(map[string]string)
				ksValue := rvValue.MapKeys()
				ksExpect := rvExpect.MapKeys()
				for _, key := range ksValue {
					mValue[gconv.String(key.Interface())] = gconv.String(rvValue.MapIndex(key).Interface())
				}
				for _, key := range ksExpect {
					mExpect[gconv.String(key.Interface())] = gconv.String(rvExpect.MapIndex(key).Interface())
				}
				for k, v := range mExpect {
					if v != mValue[k] {
						return fmt.Errorf(`[ASSERT] EXPECT VALUE map["%v"]:%v == map["%v"]:%v`+
							"\nGIVEN : %v\nEXPECT: %v", k, mValue[k], k, v, mValue, mExpect)
					}
				}
			} else {
				return fmt.Errorf(`[ASSERT] EXPECT MAP LENGTH %d == %d`, rvValue.Len(), rvExpect.Len())
			}
		} else {
			return fmt.Errorf(`[ASSERT] EXPECT VALUE TO BE A MAP, BUT GIVEN "%s"`, rvValue.Kind())
		}
	}
	return nil
}
