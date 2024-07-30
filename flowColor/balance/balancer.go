package balancer

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
)

// WithFlowColor returns a grpc.DialOption which enables flow color load balancing.
func WithFlowColor() grpc.DialOption {
	b := grpcx.Balancer
	flowColor := NewBuilderFlowColor()
	b.Register(flowColor)
	return b.WithName(flowColor.Name())
}
