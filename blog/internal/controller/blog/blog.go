package blog

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	v1 "yijunqiang/gf-micro/blog/api/blog/v1"
	"yijunqiang/gf-micro/blog/internal/service"
)

type Controller struct {
	v1.UnimplementedBlogServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterBlogServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res = &v1.CreateRes{}
	_, err = service.Blog().Create(ctx, req.Title, req.Content, req.Nickname)
	return
}

func (*Controller) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	res = &v1.EditRes{}
	err = service.Blog().Edit(ctx, req.Id, req.Title, req.Content, req.Nickname)
	return
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	res.Blog, err = service.Blog().GetById(ctx, req.Id)
	return
}

func (*Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	res.List, err = service.Blog().GetList(ctx)
	return
}

func (*Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	res = &v1.DeleteRes{}
	err = service.Blog().Delete(ctx, req.Id)
	return
}

func (*Controller) BatDelete(ctx context.Context, req *v1.BatDeleteReq) (res *v1.BatDeleteRes, err error) {
	res = &v1.BatDeleteRes{}
	res.BatNo, err = service.Blog().BatDelete(ctx, req.Ids)
	return
}

func (*Controller) GetBatDeleteStatus(ctx context.Context, req *v1.GetBatDeleteStatusReq) (res *v1.GetBatDeleteStatusRes, err error) {
	res = &v1.GetBatDeleteStatusRes{}
	res.Status, err = service.Blog().GetBatDeleteStatus(ctx, req.BatNo)
	return
}
