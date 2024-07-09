package user

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/yiqiang3344/gf-micro/example/user/api/user/v1"
	"github.com/yiqiang3344/gf-micro/example/user/internal/model/entity"
	"github.com/yiqiang3344/gf-micro/example/user/internal/service"
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
	return
}

func (*Controller) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	ret := &entity.User{}
	res.Token, ret, err = service.User().Login(ctx, req.Nickname, req.Password)
	if err != nil {
		return
	}
	if ret != nil {
		gconv.ConvertWithRefer(ret, &res.User)
	}
	return
}

func (*Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	ret, err := service.User().GetById(ctx, req.GetId())
	if err != nil {
		return
	}
	if ret != nil {
		gconv.ConvertWithRefer(ret, &res.User)
	}
	return
}

func (*Controller) GetByToken(ctx context.Context, req *v1.GetByTokenReq) (res *v1.GetByTokenRes, err error) {
	res = &v1.GetByTokenRes{}
	ret, err := service.User().GetByToken(ctx, req.GetToken())
	if err != nil {
		return
	}
	if ret != nil {
		gconv.ConvertWithRefer(ret, &res.User)
	}
	return
}

func (*Controller) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res = &v1.LogoutRes{}
	err = service.User().Logout(ctx, req.GetId())
	return
}
