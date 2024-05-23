package controller

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	v1 "web/api/blog/v1"
	"web/internal/model"
	blogMicroV1 "yijunqiang/gf-micro/blog/api/blog/v1"
)

var (
	blogConn   = grpcx.Client.MustNewGrpcClientConn("blog")
	blogClient = blogMicroV1.NewBlogClient(blogConn)
)

func (c *cBlog) BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error) {
	//todo 获取登录用户信息
	nickname := "test1"
	res = &v1.BlogCreateRes{}
	_, err = blogClient.Create(ctx, &blogMicroV1.CreateReq{
		Nickname: nickname,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (c *cBlog) BlogEdit(ctx context.Context, req *v1.BlogEditReq) (res *v1.BlogEditRes, err error) {
	//todo 获取登录用户信息
	nickname := "test1"
	res = &v1.BlogEditRes{}
	_, err = blogClient.Edit(ctx, &blogMicroV1.EditReq{
		Id:       gconv.Uint64(req.Id),
		Nickname: nickname,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (c *cBlog) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	res = &v1.BlogDetailRes{}
	ret, err := blogClient.GetOne(ctx, &blogMicroV1.GetOneReq{
		Id: gconv.Uint64(req.Id),
	})
	if err != nil {
		return
	}
	if ret.Blog == nil {
		err = gerror.NewCode(gcode.New(-1, "博客不存在", nil))
		return
	}
	res.BlogDetailOutput = &model.BlogDetailOutput{
		Id:       ret.Blog.Id,
		Title:    ret.Blog.Title,
		Content:  ret.Blog.Content,
		Nickname: ret.Blog.Nickname,
	}
	return
}

func (c *cBlog) BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error) {
	res = &v1.BlogListRes{}
	ret, err := blogClient.GetList(ctx, &blogMicroV1.GetListReq{})
	if err != nil {
		return
	}
	for _, v := range ret.List {
		res.List = append(res.List, &model.BlogDetailOutput{
			Id:       v.Id,
			Title:    v.Title,
			Content:  v.Content,
			Nickname: v.Nickname,
		})
	}
	return
}

func (c *cBlog) BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error) {
	res = &v1.BlogDeleteRes{}
	_, err = blogClient.Delete(ctx, &blogMicroV1.DeleteReq{
		Id: gconv.Uint64(req.Id),
	})
	if err != nil {
		return
	}
	return
}

func (c *cBlog) BlogBatDelete(ctx context.Context, req *v1.BlogBatDeleteReq) (res *v1.BlogBatDeleteRes, err error) {
	var ids []uint64
	res = &v1.BlogBatDeleteRes{}
	for _, v := range strings.Split(req.Ids, ",") {
		ids = append(ids, gconv.Uint64(v))
	}
	ret, err := blogClient.BatDelete(ctx, &blogMicroV1.BatDeleteReq{
		Ids: ids,
	})
	if err != nil {
		return
	}
	res.BatNo = ret.BatNo
	return
}

func (c *cBlog) BlogGetBatDeleteStatus(ctx context.Context, req *v1.BlogGetBatDeleteStatusReq) (res *v1.BlogGetBatDeleteStatusRes, err error) {
	res = &v1.BlogGetBatDeleteStatusRes{}
	ret, err := blogClient.GetBatDeleteStatus(ctx, &blogMicroV1.GetBatDeleteStatusReq{
		BatNo: req.BatNo,
	})
	if err != nil {
		return
	}
	res.Status = ret.Status
	return
}
