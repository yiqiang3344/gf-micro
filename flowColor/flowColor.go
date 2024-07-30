package flowColor

import (
	"context"
	"google.golang.org/grpc/metadata"
	"os"
)

const ColorBase = "base"
const FlowColor = "FLOW_COLOR"

func GetLocalFlowColor() *string {
	//根据环境标识获取颜色标识
	s := os.Getenv("FLOW_COLOR")
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

func Match(ctx context.Context) *bool {
	ret := false
	//流量染色标识为空，当前系统为基准系统则返回true
	if *GetCtxFlowColor(ctx) == "" && *GetLocalFlowColor() == ColorBase {
		ret = true
	}
	//流量染色标识不为空，且与当前系统颜色标识一致则返回true
	if *GetCtxFlowColor(ctx) != "" && *GetCtxFlowColor(ctx) == *GetLocalFlowColor() {
		ret = true
	}
	return &ret
}
