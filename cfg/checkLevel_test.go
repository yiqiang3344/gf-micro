package cfg

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"reflect"
	"testing"
)

func Test_checkLevel(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testWithExcel.Assert("必填判断1", checkLevel(nil, MustInput, reflect.String), false)
		testWithExcel.Assert("必填判断2", checkLevel("", MustInput, reflect.String), true)
		testWithExcel.Assert("必填判断3", checkLevel("1", MustInput, reflect.String), true)

		testWithExcel.Assert("必须为零值1", checkLevel(nil, MustInputZero, reflect.String), false)
		testWithExcel.Assert("必须为零值2", checkLevel("", MustInputZero, reflect.String), true)
		testWithExcel.Assert("必须为零值3", checkLevel(0, MustInputZero, reflect.Int64), true)
		testWithExcel.Assert("必须为零值4", checkLevel(false, MustInputZero, reflect.Bool), true)
		testWithExcel.Assert("必须为零值5", checkLevel(map[string]interface{}{}, MustInputZero, reflect.Map), true)
		testWithExcel.Assert("必须为零值6", checkLevel([]string{}, MustInputZero, reflect.Slice), true)
		testWithExcel.Assert("必须为零值7", checkLevel(1, MustInputZero, reflect.Int64), false)

		testWithExcel.Assert("必须非零值1", checkLevel(nil, MustInputNotZero, reflect.String), false)
		testWithExcel.Assert("必须非零值2", checkLevel("", MustInputNotZero, reflect.String), false)
		testWithExcel.Assert("必须非零值3", checkLevel(0, MustInputNotZero, reflect.Int64), false)
		testWithExcel.Assert("必须非零值4", checkLevel(false, MustInputNotZero, reflect.Bool), false)
		testWithExcel.Assert("必须非零值5", checkLevel(map[string]interface{}{}, MustInputNotZero, reflect.Map), false)
		testWithExcel.Assert("必须非零值6", checkLevel([]string{}, MustInputNotZero, reflect.Slice), false)
		testWithExcel.Assert("必须非零值7", checkLevel(1, MustInputNotZero, reflect.Int64), true)
		testWithExcel.Assert("必须非零值8", checkLevel(true, MustInputNotZero, reflect.Bool), true)
		testWithExcel.Assert("必须非零值9", checkLevel("1", MustInputNotZero, reflect.String), true)

		testWithExcel.Assert("建议非空1", checkLevel(nil, ProposalNotZero, reflect.String), false)
		testWithExcel.Assert("建议非空2", checkLevel("", ProposalNotZero, reflect.String), false)
		testWithExcel.Assert("建议非空3", checkLevel(0, ProposalNotZero, reflect.Int64), false)
		testWithExcel.Assert("建议非空4", checkLevel(false, ProposalNotZero, reflect.Bool), false)
		testWithExcel.Assert("建议非空5", checkLevel(map[string]interface{}{}, ProposalNotZero, reflect.Map), false)
		testWithExcel.Assert("建议非空6", checkLevel([]string{}, ProposalNotZero, reflect.Slice), false)
		testWithExcel.Assert("建议非空7", checkLevel(1, ProposalNotZero, reflect.Int64), true)
		testWithExcel.Assert("建议非空8", checkLevel(true, ProposalNotZero, reflect.Bool), true)
		testWithExcel.Assert("建议非空9", checkLevel("1", ProposalNotZero, reflect.String), true)

		testWithExcel.Assert("选填1", checkLevel(nil, OptionalInput, reflect.String), true)
		testWithExcel.Assert("选填2", checkLevel("", OptionalInput, reflect.String), true)
		testWithExcel.Assert("选填3", checkLevel(0, OptionalInput, reflect.Int64), true)
		testWithExcel.Assert("选填4", checkLevel(false, OptionalInput, reflect.Bool), true)
		testWithExcel.Assert("选填5", checkLevel(map[string]interface{}{}, OptionalInput, reflect.Map), true)
		testWithExcel.Assert("选填6", checkLevel([]string{}, OptionalInput, reflect.Slice), true)
		testWithExcel.Assert("选填7", checkLevel(1, OptionalInput, reflect.Int64), true)
		testWithExcel.Assert("选填8", checkLevel(true, OptionalInput, reflect.Bool), true)
		testWithExcel.Assert("选填9", checkLevel("1", OptionalInput, reflect.String), true)

		testWithExcel.Assert("建议为空1", checkLevel(nil, ProposalZero, reflect.String), false)
		testWithExcel.Assert("建议为空2", checkLevel("", ProposalZero, reflect.String), true)
		testWithExcel.Assert("建议为空3", checkLevel(0, ProposalZero, reflect.Int64), true)
		testWithExcel.Assert("建议为空4", checkLevel(false, ProposalZero, reflect.Bool), true)
		testWithExcel.Assert("建议为空5", checkLevel(map[string]interface{}{}, ProposalZero, reflect.Map), true)
		testWithExcel.Assert("建议为空6", checkLevel([]string{}, ProposalZero, reflect.Slice), true)
		testWithExcel.Assert("建议为空7", checkLevel(1, ProposalZero, reflect.String), false)
		testWithExcel.Assert("建议为空8", checkLevel(true, ProposalZero, reflect.Bool), false)
		testWithExcel.Assert("建议为空9", checkLevel("1", ProposalZero, reflect.String), false)
	})
}
