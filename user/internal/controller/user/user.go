package user

import (
	"context"
	v1 "yijunqiang/gf-micro/user/api/user/v1"
	"yijunqiang/gf-micro/user/internal/logging"
	"yijunqiang/gf-micro/user/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	res = &v1.CreateRes{}
	_, err = service.User().Create(ctx, req.Nickname, req.Password)
	if err != nil {
		logging.BizLog{
			Tag:     "Create",
			Message: "failed",
		}.Log(ctx)
		return
	}
	logging.BizLog{
		Tag:     "Create",
		Message: "success",
	}.Log(ctx)
	return
}

func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	res.Token, err = service.User().Login(ctx, req.Nickname, req.Password)
	if err != nil {
		logging.BizLog{
			Tag:     "Login",
			Message: "failed",
		}.Log(ctx)
		return
	}
	logging.BizLog{
		Tag:     "Login",
		Message: "success",
	}.Log(ctx)
	return
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	res.User, err = service.User().GetById(ctx, req.GetId())
	return
}
