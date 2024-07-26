package flowColor

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const ColorBase = "base"

func SetCtxFlowColor(ctx context.Context, color string) context.Context {
	if GetCtxFlowColor(ctx) == "" {
		ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
			"color": color,
		}))
		ctx = metadata.NewIncomingContext(ctx, metadata.New(map[string]string{
			"color": color,
		}))
	}
	return ctx
}

func GetCtxFlowColor(ctx context.Context) string {
	var color string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ret := md.Get("color")
		if len(ret) > 0 {
			color = ret[0]
		}
	}
	return color
}
