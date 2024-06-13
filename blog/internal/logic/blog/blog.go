package blog

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"yijunqiang/gf-micro/blog/internal/logging"
	"yijunqiang/gf-micro/blog/internal/model/entity"
	"yijunqiang/gf-micro/blog/internal/utility/mstring"

	rocketmq_client "github.com/yiqiang3344/rocketmq-client-go"
	"yijunqiang/gf-micro/blog/api/pbentity"
	"yijunqiang/gf-micro/blog/internal/dao"
	"yijunqiang/gf-micro/blog/internal/model/do"
	"yijunqiang/gf-micro/blog/internal/service"
)

type (
	sBlog struct{}
)

func init() {
	service.RegisterBlog(&sBlog{})
}

func getDbCacheFullKey(table string, key string) string {
	//gf/database/gdb/gdb.go:386中有定义缓存前缀，但没开放出来
	return fmt.Sprintf("SelectCache:%s@%s", table, key)
}

func getDbCacheKey(table string, key string) string {
	return fmt.Sprintf("%s@%s", table, key)
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
		getDbCacheFullKey(dao.Blog.Table(), "GetList"),
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
		getDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		getDbCacheFullKey(dao.Blog.Table(), "GetList"),
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
		Name:     getDbCacheKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
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
		Name:     getDbCacheKey(dao.Blog.Table(), "GetList"),
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
		getDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		getDbCacheFullKey(dao.Blog.Table(), "GetList"),
	)
	if err != nil {
		return
	}
	return
}

const BatDeleteTopic = "blog_bat_delete"

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
	batNo = mstring.RandomString(10)

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
	for _, id := range ids {
		_, err1 := producer.Send(ctx, rocketmq_client.TopicNormal, rocketmq_client.Message{
			Topic: BatDeleteTopic,
			Keys: []string{
				batNo,
			},
			Body: gconv.String(id),
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
			//todo 更新进度
			//清除缓存
			_, err1 = g.DB().GetCache().Remove(
				ctx,
				getDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%s", id)),
				getDbCacheFullKey(dao.Blog.Table(), "GetList"),
			)
			if err1 != nil {
				return err1
			}
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
	// todo 获取进度
	status = ""
	switch batNo {
	case "bat1":
		status = "success"
	case "bat2":
		status = "pending"
	default:
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "batNo不存在")
		return
	}
	return
}
