package cfg

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestTest(t *testing.T) {
	g.DumpWithType(map[string]interface{}{"1": "1"}["2"])
}
