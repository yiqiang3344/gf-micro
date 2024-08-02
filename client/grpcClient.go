package client

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcfg"
	cfg2 "github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/flowColor"
	"github.com/yiqiang3344/gf-micro/flowColor/balancer"
	"github.com/yiqiang3344/gf-micro/logging"
	"google.golang.org/grpc"
)

var grpcClientMap map[string]interface{}

// GetGrpcClient 获取grpc客户端，封装了流量染色功能和client日志
func GetGrpcClient[T any](name string, f func(cc grpc.ClientConnInterface) T, optsFunc ...OptsFunc) T {
	o := &opts{
		log: true, //默认打日志
	}
	if len(optsFunc) > 0 {
		for _, v := range optsFunc {
			v(o)
		}
	}
	if gsvc.GetRegistry() == nil {
		cfg := gcfg.Instance().MustGet(context.Background(), cfg2.REGISTRY_GRPC)
		if cfg.IsNil() {
			panic(fmt.Errorf("grpc服务注册发现未配置"))
		}
		grpcx.Resolver.Register(etcd.New(cfg.String()))
	}
	if grpcClientMap == nil {
		grpcClientMap = make(map[string]interface{})
	}
	if v, ok := grpcClientMap[name]; ok {
		return v.(T)
	}
	var (
		u []grpc.UnaryClientInterceptor
		d []grpc.DialOption
	)
	if o.log {
		u = append(u, logging.GrpcClientLoggerUnary)
	}
	u = append(u, flowColor.GrpcClientUnary)
	d = append(d, grpcx.Client.ChainUnary(u...))
	if flowColor.IsOpen() {
		d = append(d, balancer.WithFlowColor())
	}
	c := f(grpcx.Client.MustNewGrpcClientConn(name, d...))
	grpcClientMap[name] = c
	return c
}
