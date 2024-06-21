package controller

import (
	"context"
	v1 "web/api/blog/v1"
	"web/internal/service"
)

func (c *cBlog) BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error) {
	res, err = service.Blog().BlogCreate(ctx, req)
	return
}

func (c *cBlog) BlogEdit(ctx context.Context, req *v1.BlogEditReq) (res *v1.BlogEditRes, err error) {
	res, err = service.Blog().BlogEdit(ctx, req)
	return
}

func (c *cBlog) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	res, err = service.Blog().BlogDetail(ctx, req)
	return
}

func (c *cBlog) BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error) {
	res, err = service.Blog().BlogList(ctx, req)
	return
}

func (c *cBlog) BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error) {
	res, err = service.Blog().BlogDelete(ctx, req)
	return
}

func (c *cBlog) BlogBatDelete(ctx context.Context, req *v1.BlogBatDeleteReq) (res *v1.BlogBatDeleteRes, err error) {
	res, err = service.Blog().BlogBatDelete(ctx, req)
	return
}

func (c *cBlog) BlogGetBatDeleteStatus(ctx context.Context, req *v1.BlogGetBatDeleteStatusReq) (res *v1.BlogGetBatDeleteStatusRes, err error) {
	res, err = service.Blog().BlogGetBatDeleteStatus(ctx, req)
	return
}
