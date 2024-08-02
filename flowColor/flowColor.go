package flowColor

import (
	"context"
	"github.com/gogf/gf/v2/os/gcfg"
	"google.golang.org/grpc/metadata"
	"os"
)

const ColorBase = "base"
const FlowColor = "FLOW_COLOR"

func IsOpen() bool {
	return gcfg.Instance().MustGet(context.Background(), "flowColor.open").Bool()
}

func GetLocalFlowColor() *string {
	//根据环境标识获取颜色标识
	s := os.Getenv(FlowColor)
	return &s
}

func IsBase() *bool {
	ret := false
	if *GetLocalFlowColor() == ColorBase {
		ret = true
	}
	return &ret
}

func SetCtxFlowColor(ctx context.Context, defaultColor string) context.Context {
	if *GetCtxFlowColor(ctx) == "" {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			"color": defaultColor,
		}))
		ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
			"color": defaultColor,
		}))
	}
	return ctx
}

func GetCtxFlowColor(ctx context.Context) *string {
	var color string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ret := md.Get("color")
		if len(ret) > 0 {
			color = ret[0]
		}
	}
	return &color
}
