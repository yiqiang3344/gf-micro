package blog

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/yiqiang3344/gf-micro/auth"
	blogMicroV1 "github.com/yiqiang3344/gf-micro/example/blog/api/blog/v1"
	"strings"
	v1 "web/api/blog/v1"
	"web/internal/logging"
	"web/internal/model"
	"web/internal/service"
)

type sBlog struct{}

func New() *sBlog {
	return &sBlog{}
}

func init() {
	service.RegisterBlog(New())
}

var blogClient blogMicroV1.BlogClient

func getBlogClient() blogMicroV1.BlogClient {
	if blogClient == nil {
		blogClient = blogMicroV1.NewBlogClient(grpcx.Client.MustNewGrpcClientConn("blog"))
	}
	return blogClient
}

func (c *sBlog) BlogCreate(ctx context.Context, req *v1.BlogCreateReq) (res *v1.BlogCreateRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "BlogCreate",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "BlogCreate",
				Message: "success",
			}.Log(ctx)
		}
	}()

	currentUser, err := auth.GetStandardAuth().GetUserWithCheck(ctx)
	if err != nil {
		return
	}
	res = &v1.BlogCreateRes{}
	_, err = getBlogClient().Create(ctx, &blogMicroV1.CreateReq{
		Nickname: currentUser.Nickname,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		return
	}
	return
}

func (c *sBlog) BlogEdit(ctx context.Context, req *v1.BlogEditReq) (res *v1.BlogEditRes, err error) {
	currentUser, err := auth.GetStandardAuth().GetUserWithCheck(ctx)
	if err != nil {
		return
	}
	res = &v1.BlogEditRes{}
	_, err = getBlogClient().Edit(ctx, &blogMicroV1.EditReq{
		Id:       req.Id,
		Nickname: currentUser.Nickname,
		Title:    req.Title,
		Content:  req.Content,
	})
	if err != nil {
		return
	}
	return
}

func (c *sBlog) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	res = &v1.BlogDetailRes{}
	ret, err := getBlogClient().GetOne(ctx, &blogMicroV1.GetOneReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	if ret.Blog == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "博客不存在")
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

func (c *sBlog) BlogList(ctx context.Context, req *v1.BlogListReq) (res *v1.BlogListRes, err error) {
	res = &v1.BlogListRes{
		List: []*model.BlogDetailOutput{},
	}
	ret, err := getBlogClient().GetList(ctx, &blogMicroV1.GetListReq{})
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

func (c *sBlog) BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "BlogDelete",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "BlogDelete",
				Message: "success",
			}.Log(ctx)
		}
	}()

	res = &v1.BlogDeleteRes{}
	currentUser, err := auth.GetStandardAuth().GetUserWithCheck(ctx)
	if err != nil {
		return
	}

	ret, err := getBlogClient().GetOne(ctx, &blogMicroV1.GetOneReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	if ret.Blog == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "博客不存在")
		return
	}
	if ret.Blog.Nickname != currentUser.Nickname {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "只能删除自己的博客")
		return
	}

	_, err = getBlogClient().Delete(ctx, &blogMicroV1.DeleteReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	return
}

func (c *sBlog) BlogBatDelete(ctx context.Context, req *v1.BlogBatDeleteReq) (res *v1.BlogBatDeleteRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "BlogBatDelete",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "BlogBatDelete",
				Message: "success",
			}.Log(ctx)
		}
	}()
	var ids []string
	res = &v1.BlogBatDeleteRes{}

	currentUser, err := auth.GetStandardAuth().GetUserWithCheck(ctx)
	if err != nil {
		return
	}

	for _, id := range strings.Split(req.Ids, ",") {
		var ret *blogMicroV1.GetOneRes
		ret, err = getBlogClient().GetOne(ctx, &blogMicroV1.GetOneReq{
			Id: id,
		})
		if err != nil {
			return
		}
		if ret.Blog == nil {
			err = gerror.NewCodef(gcode.CodeBusinessValidationFailed, "博客不存在:%d", id)
			return
		}
		if ret.Blog.Nickname != currentUser.Nickname {
			err = gerror.NewCodef(gcode.CodeBusinessValidationFailed, "只能删除自己的博客:%d", id)
			return
		}
		ids = append(ids, id)
	}
	ret, err := getBlogClient().BatDelete(ctx, &blogMicroV1.BatDeleteReq{
		Ids: ids,
	})
	if err != nil {
		return
	}
	res.BatNo = ret.BatNo
	return
}

func (c *sBlog) BlogGetBatDeleteStatus(ctx context.Context, req *v1.BlogGetBatDeleteStatusReq) (res *v1.BlogGetBatDeleteStatusRes, err error) {
	res = &v1.BlogGetBatDeleteStatusRes{}
	ret, err := getBlogClient().GetBatDeleteStatus(ctx, &blogMicroV1.GetBatDeleteStatusReq{
		BatNo: req.BatNo,
	})
	if err != nil {
		return
	}
	res.Status = ret.Status
	return
}
