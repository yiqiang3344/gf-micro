package flowColor

import (
	"context"
	"google.golang.org/grpc"
)

// GrpcServerUnary grpc服务端流量染色拦截器
func GrpcServerUnary(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	if IsOpen() {
		ctx = SetCtxFlowColor(ctx, ColorBase)
	}
	return handler(ctx, req)
}

// GrpcClientUnary grpc客户端流量染色拦截器
func GrpcClientUnary(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if IsOpen() {
		ctx = SetCtxFlowColor(ctx, ColorBase)
	}
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
