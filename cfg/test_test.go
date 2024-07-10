package cfg

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"reflect"
	"testing"
)

func TestRule(t *testing.T) {
	extra := make(map[ExtraKey]interface{})

	gtest.C(t, func(t *gtest.T) {
		//testWithExcel.Assert("必填判断1", checkLevel(nil, MustInput), false)
		//testWithExcel.Assert("必填判断2", checkLevel("", MustInput), true)
		//testWithExcel.Assert("必填判断3", checkLevel("1", MustInput), true)
		//testWithExcel.Assert("必须为零值1", checkLevel(nil, MustInputZero), false)
		//testWithExcel.Assert("必须为零值2", checkLevel("", MustInputZero), true)
		//testWithExcel.Assert("必须为零值3", checkLevel(0, MustInputZero), true)
		//testWithExcel.Assert("必须为零值4", checkLevel(false, MustInputZero), true)
		//testWithExcel.Assert("必须为零值5", checkLevel(map[string]interface{}{}, MustInputZero), false)
		//testWithExcel.Assert("必须为零值6", checkLevel([]string{}, MustInputZero), false)
		//testWithExcel.Assert("必须为零值7", checkLevel(1, MustInputZero), false)
		//
		//testWithExcel.Assert("必须非零值1", checkLevel(nil, MustInputNotZero), false)
		//testWithExcel.Assert("必须非零值2", checkLevel("", MustInputNotZero), false)
		//testWithExcel.Assert("必须非零值3", checkLevel(0, MustInputNotZero), false)
		//testWithExcel.Assert("必须非零值4", checkLevel(false, MustInputNotZero), false)
		//testWithExcel.Assert("必须非零值5", checkLevel(map[string]interface{}{}, MustInputNotZero), true)
		//testWithExcel.Assert("必须非零值6", checkLevel([]string{}, MustInputNotZero), true)
		//testWithExcel.Assert("必须非零值7", checkLevel(1, MustInputNotZero), true)
		//testWithExcel.Assert("必须非零值8", checkLevel(true, MustInputNotZero), true)
		//testWithExcel.Assert("必须非零值9", checkLevel("1", MustInputNotZero), true)
		//
		//testWithExcel.Assert("建议非空1", checkLevel(nil, ProposalNotEmpty), false)
		//testWithExcel.Assert("建议非空2", checkLevel("", ProposalNotEmpty), false)
		//testWithExcel.Assert("建议非空3", checkLevel(0, ProposalNotEmpty), false)
		//testWithExcel.Assert("建议非空4", checkLevel(false, ProposalNotEmpty), false)
		//testWithExcel.Assert("建议非空5", checkLevel(map[string]interface{}{}, ProposalNotEmpty), true)
		//testWithExcel.Assert("建议非空6", checkLevel([]string{}, ProposalNotEmpty), true)
		//testWithExcel.Assert("建议非空7", checkLevel(1, ProposalNotEmpty), true)
		//testWithExcel.Assert("建议非空8", checkLevel(true, ProposalNotEmpty), true)
		//testWithExcel.Assert("建议非空9", checkLevel("1", ProposalNotEmpty), true)
		//
		//testWithExcel.Assert("选填1", checkLevel(nil, OptionalInput), true)
		//testWithExcel.Assert("选填2", checkLevel("", OptionalInput), true)
		//testWithExcel.Assert("选填3", checkLevel(0, OptionalInput), true)
		//testWithExcel.Assert("选填4", checkLevel(false, OptionalInput), true)
		//testWithExcel.Assert("选填5", checkLevel(map[string]interface{}{}, OptionalInput), true)
		//testWithExcel.Assert("选填6", checkLevel([]string{}, OptionalInput), true)
		//testWithExcel.Assert("选填7", checkLevel(1, OptionalInput), true)
		//testWithExcel.Assert("选填8", checkLevel(true, OptionalInput), true)
		//testWithExcel.Assert("选填9", checkLevel("1", OptionalInput), true)
		//
		//testWithExcel.Assert("建议为空1", checkLevel(nil, ProposalEmpty), true)
		//testWithExcel.Assert("建议为空2", checkLevel("", ProposalEmpty), true)
		//testWithExcel.Assert("建议为空3", checkLevel(0, ProposalEmpty), true)
		//testWithExcel.Assert("建议为空4", checkLevel(false, ProposalEmpty), true)
		//testWithExcel.Assert("建议为空5", checkLevel(map[string]interface{}{}, ProposalEmpty), false)
		//testWithExcel.Assert("建议为空6", checkLevel([]string{}, ProposalEmpty), false)
		//testWithExcel.Assert("建议为空7", checkLevel(1, ProposalEmpty), false)
		//testWithExcel.Assert("建议为空8", checkLevel(true, ProposalEmpty), false)
		//testWithExcel.Assert("建议为空9", checkLevel("1", ProposalEmpty), false)
		//
		//testWithExcel.Assert("checkKindBool1", fmt.Sprint(checkKind(true, reflect.Bool)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindBool2", fmt.Sprint(checkKind(false, reflect.Bool)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindBool3", fmt.Sprint(checkKind("", reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//testWithExcel.Assert("checkKindBool4", fmt.Sprint(checkKind(0, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//testWithExcel.Assert("checkKindBool5", fmt.Sprint(checkKind(1, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//testWithExcel.Assert("checkKindBool5", fmt.Sprint(checkKind(0.1, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//testWithExcel.Assert("checkKindBool6", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//testWithExcel.Assert("checkKindBool7", fmt.Sprint(checkKind([]string{}, reflect.Bool)), fmt.Sprintf("%t%s", false, "Bool"))
		//
		//testWithExcel.Assert("checkKindInt1", fmt.Sprint(checkKind(true, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt2", fmt.Sprint(checkKind(false, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt3", fmt.Sprint(checkKind("", reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt4", fmt.Sprint(checkKind(0, reflect.Int)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindInt5", fmt.Sprint(checkKind(1, reflect.Int)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindInt6", fmt.Sprint(checkKind(-1, reflect.Int)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindInt7", fmt.Sprint(checkKind(0.1, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt8", fmt.Sprint(checkKind(-0.1, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//testWithExcel.Assert("checkKindInt10", fmt.Sprint(checkKind([]string{}, reflect.Int)), fmt.Sprintf("%t%s", false, "Int"))
		//
		//testWithExcel.Assert("checkKindUint1", fmt.Sprint(checkKind(true, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint2", fmt.Sprint(checkKind(false, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint3", fmt.Sprint(checkKind("", reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint4", fmt.Sprint(checkKind(0, reflect.Uint)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindUint5", fmt.Sprint(checkKind(1, reflect.Uint)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindUint6", fmt.Sprint(checkKind(-1, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint7", fmt.Sprint(checkKind(0.1, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint8", fmt.Sprint(checkKind(-0.1, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//testWithExcel.Assert("checkKindUint10", fmt.Sprint(checkKind([]string{}, reflect.Uint)), fmt.Sprintf("%t%s", false, "Uint"))
		//
		//testWithExcel.Assert("checkKindFloat1", fmt.Sprint(checkKind(true, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat2", fmt.Sprint(checkKind(false, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat3", fmt.Sprint(checkKind("", reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat4", fmt.Sprint(checkKind(0, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat5", fmt.Sprint(checkKind(1, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat6", fmt.Sprint(checkKind(-1, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat7", fmt.Sprint(checkKind(0.1, reflect.Float32)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindFloat8", fmt.Sprint(checkKind(-0.1, reflect.Float32)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindFloat9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//testWithExcel.Assert("checkKindFloat10", fmt.Sprint(checkKind([]string{}, reflect.Float32)), fmt.Sprintf("%t%s", false, "Float"))
		//
		//testWithExcel.Assert("checkKindMap1", fmt.Sprint(checkKind(true, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap2", fmt.Sprint(checkKind(false, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap3", fmt.Sprint(checkKind("", reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap4", fmt.Sprint(checkKind(0, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap5", fmt.Sprint(checkKind(1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap6", fmt.Sprint(checkKind(-1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap7", fmt.Sprint(checkKind(0.1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap8", fmt.Sprint(checkKind(-0.1, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//testWithExcel.Assert("checkKindMap9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Map)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindMap10", fmt.Sprint(checkKind([]string{}, reflect.Map)), fmt.Sprintf("%t%s", false, "Map"))
		//
		//testWithExcel.Assert("checkKindSlice1", fmt.Sprint(checkKind(true, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice2", fmt.Sprint(checkKind(false, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice3", fmt.Sprint(checkKind("", reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice4", fmt.Sprint(checkKind(0, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice5", fmt.Sprint(checkKind(1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice6", fmt.Sprint(checkKind(-1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice7", fmt.Sprint(checkKind(0.1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice8", fmt.Sprint(checkKind(-0.1, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.Slice)), fmt.Sprintf("%t%s", false, "Slice"))
		//testWithExcel.Assert("checkKindSlice10", fmt.Sprint(checkKind([]string{}, reflect.Slice)), fmt.Sprintf("%t%s", true, ""))
		//
		//testWithExcel.Assert("checkKindString1", fmt.Sprint(checkKind(true, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString2", fmt.Sprint(checkKind(false, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString3", fmt.Sprint(checkKind("", reflect.String)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkKindString4", fmt.Sprint(checkKind(0, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString5", fmt.Sprint(checkKind(1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString6", fmt.Sprint(checkKind(-1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString7", fmt.Sprint(checkKind(0.1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString8", fmt.Sprint(checkKind(-0.1, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString9", fmt.Sprint(checkKind(map[string]interface{}{}, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//testWithExcel.Assert("checkKindString10", fmt.Sprint(checkKind([]string{}, reflect.String)), fmt.Sprintf("%t%s", false, "String"))
		//
		//extra = map[ExtraKey]interface{}{
		//	EQ: "1",
		//}
		//testWithExcel.Assert("checkExtraEQString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		//testWithExcel.Assert("checkExtraEQString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkExtraEQString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		//testWithExcel.Assert("checkExtraEQString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		//testWithExcel.Assert("checkExtraEQString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		//testWithExcel.Assert("checkExtraEQString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		//testWithExcel.Assert("checkExtraEQString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))
		//testWithExcel.Assert("checkExtraEQString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", false, "等于string1"))

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
		testWithExcel.Assert("checkExtraEQInt1", fmt.Sprint(checkExtra(true, reflect.Int, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt2", fmt.Sprint(checkExtra("1", reflect.Int, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt3", fmt.Sprint(checkExtra(0, reflect.Int, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))
		testWithExcel.Assert("checkExtraEQInt4", fmt.Sprint(checkExtra(1, reflect.Int, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt5", fmt.Sprint(checkExtra(-1, reflect.Int, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt6", fmt.Sprint(checkExtra(0.1, reflect.Int, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraEQInt7", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.Int, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))
		testWithExcel.Assert("checkExtraEQInt8", fmt.Sprint(checkExtra([]string{}, reflect.Int, extra)), fmt.Sprintf("%t%s", false, "等于booltrue"))

		extra = map[ExtraKey]interface{}{
			NE: "1",
		}
		testWithExcel.Assert("checkExtraNEString1", fmt.Sprint(checkExtra(true, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString2", fmt.Sprint(checkExtra(false, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString3", fmt.Sprint(checkExtra("1", reflect.String, extra)), fmt.Sprintf("%t%s", false, "不等于string1"))
		testWithExcel.Assert("checkExtraNEString4", fmt.Sprint(checkExtra(0, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString5", fmt.Sprint(checkExtra(1, reflect.String, extra)), fmt.Sprintf("%t%s", false, "不等于string1"))
		testWithExcel.Assert("checkExtraNEString6", fmt.Sprint(checkExtra(-1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString7", fmt.Sprint(checkExtra(0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString8", fmt.Sprint(checkExtra(-0.1, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString9", fmt.Sprint(checkExtra(map[string]interface{}{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
		testWithExcel.Assert("checkExtraNEString10", fmt.Sprint(checkExtra([]string{}, reflect.String, extra)), fmt.Sprintf("%t%s", true, ""))
	})
}
