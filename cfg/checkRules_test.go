package cfg

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"reflect"
	"testing"
)

func Test_checkRules(t *testing.T) {
	ctx := context.Background()
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("./")
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("testConfig.yaml")

	type Case struct {
		Name   string
		Env    ENV
		Expect string
		opt    checkOpt
	}

	rules := []Case{
		//MustInput
		{Name: "MustInput1", Expect: "false,[错误]none值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput2", Expect: "false,[错误]testStrNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput3", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput4", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput5", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput6", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput7", Expect: "false,[错误]testBoolNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput8", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput9", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput10", Expect: "false,[错误]testIntNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput11", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput12", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput13", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput14", Expect: "false,[错误]testFloatNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput15", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput16", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput17", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput18", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput19", Expect: "false,[错误]testMapNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput20", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput21", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput22", Expect: "false,[错误]testSliceNil值为nil,期望必填", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Map, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput23", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput24", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput25", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInput26", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: MustInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		//MustInputZero
		{Name: "MustInputZero1", Expect: "false,[错误]none值为nil,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero2", Expect: "false,[错误]testStrNil值为nil,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero3", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero4", Expect: "false,[错误]testStr0值为string0,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero5", Expect: "false,[错误]testStr1值为string1,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero6", Expect: "false,[错误]testStr2值为string2,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero7", Expect: "false,[错误]testBoolNil值为nil,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero8", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero9", Expect: "false,[错误]testBooltrue值为booltrue,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero10", Expect: "false,[错误]testIntNil值为nil,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero11", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero12", Expect: "false,[错误]testInt1值为int641,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero13", Expect: "false,[错误]testInt2值为int642,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero14", Expect: "false,[错误]testFloatNil值为nil,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero15", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero16", Expect: "false,[错误]testFloat0p1值为float640.1,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero17", Expect: "false,[错误]testFloat1p1值为float641.1,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero18", Expect: "false,[错误]testFloat2p1值为float642.1,期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero19", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero20", Expect: "false,[错误]testMap1值为map[string]interface {}map[1:1],期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero21", Expect: "false,[错误]testMap2值为map[string]interface {}map[1:1 2:1],期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero22", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Map, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero23", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero24", Expect: "false,[错误]testSlice1值为[]interface {}[1],期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero25", Expect: "false,[错误]testSlice2值为[]interface {}[1 2],期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputZero26", Expect: "false,[错误]testSlice2Str值为[]interface {}[1 2],期望必须为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: MustInputZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		//MustInputNotZero
		{Name: "MustInputNotZero1", Expect: "false,[错误]none值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero2", Expect: "false,[错误]testStrNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero3", Expect: "false,[错误]testStrEmpty值为string,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero4", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero5", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero6", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero7", Expect: "false,[错误]testBoolNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero8", Expect: "false,[错误]testBoolfalse值为boolfalse,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero9", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero10", Expect: "false,[错误]testIntNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero11", Expect: "false,[错误]testInt0值为int640,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero12", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero13", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero14", Expect: "false,[错误]testFloatNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero15", Expect: "false,[错误]testFloat0p0值为int640,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero16", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero17", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero18", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero19", Expect: "false,[错误]testMapNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero20", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero21", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero22", Expect: "false,[错误]testSliceNil值为nil,期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Map, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero23", Expect: "false,[错误]testSliceEmpty值为[]interface {}[],期望必须非零值", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero24", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero25", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "MustInputNotZero26", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: MustInputNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		//ProposalNotZero
		{Name: "ProposalNotZero1", Expect: "false,[建议]none值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero2", Expect: "false,[建议]testStrNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero3", Expect: "false,[建议]testStrEmpty值为string,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero4", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero5", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero6", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero7", Expect: "false,[建议]testBoolNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero8", Expect: "false,[建议]testBoolfalse值为boolfalse,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero9", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero10", Expect: "false,[建议]testIntNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero11", Expect: "false,[建议]testInt0值为int640,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero12", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero13", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero14", Expect: "false,[建议]testFloatNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero15", Expect: "false,[建议]testFloat0p0值为int640,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero16", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero17", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero18", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero19", Expect: "false,[建议]testMapNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero20", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero21", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero22", Expect: "false,[建议]testSliceNil值为nil,期望非零值", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Slice, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero23", Expect: "false,[建议]testSliceEmpty值为[]interface {}[],期望非零值", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero24", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero25", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalNotZero26", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: ProposalNotZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		//OptionalInput
		{Name: "OptionalInput1", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput2", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput3", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput4", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput5", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput6", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput7", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput8", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput9", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput10", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput11", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput12", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput13", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput14", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput15", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput16", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput17", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput18", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput19", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput20", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput21", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput22", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Slice, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput23", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput24", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput25", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "OptionalInput26", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: OptionalInput, Env: []ENV{DEV, PROD}, Extra: nil}},
		//ProposalZero
		{Name: "ProposalZero1", Expect: "false,[提醒]none值为nil,期望为零值", Env: DEV, opt: checkOpt{Pattern: "none", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero2", Expect: "false,[提醒]testStrNil值为nil,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testStrNil", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero3", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testStrEmpty", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero4", Expect: "false,[提醒]testStr0值为string0,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testStr0", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero5", Expect: "false,[提醒]testStr1值为string1,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testStr1", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero6", Expect: "false,[提醒]testStr2值为string2,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testStr2", Kind: reflect.String, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero7", Expect: "false,[提醒]testBoolNil值为nil,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testBoolNil", Kind: reflect.Bool, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero8", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testBoolfalse", Kind: reflect.Bool, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero9", Expect: "false,[提醒]testBooltrue值为booltrue,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testBooltrue", Kind: reflect.Bool, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero10", Expect: "false,[提醒]testIntNil值为nil,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testIntNil", Kind: reflect.Int64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero11", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testInt0", Kind: reflect.Int64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero12", Expect: "false,[提醒]testInt1值为int641,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testInt1", Kind: reflect.Int64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero13", Expect: "false,[提醒]testInt2值为int642,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testInt2", Kind: reflect.Int64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero14", Expect: "false,[提醒]testFloatNil值为nil,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testFloatNil", Kind: reflect.Float64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero15", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testFloat0p0", Kind: reflect.Float64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero16", Expect: "false,[提醒]testFloat0p1值为float640.1,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat0p1", Kind: reflect.Float64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero17", Expect: "false,[提醒]testFloat1p1值为float641.1,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat1p1", Kind: reflect.Float64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero18", Expect: "false,[提醒]testFloat2p1值为float642.1,期望为零值", Env: DEV, opt: checkOpt{Pattern: "testFloat2p1", Kind: reflect.Float64, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero19", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testMapNil", Kind: reflect.Map, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero20", Expect: "false,[提醒]testMap1值为map[string]interface {}map[1:1],期望为零值", Env: DEV, opt: checkOpt{Pattern: "testMap1", Kind: reflect.Map, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero21", Expect: "false,[提醒]testMap2值为map[string]interface {}map[1:1 2:1],期望为零值", Env: DEV, opt: checkOpt{Pattern: "testMap2", Kind: reflect.Map, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero22", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceNil", Kind: reflect.Slice, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero23", Expect: "true,", Env: DEV, opt: checkOpt{Pattern: "testSliceEmpty", Kind: reflect.Slice, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero24", Expect: "false,[提醒]testSlice1值为[]interface {}[1],期望为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice1", Kind: reflect.Slice, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero25", Expect: "false,[提醒]testSlice2值为[]interface {}[1 2],期望为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice2", Kind: reflect.Slice, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
		{Name: "ProposalZero26", Expect: "false,[提醒]testSlice2Str值为[]interface {}[1 2],期望为零值", Env: DEV, opt: checkOpt{Pattern: "testSlice2Str", Kind: reflect.Slice, Level: ProposalZero, Env: []ENV{DEV, PROD}, Extra: nil}},
	}
	for _, v := range rules {
		gtest.C(t, func(t *gtest.T) {
			defer func(rule Case) {
				if e := recover(); e != nil {
					t.Errorf("ERROR %s:%v", v.Name, e)
				}
			}(v)
			b, s := checkRule(gcfg.Instance().MustGet(ctx, v.opt.Pattern).Interface(), v.Env, v.opt)
			testWithExcel.Assert(v.Name, fmt.Sprintf("%t,%s", b, s), v.Expect)
		})
	}
}
