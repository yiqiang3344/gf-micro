// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "web/api/blog/v1"
)

type (
	IBlog interface {
		BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error)
		BlogEdit(ctx context.Context, req *v1.BlogEditReq) (res *v1.BlogEditRes, err error)
		BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error)
		BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error)
		BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error)
		BlogBatDelete(ctx context.Context, req *v1.BlogBatDeleteReq) (res *v1.BlogBatDeleteRes, err error)
		BlogGetBatDeleteStatus(ctx context.Context, req *v1.BlogGetBatDeleteStatusReq) (res *v1.BlogGetBatDeleteStatusRes, err error)
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
