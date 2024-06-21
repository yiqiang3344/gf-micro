package logging

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

// UnaryLogger is the default unary interceptor for logging purpose.
func UnaryLogger(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	var (
		start    = time.Now()
		res, err = handler(ctx, req)
		duration = time.Since(start)
	)
	handleAccessLog(ctx, err, duration, info, req, res)
	handleErrorLog(ctx, err, info, req)
	return res, err
}

// handleAccessLog handles the access logging for server.
func handleAccessLog(
	ctx context.Context, err error, duration time.Duration, info *grpc.UnaryServerInfo, req, res interface{},
) {
	GrpcAccessLog{
		Method: info.FullMethod,
		Cost:   fmt.Sprintf("%.3fms", float64(duration)/1e6),
		Req:    req,
		Res:    res,
	}.Log(ctx, err)
}

// handleErrorLog handles the error logging for server.
func handleErrorLog(ctx context.Context, err error, info *grpc.UnaryServerInfo, req interface{}) {
	if err == nil {
		return
	}
	ErrorLog{
		Method: info.FullMethod,
		Req:    req,
	}.Log(ctx, err)
}
