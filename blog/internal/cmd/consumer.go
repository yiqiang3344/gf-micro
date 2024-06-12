package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gproc"
	"os"
	"yijunqiang/gf-micro/blog/internal/logging"
	"yijunqiang/gf-micro/blog/internal/service"
)

var (
	batDeleteBlogConsumer = &gcmd.Command{
		Name:  "batDeleteBlogConsumer",
		Usage: "./main batDeleteBlogConsumer",
		Brief: "批量删除博客的消费者",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// grpc服务注册发现
			grpcx.Resolver.Register(etcd.New(gcfg.Instance().MustGet(ctx, "registry.etcd").String()))

			//日志json化
			glog.SetDefaultHandler(logging.HandlerJson)

			// 链路追踪初始化
			if !gcfg.Instance().MustGet(ctx, "otlp").IsNil() {
				shutdown, err := otlpgrpc.Init(
					gcfg.Instance().MustGet(ctx, "appName").String(),
					gcfg.Instance().MustGet(ctx, "otlp.endpoint").String(),
					gcfg.Instance().MustGet(ctx, "otlp.traceToken").String(),
				)
				if err != nil {
					g.Log().Debugf(ctx, "otlp初始化失败:%v\n", err)
				} else {
					defer shutdown()
				}
			}

			stopFunc, err := service.Blog().BatDeleteConsumer(ctx)
			if err != nil {
				g.Log("debug").Errorf(ctx, "batDeleteBlogConsumer异常:%+v", err)
				panic(err)
			}

			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				stopFunc()
			})
			gproc.Listen()
			return
		},
	}
)
