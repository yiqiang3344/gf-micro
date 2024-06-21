// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/yiqiang3344/gf-micro/example/blog/api/pbentity"
	"github.com/yiqiang3344/gf-micro/example/blog/internal/model/entity"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/xxl-job/xxl-job-executor-go"
)

type (
	IBlog interface {
		Create(ctx context.Context, title string, content string, nickname string) (blog *entity.Blog, err error)
		Edit(ctx context.Context, id uint64, title string, content string, nickname string) (err error)
		GetById(ctx context.Context, id uint64) (ret *pbentity.Blog, err error)
		GetList(ctx context.Context) (list []*pbentity.Blog, err error)
		Delete(ctx context.Context, id uint64) (err error)
		BatDelete(ctx context.Context, ids []uint64) (batNo string, err error)
		GetBatDeleteStatus(ctx context.Context, batNo string) (status string, err error)
		BatDeleteConsumer(ctx context.Context, parser *gcmd.Parser) (stopFunc func(), err error)
		Stats(ctx context.Context, param *xxl.RunReq) (msg string)
	}
)

var (
	localBlog IBlog
)

func Blog() IBlog {
	if localBlog == nil {
		panic("implement not found for interface IBlog, forgot register?")
	}
	return localBlog
}

func RegisterBlog(i IBlog) {
	localBlog = i
}
