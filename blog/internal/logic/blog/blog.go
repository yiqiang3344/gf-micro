package blog

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
	"yijunqiang/gf-micro/blog/internal/logging"
	"yijunqiang/gf-micro/blog/internal/model/entity"
	rocketmq_client "yijunqiang/gf-micro/blog/internal/utility/mq"
	"yijunqiang/gf-micro/blog/internal/utility/mstring"

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

func (s *sBlog) Create(ctx context.Context, title string, content string, nickname string) (blog *entity.Blog, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
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
	return
}

func (s *sBlog) Edit(ctx context.Context, id uint64, title string, content string, nickname string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
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
	return
}

func (s *sBlog) GetById(ctx context.Context, id uint64) (*pbentity.Blog, error) {
	var blog *pbentity.Blog
	err := dao.Blog.Ctx(ctx).Where(do.Blog{
		Id: id,
	}).Scan(&blog)
	return blog, err
}

func (s *sBlog) GetList(ctx context.Context) (list []*pbentity.Blog, err error) {
	list = []*pbentity.Blog{}
	err = dao.Blog.Ctx(ctx).Scan(&list)
	if err != nil {
		return
	}
	return
}

func (s *sBlog) Delete(ctx context.Context, id uint64) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
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

	//写入消息队列
	mqClient, err := rocketmq_client.GetClient(&rocketmq_client.ClientCfg{
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
	})
	if err != nil {
		g.Log().Debugf(ctx, "消息队列客户端初始化失败：%+v", err)
		return
	}
	err = mqClient.StartProducer(
		ctx,
		rocketmq_client.WithProducerOptionTopics([]string{BatDeleteTopic}),
		rocketmq_client.WithProducerOptionMaxAttempts(2),
	)
	if err != nil {
		g.Log().Debugf(ctx, "消息队列生产者初始化失败：%+v", err)
		return
	}
	batNo = mstring.RandomString(10)
	// todo 生成进度缓存

	defer mqClient.StopProducer()
	for _, id := range ids {
		_, err1 := mqClient.Send(ctx, rocketmq_client.TopicNormal, rocketmq_client.Message{
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

// BatDeleteConsumer 删除博客消费逻辑
func (s *sBlog) BatDeleteConsumer(ctx context.Context) (stopFunc func(), err error) {
	mqClient, err := rocketmq_client.GetClient(&rocketmq_client.ClientCfg{
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
	})
	if err != nil {
		g.Log().Debugf(ctx, "消息队列客户端初始化失败：%v", err)
		return nil, gerror.NewCodef(gcode.CodeInternalError, "消息队列客户端初始化失败:%v", err)
	}
	stopFunc, err = mqClient.SimpleConsume(
		ctx,
		func(ctx context.Context, msg *rmq_client.MessageView, consumer rmq_client.SimpleConsumer) {
			id := string(msg.GetBody())
			_, err1 := dao.Blog.Ctx(ctx).Where(do.Blog{
				Id: id,
			}).Delete()
			//todo 更新进度
			if err1 != nil {
				g.Log().Debugf(ctx, "删除博客[%s]失败:%s", id, err1)
				return
			}
			err1 = consumer.Ack(ctx, msg)
			if err1 != nil {
				g.Log().Debugf(ctx, "删除博客ACK[%s]失败:%s", id, err1)
			} else {
				g.Log().Debugf(ctx, "删除博客ACK[%s]成功", id)
			}
			return
		},
		rocketmq_client.WithConsumerOptionAwaitDuration(5*time.Second),
		rocketmq_client.WithConsumerOptionInvisibleDuration(10*time.Second),
		rocketmq_client.WithConsumerOptionSubExpressions(map[string]*rmq_client.FilterExpression{
			BatDeleteTopic: rmq_client.SUB_ALL,
		}),
	)
	if err != nil {
		return nil, gerror.NewCodef(gcode.CodeInternalError, "消费异常:%v", err)
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
