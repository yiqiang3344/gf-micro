package logging

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"time"
)

// GrpcServerLoggerUnary grpc服务端日志拦截器
func GrpcServerLoggerUnary(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	var (
		start    = time.Now()
		res, err = handler(ctx, req)
		duration = time.Since(start)
	)
	handleGrpcAccessLog(ctx, err, duration, info, req, res)
	handleGrpcErrorLog(ctx, err, info, req)
	return res, err
}

// GrpcClientLoggerUnary grpc客户端日志拦截器
func GrpcClientLoggerUnary(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(start)
	handleGrpcClientLog(ctx, method, duration, req.(proto.Message), reply.(proto.Message), err)
	return err
}

// handleGrpcAccessLog 处理服务端访问日志
func handleGrpcAccessLog(
	ctx context.Context, err error, duration time.Duration, info *grpc.UnaryServerInfo, req, res interface{},
) {
	GrpcAccessLog{
		Path: info.FullMethod,
		Cost: fmt.Sprintf("%.3fms", float64(duration)/1e6),
		Req:  req,
		Res:  res,
	}.Log(ctx, err)
}

// handleGrpcClientLog 处理客户端对外访问的日志
func handleGrpcClientLog(ctx context.Context, path string, duration time.Duration, req, res interface{}, err error) {
	GrpcClientLog{
		Path: path,
		Cost: fmt.Sprintf("%.3fms", float64(duration)/1e6),
		Req:  req,
		Res:  res,
	}.Log(ctx, err)
}

// handleGrpcErrorLog handles the error logging for server.
func handleGrpcErrorLog(ctx context.Context, err error, info *grpc.UnaryServerInfo, req interface{}) {
	if err == nil {
		return
	}
	ErrorLog{
		Method: info.FullMethod,
		Req:    req,
	}.Log(ctx, err)
}
