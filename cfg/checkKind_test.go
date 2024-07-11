package cfg

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"reflect"
	"testing"
)

func Test_checkKind(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testWithExcel.Assert("checkKindBool1", fmt.Sprint(checkKind(true, reflect.Bool)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindBool2", fmt.Sprint(checkKind(false, reflect.Bool)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindBool3", fmt.Sprint(checkKind("", reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		testWithExcel.Assert("checkKindBool4", fmt.Sprint(checkKind(0, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		testWithExcel.Assert("checkKindBool5", fmt.Sprint(checkKind(1, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		testWithExcel.Assert("checkKindBool5", fmt.Sprint(checkKind(0.1, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		testWithExcel.Assert("checkKindBool6", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		testWithExcel.Assert("checkKindBool7", fmt.Sprint(checkKind([]string{}, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))

		testWithExcel.Assert("checkKindInt1", fmt.Sprint(checkKind(true, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt2", fmt.Sprint(checkKind(false, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt3", fmt.Sprint(checkKind("", reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt4", fmt.Sprint(checkKind(0, reflect.Int64)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindInt5", fmt.Sprint(checkKind(1, reflect.Int64)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindInt6", fmt.Sprint(checkKind(-1, reflect.Int64)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindInt7", fmt.Sprint(checkKind(0.1, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt8", fmt.Sprint(checkKind(-0.1, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))
		testWithExcel.Assert("checkKindInt10", fmt.Sprint(checkKind([]string{}, reflect.Int64)), fmt.Sprintf("%t%s", false, "Int"))

		testWithExcel.Assert("checkKindFloat1", fmt.Sprint(checkKind(true, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat2", fmt.Sprint(checkKind(false, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat3", fmt.Sprint(checkKind("", reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat4", fmt.Sprint(checkKind(0, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat5", fmt.Sprint(checkKind(1, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat6", fmt.Sprint(checkKind(-1, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat7", fmt.Sprint(checkKind(0.1, reflect.Float64)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindFloat8", fmt.Sprint(checkKind(-0.1, reflect.Float64)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindFloat9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))
		testWithExcel.Assert("checkKindFloat10", fmt.Sprint(checkKind([]string{}, reflect.Float64)), fmt.Sprintf("%t%s", false, "Float"))

		testWithExcel.Assert("checkKindMap1", fmt.Sprint(checkKind(true, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap2", fmt.Sprint(checkKind(false, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap3", fmt.Sprint(checkKind("", reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap4", fmt.Sprint(checkKind(0, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap5", fmt.Sprint(checkKind(1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap6", fmt.Sprint(checkKind(-1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap7", fmt.Sprint(checkKind(0.1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap8", fmt.Sprint(checkKind(-0.1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		testWithExcel.Assert("checkKindMap9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Map)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindMap10", fmt.Sprint(checkKind([]string{}, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))

		testWithExcel.Assert("checkKindSlice1", fmt.Sprint(checkKind(true, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice2", fmt.Sprint(checkKind(false, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice3", fmt.Sprint(checkKind("", reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice4", fmt.Sprint(checkKind(0, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice5", fmt.Sprint(checkKind(1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice6", fmt.Sprint(checkKind(-1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice7", fmt.Sprint(checkKind(0.1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice8", fmt.Sprint(checkKind(-0.1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		testWithExcel.Assert("checkKindSlice10", fmt.Sprint(checkKind([]string{}, reflect.Slice)), fmt.Sprintf("%t%s", true, ""))

		testWithExcel.Assert("checkKindString1", fmt.Sprint(checkKind(true, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString2", fmt.Sprint(checkKind(false, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString3", fmt.Sprint(checkKind("", reflect.String)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkKindString4", fmt.Sprint(checkKind(0, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString5", fmt.Sprint(checkKind(1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString6", fmt.Sprint(checkKind(-1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString7", fmt.Sprint(checkKind(0.1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString8", fmt.Sprint(checkKind(-0.1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		testWithExcel.Assert("checkKindString10", fmt.Sprint(checkKind([]string{}, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
	})
}
