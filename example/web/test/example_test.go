package test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func TestExample(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		caseName := "case1"
		Assert(caseName, "", "")
	})
}
