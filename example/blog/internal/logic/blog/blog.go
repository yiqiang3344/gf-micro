package blog

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/yiqiang3344/gf-micro/cache"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/dao"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/logging"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/model/do"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/model/entity"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/service"
	"github.com/yiqiang3344/gf-micro/flowColor"
	rocketmq_client "github.com/yiqiang3344/rocketmq-client-go"
	"time"
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
		cache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
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
		cache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		cache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
	)
	if err != nil {
		return
	}
	return
}

func (s *sBlog) GetById(ctx context.Context, id uint64) (blog *entity.Blog, err error) {
	err = dao.Blog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     cache.GetDbCacheKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
	}).Where(do.Blog{
		Id: id,
	}).Scan(&blog)
	if err != nil {
		return
	}
	return
}

func (s *sBlog) GetList(ctx context.Context) (list []*entity.Blog, err error) {
	err = dao.Blog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     cache.GetDbCacheKey(dao.Blog.Table(), "GetList"),
	}).Scan(&list)
	if err != nil {
		return
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
		cache.GetDbCacheFullKey(dao.Blog.Table(), fmt.Sprintf("GetById:%d", id)),
		cache.GetDbCacheFullKey(dao.Blog.Table(), "GetList"),
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

	if len(ids) == 0 {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "ids不能为空")
		return
	}

	idsList := garray.NewSortedIntArray()
	for _, v := range ids {
		//检查博客是否存在
		blog, err1 := s.GetById(ctx, v)
		if err1 != nil {
			return "", err1
		}
		if blog == nil {
			return "", gerror.NewCode(gcode.CodeBusinessValidationFailed, fmt.Sprintf("博客[%d]不存在", v))
		}
		idsList.Add(gconv.Int(v))
	}

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
			FlowColor: flowColor.GetCtxFlowColor(ctx),
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
	err = g.Redis().SetEX(ctx, getBatDeleteProgressKey(batNo), idsList.Len(), int64(time.Hour.Seconds()))
	if err != nil {
		return "", err
	}
	for _, id := range idsList.Slice() {
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
