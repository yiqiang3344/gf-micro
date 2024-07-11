package cfg

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"reflect"
	"testing"
)

func Test_checkExtra(t *testing.T) {
	extra := make(map[ExtraKey]interface{})

	gtest.C(t, func(t *gtest.T) {
		extra = map[ExtraKey]interface{}{
			EQ: "1",
		}
		testWithExcel.Assert("checkExtraEQString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		testWithExcel.Assert("checkExtraEQString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		testWithExcel.Assert("checkExtraEQString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		testWithExcel.Assert("checkExtraEQString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		testWithExcel.Assert("checkExtraEQString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		testWithExcel.Assert("checkExtraEQString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))

		extra = map[ExtraKey]interface{}{
			EQ: true,
		}
		testWithExcel.Assert("checkExtraEQBool1", fmt.Sprint(checkExtra(true, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQBool2", fmt.Sprint(checkExtra("1", reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQBool3", fmt.Sprint(checkExtra(0, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))
		testWithExcel.Assert("checkExtraEQBool4", fmt.Sprint(checkExtra(1, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQBool5", fmt.Sprint(checkExtra(-1, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQBool6", fmt.Sprint(checkExtra(0.1, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQBool7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))
		testWithExcel.Assert("checkExtraEQBool8", fmt.Sprint(checkExtra([]string{}, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))

		extra = map[ExtraKey]interface{}{
			EQ: 1,
		}
		testWithExcel.Assert("checkExtraEQInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "等于int1"))
		testWithExcel.Assert("checkExtraEQInt4", fmt.Sprint(checkExtra(1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "等于int1"))
		testWithExcel.Assert("checkExtraEQInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "等于int1"))
		testWithExcel.Assert("checkExtraEQInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "等于int1"))
		testWithExcel.Assert("checkExtraEQInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "等于int1"))

		extra = map[ExtraKey]interface{}{
			EQ: 1.0,
		}
		testWithExcel.Assert("checkExtraEQFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "等于float641"))
		testWithExcel.Assert("checkExtraEQFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "等于float641"))
		testWithExcel.Assert("checkExtraEQFloat4", fmt.Sprint(checkExtra(1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "等于float641"))
		testWithExcel.Assert("checkExtraEQFloat6", fmt.Sprint(checkExtra(1.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "等于float641"))
		testWithExcel.Assert("checkExtraEQFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "等于float641"))

		extra = map[ExtraKey]interface{}{
			EQ: map[string]interface{}{"a": 1},
		}
		testWithExcel.Assert("checkExtraEQMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraEQMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQMap8", fmt.Sprint(checkExtra([]string{}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "等于map[string]interface {}map[a:1]"))

		extra = map[ExtraKey]interface{}{
			EQ: []int{1, 2},
		}
		testWithExcel.Assert("checkExtraEQSlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "等于[]int[1 2]"))
		testWithExcel.Assert("checkExtraEQSlice8", fmt.Sprint(checkExtra([]int{1, 2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: "1",
		}
		testWithExcel.Assert("checkExtraNEString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", false, "不等于string1"))
		testWithExcel.Assert("checkExtraNEString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "不等于string1"))
		testWithExcel.Assert("checkExtraNEString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: true,
		}
		testWithExcel.Assert("checkExtraNEBool1", fmt.Sprint(checkExtra(true, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "不等于booltrue"))
		testWithExcel.Assert("checkExtraNEBool2", fmt.Sprint(checkExtra("1", reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "不等于booltrue"))
		testWithExcel.Assert("checkExtraNEBool3", fmt.Sprint(checkExtra(0, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEBool4", fmt.Sprint(checkExtra(1, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "不等于booltrue"))
		testWithExcel.Assert("checkExtraNEBool5", fmt.Sprint(checkExtra(-1, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "不等于booltrue"))
		testWithExcel.Assert("checkExtraNEBool6", fmt.Sprint(checkExtra(0.1, reflect.Bool, extra)), fmt.Sprintf("%t%s", false, "不等于booltrue"))
		testWithExcel.Assert("checkExtraNEBool7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEBool8", fmt.Sprint(checkExtra([]string{}, reflect.Bool, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: 1,
		}
		testWithExcel.Assert("checkExtraNEInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "不等于int1"))
		testWithExcel.Assert("checkExtraNEInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "不等于int1"))
		testWithExcel.Assert("checkExtraNEInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEInt4", fmt.Sprint(checkExtra(1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "不等于int1"))
		testWithExcel.Assert("checkExtraNEInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: 1.0,
		}
		testWithExcel.Assert("checkExtraNEFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不等于float641"))
		testWithExcel.Assert("checkExtraNEFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEFloat4", fmt.Sprint(checkExtra(1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不等于float641"))
		testWithExcel.Assert("checkExtraNEFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEFloat6", fmt.Sprint(checkExtra(1.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不等于float641"))
		testWithExcel.Assert("checkExtraNEFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: map[string]interface{}{"a": 1},
		}
		testWithExcel.Assert("checkExtraNEMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "不等于map[string]interface {}map[a:1]"))
		testWithExcel.Assert("checkExtraNEMap8", fmt.Sprint(checkExtra([]string{}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NE: []int{1, 2},
		}
		testWithExcel.Assert("checkExtraNESlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNESlice8", fmt.Sprint(checkExtra([]int{1, 2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "不等于[]int[1 2]"))

		extra = map[ExtraKey]interface{}{
			GT: "1",
		}
		testWithExcel.Assert("checkExtraGTString1", fmt.Sprint(checkExtra(false, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTString3", fmt.Sprint(checkExtra("2", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTString4", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于string1"))
		testWithExcel.Assert("checkExtraGTString5", fmt.Sprint(checkExtra(2, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于string1"))
		testWithExcel.Assert("checkExtraGTString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于string1"))
		testWithExcel.Assert("checkExtraGTString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			GT: 1,
		}
		testWithExcel.Assert("checkExtraGTInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))
		testWithExcel.Assert("checkExtraGTInt2", fmt.Sprint(checkExtra("2", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))
		testWithExcel.Assert("checkExtraGTInt4", fmt.Sprint(checkExtra(2, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))
		testWithExcel.Assert("checkExtraGTInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))
		testWithExcel.Assert("checkExtraGTInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))
		testWithExcel.Assert("checkExtraGTInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于int1"))

		extra = map[ExtraKey]interface{}{
			GT: 1.0,
		}
		testWithExcel.Assert("checkExtraGTFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于float641"))
		testWithExcel.Assert("checkExtraGTFloat2", fmt.Sprint(checkExtra("2", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于float641"))
		testWithExcel.Assert("checkExtraGTFloat4", fmt.Sprint(checkExtra(2, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于float641"))
		testWithExcel.Assert("checkExtraGTFloat6", fmt.Sprint(checkExtra(2.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于float641"))
		testWithExcel.Assert("checkExtraGTFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于float641"))

		extra = map[ExtraKey]interface{}{
			GT: 1,
		}
		testWithExcel.Assert("checkExtraGTMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGTMap8", fmt.Sprint(checkExtra([]string{}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))

		extra = map[ExtraKey]interface{}{
			GT: 1,
		}
		testWithExcel.Assert("checkExtraGTSlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于1"))
		testWithExcel.Assert("checkExtraGTSlice8", fmt.Sprint(checkExtra([]int{1, 2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			GE: "1",
		}
		testWithExcel.Assert("checkExtraGEString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于等于string1"))
		testWithExcel.Assert("checkExtraGEString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于等于string1"))
		testWithExcel.Assert("checkExtraGEString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "大于等于string1"))
		testWithExcel.Assert("checkExtraGEString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			GE: 1,
		}
		testWithExcel.Assert("checkExtraGEInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于等于int1"))
		testWithExcel.Assert("checkExtraGEInt4", fmt.Sprint(checkExtra(1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于等于int1"))
		testWithExcel.Assert("checkExtraGEInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于等于int1"))
		testWithExcel.Assert("checkExtraGEInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于等于int1"))
		testWithExcel.Assert("checkExtraGEInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "大于等于int1"))

		extra = map[ExtraKey]interface{}{
			GE: 1.0,
		}
		testWithExcel.Assert("checkExtraGEFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于等于float641"))
		testWithExcel.Assert("checkExtraGEFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于等于float641"))
		testWithExcel.Assert("checkExtraGEFloat4", fmt.Sprint(checkExtra(1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于等于float641"))
		testWithExcel.Assert("checkExtraGEFloat6", fmt.Sprint(checkExtra(1.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于等于float641"))
		testWithExcel.Assert("checkExtraGEFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "大于等于float641"))

		extra = map[ExtraKey]interface{}{
			GE: 1,
		}
		testWithExcel.Assert("checkExtraGEMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))
		testWithExcel.Assert("checkExtraGEMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGEMap8", fmt.Sprint(checkExtra([]string{}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))

		extra = map[ExtraKey]interface{}{
			GE: 1,
		}
		testWithExcel.Assert("checkExtraGESlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice8", fmt.Sprint(checkExtra([]int{1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraGESlice9", fmt.Sprint(checkExtra([]int{}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度大于等于1"))

		extra = map[ExtraKey]interface{}{
			LT: "1",
		}
		testWithExcel.Assert("checkExtraLTString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于string1"))
		testWithExcel.Assert("checkExtraLTString3", fmt.Sprint(checkExtra("0", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTString4", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于string1"))
		testWithExcel.Assert("checkExtraLTString5", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于string1"))
		testWithExcel.Assert("checkExtraLTString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于string1"))

		extra = map[ExtraKey]interface{}{
			LT: 1,
		}
		testWithExcel.Assert("checkExtraLTInt1", fmt.Sprint(checkExtra(false, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt2", fmt.Sprint(checkExtra("0", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt4", fmt.Sprint(checkExtra(2, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "小于int1"))
		testWithExcel.Assert("checkExtraLTInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LT: 1.0,
		}
		testWithExcel.Assert("checkExtraLTFloat1", fmt.Sprint(checkExtra(false, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat2", fmt.Sprint(checkExtra("0", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat4", fmt.Sprint(checkExtra(2, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "小于float641"))
		testWithExcel.Assert("checkExtraLTFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat6", fmt.Sprint(checkExtra(0.1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LT: 1,
		}
		testWithExcel.Assert("checkExtraLTMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTMap8", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLTMap9", fmt.Sprint(checkExtra([]string{}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LT: 1,
		}
		testWithExcel.Assert("checkExtraLTSlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于1"))
		testWithExcel.Assert("checkExtraLTSlice8", fmt.Sprint(checkExtra([]int{}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LE: "1",
		}
		testWithExcel.Assert("checkExtraLEString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于等于string1"))
		testWithExcel.Assert("checkExtraLEString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于等于string1"))
		testWithExcel.Assert("checkExtraLEString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "小于等于string1"))

		extra = map[ExtraKey]interface{}{
			LE: 1,
		}
		testWithExcel.Assert("checkExtraLEInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt4", fmt.Sprint(checkExtra(2, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "小于等于int1"))
		testWithExcel.Assert("checkExtraLEInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LE: 1.0,
		}
		testWithExcel.Assert("checkExtraLEFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEFloat4", fmt.Sprint(checkExtra(2, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "小于等于float641"))
		testWithExcel.Assert("checkExtraLEFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEFloat6", fmt.Sprint(checkExtra(2.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "小于等于float641"))
		testWithExcel.Assert("checkExtraLEFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			LE: 1,
		}
		testWithExcel.Assert("checkExtraLEMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度小于等于1"))
		testWithExcel.Assert("checkExtraLEMap8", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLEMap9", fmt.Sprint(checkExtra([]int{1, 2, 3}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "长度小于等于1"))

		extra = map[ExtraKey]interface{}{
			LE: 1,
		}
		testWithExcel.Assert("checkExtraLESlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1, "b": 1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice8", fmt.Sprint(checkExtra([]int{1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraLESlice9", fmt.Sprint(checkExtra([]int{1, 2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "长度小于等于1"))

		extra = map[ExtraKey]interface{}{
			IN: []string{"1"},
		}
		testWithExcel.Assert("checkExtraINString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraINString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINString4", fmt.Sprint(checkExtra(2, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraINString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraINString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraINString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraINString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "被包含于[]string[1]"))

		extra = map[ExtraKey]interface{}{
			IN: []int{1},
		}
		testWithExcel.Assert("checkExtraINInt1", fmt.Sprint(checkExtra(false, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraINInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINInt3", fmt.Sprint(checkExtra(1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINInt4", fmt.Sprint(checkExtra(2, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraINInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraINInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraINInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraINInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]int[1]"))

		extra = map[ExtraKey]interface{}{
			IN: []float64{1.0},
		}
		testWithExcel.Assert("checkExtraINFloat1", fmt.Sprint(checkExtra(false, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraINFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINFloat3", fmt.Sprint(checkExtra(1.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraINFloat4", fmt.Sprint(checkExtra(2, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraINFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraINFloat6", fmt.Sprint(checkExtra(0.1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraINFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraINFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "被包含于[]float64[1]"))

		extra = map[ExtraKey]interface{}{
			NI: []string{"1"},
		}
		testWithExcel.Assert("checkExtraNIString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString2", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraNIString3", fmt.Sprint(checkExtra("2", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]string[1]"))
		testWithExcel.Assert("checkExtraNIString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NI: []int{1},
		}
		testWithExcel.Assert("checkExtraNIInt1", fmt.Sprint(checkExtra(true, reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraNIInt2", fmt.Sprint(checkExtra("1", reflect.Int64, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]int[1]"))
		testWithExcel.Assert("checkExtraNIInt3", fmt.Sprint(checkExtra(0, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIInt4", fmt.Sprint(checkExtra(2, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIInt5", fmt.Sprint(checkExtra(-1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIInt6", fmt.Sprint(checkExtra(0.1, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			NI: []float64{1.0},
		}
		testWithExcel.Assert("checkExtraNIFloat1", fmt.Sprint(checkExtra(true, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIFloat2", fmt.Sprint(checkExtra("1", reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraNIFloat3", fmt.Sprint(checkExtra(0, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIFloat4", fmt.Sprint(checkExtra(1, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraNIFloat5", fmt.Sprint(checkExtra(-1, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIFloat6", fmt.Sprint(checkExtra(1.0, reflect.Float64, extra)), fmt.Sprintf("%t%s", false, "不被包含于[]float64[1]"))
		testWithExcel.Assert("checkExtraNIFloat7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNIFloat8", fmt.Sprint(checkExtra([]string{}, reflect.Float64, extra)), fmt.Sprintf("%t%s", true, ""))

		extra = map[ExtraKey]interface{}{
			CO: "a",
		}
		testWithExcel.Assert("checkExtraCOMap1", fmt.Sprint(checkExtra(true, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap2", fmt.Sprint(checkExtra("1", reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap3", fmt.Sprint(checkExtra(0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap4", fmt.Sprint(checkExtra(1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap5", fmt.Sprint(checkExtra(-1, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap6", fmt.Sprint(checkExtra(1.0, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap7", fmt.Sprint(checkExtra(map[string]interface{}{"b": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))
		testWithExcel.Assert("checkExtraCOMap8", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Map, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraCOMap9", fmt.Sprint(checkExtra([]interface{}{"a", "b"}, reflect.Map, extra)), fmt.Sprintf("%t%s", false, "包含key:a"))

		extra = map[ExtraKey]interface{}{
			CO: 1,
		}
		testWithExcel.Assert("checkExtraCOSlice1", fmt.Sprint(checkExtra(true, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice2", fmt.Sprint(checkExtra("1", reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice3", fmt.Sprint(checkExtra(0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice4", fmt.Sprint(checkExtra(1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice5", fmt.Sprint(checkExtra(-1, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice6", fmt.Sprint(checkExtra(1.0, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice7", fmt.Sprint(checkExtra(map[string]interface{}{"a": 1}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
		testWithExcel.Assert("checkExtraCOSlice8", fmt.Sprint(checkExtra([]interface{}{1, 2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraCOSlice9", fmt.Sprint(checkExtra([]interface{}{2}, reflect.Slice, extra)), fmt.Sprintf("%t%s", false, "包含int1"))
	})
}
