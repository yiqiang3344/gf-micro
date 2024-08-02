package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"github.com/yiqiang3344/gf-micro/cfg"
	"github.com/yiqiang3344/gf-micro/logging"
)

type MiddlewareForCmd func(ctx context.Context) (stopFunc func())

// GetGrpcMiddleware grpc服务命令全局中间件
func GetGrpcMiddleware() MiddlewareForCmd {
	return func(ctx context.Context) (stopFunc func()) {
		var shutdownArr []func()

		//配置中心
		if !gcfg.Instance().MustGet(ctx, "apollo").IsNil() {
			adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
			if err != nil {
				panic(err)
			}
			gcfg.Instance().SetAdapter(adapter)
		}

		//日志json化
		glog.SetDefaultHandler(logging.HandlerJson)

		// 链路追踪初始化
		if !gcfg.Instance().MustGet(ctx, cfg.OTLP).IsNil() {
			shutdown, err := otlpgrpc.Init(
				gcfg.Instance().MustGet(ctx, cfg.APPNAME).String(),
				gcfg.Instance().MustGet(ctx, "otlp.endpoint").String(),
				gcfg.Instance().MustGet(ctx, "otlp.traceToken").String(),
			)
			if err != nil {
				g.Log().Debugf(ctx, "otlp初始化失败:%v\n", err)
			} else {
				shutdownArr = append(shutdownArr, shutdown)
			}
		}

		// grpc服务注册发现
		if !gcfg.Instance().MustGet(ctx, cfg.REGISTRY_GRPC).IsNil() {
			grpcx.Resolver.Register(etcd.New(gcfg.Instance().MustGet(ctx, cfg.REGISTRY_GRPC).String()))
		}

		// 设置db全局redis缓存配置
		if !gcfg.Instance().MustGet(ctx, "redis.db").IsNil() {
			g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis("db")))
		}

		return func() {
			for _, v := range shutdownArr {
				v()
			}
			return
		}
	}
}

// GetHttpMiddleware http服务命令全局中间件
func GetHttpMiddleware() MiddlewareForCmd {
	return func(ctx context.Context) (stopFunc func()) {
		var shutdownArr []func()

		//配置中心
		if !gcfg.Instance().MustGet(ctx, "apollo").IsNil() {
			adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
			if err != nil {
				panic(err)
			}
			gcfg.Instance().SetAdapter(adapter)
		}

		// 链路追踪初始化
		if !gcfg.Instance().MustGet(ctx, cfg.OTLP).IsNil() {
			shutdown, err := otlpgrpc.Init(
				gcfg.Instance().MustGet(ctx, cfg.APPNAME).String(),
				gcfg.Instance().MustGet(ctx, "otlp.endpoint").String(),
				gcfg.Instance().MustGet(ctx, "otlp.traceToken").String(),
			)
			if err != nil {
				g.Log().Debugf(ctx, "otlp初始化失败:%v\n", err)
			} else {
				shutdownArr = append(shutdownArr, shutdown)
			}
		}

		// ghttp服务注册发现
		if !gcfg.Instance().MustGet(ctx, cfg.REGISTRY_HTTP).IsNil() {
			gsvc.SetRegistry(etcd.New(gcfg.Instance().MustGet(ctx, cfg.REGISTRY_HTTP).String()))
		}

		// 设置db全局redis缓存配置
		if !gcfg.Instance().MustGet(ctx, "redis.db").IsNil() {
			g.DB().GetCache().SetAdapter(gcache.NewAdapterRedis(g.Redis("db")))
		}

		return func() {
			for _, v := range shutdownArr {
				v()
			}
			return
		}
	}
}
