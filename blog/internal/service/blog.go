// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"yijunqiang/gf-micro/blog/api/pbentity"
	"yijunqiang/gf-micro/blog/internal/model/entity"
)

type (
	IBlog interface {
		Create(ctx context.Context, title string, content string, nickname string) (*entity.Blog, error)
		Edit(ctx context.Context, id uint64, title string, content string, nickname string) (err error)
		GetById(ctx context.Context, id uint64) (*pbentity.Blog, error)
		GetList(ctx context.Context) (list []*pbentity.Blog, err error)
		Delete(ctx context.Context, id uint64) (err error)
		BatDelete(ctx context.Context, ids []uint64) (batNo string, err error)
		GetBatDeleteStatus(ctx context.Context, batNo string) (status string, err error)
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
