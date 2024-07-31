package client

import (
	"context"
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcfg"
	cfg2 "github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/flowColor"
	"github.com/yiqiang3344/gf-micro/flowColor/balancer"
	"github.com/yiqiang3344/gf-micro/logging"
)

// GetHttpClient 获取http客户端，封装了流量染色，client日志
func GetHttpClient(optsFunc ...OptsFunc) *gclient.Client {
	o := &opts{
		log:           true,  //默认打日志
		httpDiscovery: false, //默认不开启服务注册发现
	}
	if len(optsFunc) > 0 {
		for _, v := range optsFunc {
			v(o)
		}
	}
	client := g.Client()
	if flowColor.IsOpen() {
		// 如果开启流量染色，则必须开启服务注册发现
		o.httpDiscovery = true
		client.SetBuilder(balancer.NewBuilderFlowColor())
		client.Use(flowColor.HttpClientMiddleware)
	}
	if o.httpDiscovery {
		if gsvc.GetRegistry() == nil {
			cfg := gcfg.Instance().MustGet(context.Background(), cfg2.REGISTRY_GRPC)
			if cfg.IsNil() {
				panic(fmt.Errorf("gscv服务注册发现未配置"))
			}
			client.SetDiscovery(etcd.New(cfg.String()))
		}
	} else {
		client.SetDiscovery(nil)
	}
	if o.log {
		client.Use(logging.HttpClientLogMiddleware)
	}
	return client
}
