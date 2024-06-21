package blog

import (
	"context"
	"errors"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/yiqiang3344/gf-micro/cache"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/dao"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/model/do"
	rocketmq_client "github.com/yiqiang3344/rocketmq-client-go"
	"time"
)

const BatDeleteConsumerGroup = "blog_bat_delete_consumer"

func (s *sBlog) BatDeleteConsumer(ctx context.Context, parser *gcmd.Parser) (stopFunc func(), err error) {
	stopFunc, err = rocketmq_client.SimpleConsume4Gf(
		ctx,
		&rocketmq_client.Config{
			Endpoint:      g.Cfg().MustGet(ctx, "rocketmq.endpoint").String(),
			NameSpace:     g.Cfg().MustGet(ctx, "rocketmq.namespace").String(),
			AccessKey:     g.Cfg().MustGet(ctx, "rocketmq.accessKey").String(),
			AccessSecret:  g.Cfg().MustGet(ctx, "rocketmq.accessSecret").String(),
			ConsumerGroup: BatDeleteConsumerGroup,
			LogPath:       g.Cfg().MustGet(ctx, "rocketmq.logPath").String(),
			LogStdout:     g.Cfg().MustGet(ctx, "rocketmq.logStdout").Bool(),
			Debug:         g.Cfg().MustGet(ctx, "rocketmq.debug").Bool(),
			DebugHandlerFunc: func(msg string) {
				g.Log().Debug(ctx, msg)
			},
		},
		func(ctx context.Context, msg *rmq_client.MessageView, consumer rocketmq_client.Consumer) error {
			id := string(msg.GetBody())
			_, err1 := dao.Blog.Ctx(ctx).Where(do.Blog{
				Id: id,
			}).Delete()
			if err1 != nil {
				g.Log().Debugf(ctx, "删除博客[%s]失败:%s", id, err1)
				return err1
			}
			//清除缓存
			_, err1 = g.DB().GetCache().Remove(
				ctx,
				cache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%s", id)),
				cache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
			)
			if err1 != nil {
				g.Log().Debugf(ctx, "博客[%s]清除缓存失败:%s", id, err1)
				return err1
			}
			//更新进度
			batNo, ok := msg.GetProperties()["batNo"]
			if !ok {
				err1 = errors.New("batNo不存在")
				g.Log().Debugf(ctx, "删除博客[%s]失败:%s", id, err1)
				return err1
			}
			ret, err1 := g.Redis().Decr(ctx, getBatDeleteProgressKey(batNo))
			if err1 != nil {
				g.Log().Debugf(ctx, "删除博客[%s]失败:%s", id, err1)
				return err1
			}
			g.Log().Debugf(ctx, "[%s]删除博客进度:剩余[%d]", batNo, ret)
			err1 = consumer.Ack(ctx)
			if err1 != nil {
				g.Log().Debugf(ctx, "删除博客ACK[%s]失败:%s", id, err1)
				return err1
			}
			g.Log().Debugf(ctx, "删除博客ACK[%s]成功", id)
			return nil
		},
		rocketmq_client.WithConsumerOptionAwaitDuration(5*time.Second),
		rocketmq_client.WithConsumerOptionInvisibleDuration(10*time.Second),
		rocketmq_client.WithConsumerOptionSubExpressions(map[string]*rmq_client.FilterExpression{
			BatDeleteTopic: rmq_client.SUB_ALL,
		}),
	)
	if err != nil {
		g.Log().Debugf(ctx, "消息队列消费者初始化失败：%v", err)
		return nil, gerror.NewCodef(gcode.CodeInternalError, "消息队列消费者初始化失败:%v", err)
	}
	return
}
