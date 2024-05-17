package controller

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "web/api/user/v1"
	"web/internal/model"
	userMicroV1 "yijunqiang/gf-micro/user/api/user/v1"
)

var (
	conn   = grpcx.Client.MustNewGrpcClientConn("user")
	client = userMicroV1.NewUserClient(conn)
)

func (c *cUser) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	res = &v1.UserCreateRes{}
	_, err = client.Create(ctx, &userMicroV1.CreateReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (c *cUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res = &v1.UserLoginRes{}
	ret, err := client.Login(ctx, &userMicroV1.LoginReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	res.Token = ret.Token
	return
}

func (c *cUser) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	res = &v1.UserDetailRes{}
	ret, err := client.GetOne(ctx, &userMicroV1.GetOneReq{
		Id: req.Id,
	})
	if err != nil {
		return
	}
	if ret.User == nil {
		err = gerror.NewCode(gcode.New(-1, "用户不存在", nil))
		return
	}
	res.UserDetailOutput = &model.UserDetailOutput{
		Id:       ret.User.Id,
		Nickname: ret.User.Nickname,
	}
	return
}
