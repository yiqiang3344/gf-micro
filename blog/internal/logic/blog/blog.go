package blog

import (
	"context"
	"errors"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	rocketmq_client "github.com/yiqiang3344/rocketmq-client-go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"yijunqiang/gf-micro/blog/api/pbentity"
	"yijunqiang/gf-micro/blog/internal/dao"
	"yijunqiang/gf-micro/blog/internal/logging"
	"yijunqiang/gf-micro/blog/internal/model/do"
	"yijunqiang/gf-micro/blog/internal/model/entity"
	"yijunqiang/gf-micro/blog/internal/service"
	"yijunqiang/gf-micro/blog/internal/utility/mcache"
)

type (
	sBlog struct{}
)

func init() {
	service.RegisterBlog(&sBlog{})
}

func (s *sBlog) Create(ctx context.Context, title string, content string, nickname string) (blog *entity.Blog, err error) {
	defer func() {
		if err != nil {
			logging.BizLog{
				Tag:     "Create",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Create",
				Message: "success",
			}.Log(ctx)
		}
	}()
	blog = &entity.Blog{
		Title:    title,
		Content:  content,
		Nickname: nickname,
	}
	_, err = dao.Blog.Ctx(ctx).Data(blog).Insert()
	if err != nil {
		return
	}
	//清除缓存
	_, err = g.DB().GetCache().Remove(
		ctx,
		mcache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
	)
	return
}

func (s *sBlog) Edit(ctx context.Context, id uint64, title string, content string, nickname string) (err error) {
	defer func() {
		if err != nil {
			logging.BizLog{
				Tag:     "Edit",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Edit",
				Message: "success",
			}.Log(ctx)
		}
	}()
	blog, err := s.GetById(ctx, id)
	if err != nil {
		return
	}
	if blog == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "博客不存在")
		return
	}
	if blog.Nickname != nickname {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "只能编辑自己的博客")
		return
	}
	_, err = dao.Blog.Ctx(ctx).Where(do.Blog{
		Id: id,
	}).Data(do.Blog{
		Title:   title,
		Content: content,
	}).Update()
	if err != nil {
		return
	}
	//清除缓存
	_, err = g.DB().GetCache().Remove(
		ctx,
		mcache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		mcache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
	)
	if err != nil {
		return
	}
	return
}

func (s *sBlog) GetById(ctx context.Context, id uint64) (ret *pbentity.Blog, err error) {
	var blog *entity.Blog
	err = dao.Blog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     mcache.GetDbCacheKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
	}).Where(do.Blog{
		Id: id,
	}).Scan(&blog)
	if err != nil {
		return
	}
	if blog != nil {
		ret = &pbentity.Blog{}
		gconv.ConvertWithRefer(blog, ret)
		ret.CreateAt = timestamppb.New(blog.CreateAt.Time)
		ret.UpdateAt = timestamppb.New(blog.UpdateAt.Time)
	}
	return
}

func (s *sBlog) GetList(ctx context.Context) (list []*pbentity.Blog, err error) {
	var listRet []entity.Blog
	err = dao.Blog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     mcache.GetDbCacheKey(dao.Blog.Table(), "GetList"),
	}).Scan(&listRet)
	if err != nil {
		return
	}
	for _, blog := range listRet {
		b := &pbentity.Blog{}
		gconv.ConvertWithRefer(blog, b)
		b.CreateAt = timestamppb.New(blog.CreateAt.Time)
		b.UpdateAt = timestamppb.New(blog.UpdateAt.Time)
		list = append(list, b)
	}

	return
}

func (s *sBlog) Delete(ctx context.Context, id uint64) (err error) {
	defer func() {
		if err != nil {
			logging.BizLog{
				Tag:     "Delete",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Delete",
				Message: "success",
			}.Log(ctx)
		}
	}()
	blog, err := s.GetById(ctx, id)
	if err != nil {
		return
	}
	if blog == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "博客不存在")
		return
	}
	_, err = dao.Blog.Ctx(ctx).Where(do.Blog{
		Id: id,
	}).Delete()
	if err != nil {
		return
	}
	//清除缓存
	_, err = g.DB().GetCache().Remove(
		ctx,
		mcache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		mcache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
	)
	if err != nil {
		return
	}
	return
}

const BatDeleteTopic = "blog_bat_delete"

func getBatDeleteProgressKey(batNo string) string {
	return fmt.Sprintf("progress:BatDelete:%s", batNo)
}

func (s *sBlog) BatDelete(ctx context.Context, ids []uint64) (batNo string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "BatDelete",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "BatDelete",
				Message: "success",
			}.Log(ctx)
		}
	}()

	//写入消息队列
	producer, err := rocketmq_client.GetGfProducer(
		&rocketmq_client.Config{
			Endpoint:     g.Cfg().MustGet(ctx, "rocketmq.endpoint").String(),
			NameSpace:    g.Cfg().MustGet(ctx, "rocketmq.namespace").String(),
			AccessKey:    g.Cfg().MustGet(ctx, "rocketmq.accessKey").String(),
			AccessSecret: g.Cfg().MustGet(ctx, "rocketmq.accessSecret").String(),
			LogPath:      g.Cfg().MustGet(ctx, "rocketmq.logPath").String(),
			LogStdout:    g.Cfg().MustGet(ctx, "rocketmq.logStdout").Bool(),
			Debug:        g.Cfg().MustGet(ctx, "rocketmq.debug").Bool(),
			DebugHandlerFunc: func(msg string) {
				g.Log().Debug(ctx, msg)
			},
		},
		rocketmq_client.WithProducerOptionTopics(BatDeleteTopic),
		rocketmq_client.WithProducerOptionMaxAttempts(3),
	)
	if err != nil {
		g.Log().Debugf(ctx, "消息队列生产者初始化失败：%v", err)
		return
	}
	defer producer.Stop()

	//生成批次号，及初始化进度信息
	batNo = guid.S()
	idsMap := gmap.NewListMap()
	for _, v := range ids {
		idsMap.Set(v, 1)
	}
	err = g.Redis().SetEX(ctx, getBatDeleteProgressKey(batNo), idsMap.Size(), int64(time.Hour.Seconds()))
	if err != nil {
		return
	}
	for id, _ := range idsMap.Keys() {
		_, err1 := producer.Send(ctx, rocketmq_client.TopicNormal, rocketmq_client.Message{
			Topic: BatDeleteTopic,
			Keys: []string{
				batNo,
			},
			Body: gconv.String(id),
			Properties: map[string]string{
				"batNo": batNo,
			},
		})
		if err1 != nil {
			g.Log().Debugf(ctx, "删除博客[%d]消息队列生产失败:%v", id, err1)
		} else {
			g.Log().Debugf(ctx, "删除博客[%d]消息队列生产成功", id)
		}
	}

	return
}

const BatDeleteConsumerGroup = "blog_bat_delete_consumer"

func (s *sBlog) BatDeleteConsumer(ctx context.Context) (stopFunc func(), err error) {
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
				mcache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%s", id)),
				mcache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
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

func (s *sBlog) GetBatDeleteStatus(ctx context.Context, batNo string) (status string, err error) {
	ret, err := g.Redis().Get(ctx, getBatDeleteProgressKey(batNo))
	if err != nil {
		return
	}
	if ret.IsNil() {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "批次不存在或已超过有效期")
		return
	}
	status = "pending"
	if ret.Int() <= 0 {
		status = "success"
	}
	return
}
