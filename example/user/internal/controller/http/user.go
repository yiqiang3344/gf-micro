package controller

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/yiqiang3344/gf-micro/cmd"
	v1 "github.com/yiqiang3344/gf-micro/example/user/api/user/v1"
	"github.com/yiqiang3344/gf-micro/example/user/internal/controller/user"
)

type User struct {
}

func (c *User) Create(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &v1.CreateReq{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new(user.Controller).Create(ctx, req.(*v1.CreateReq))
		return
	})
}

func (c *User) Login(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &v1.LoginReq{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new(user.Controller).Login(ctx, req.(*v1.LoginReq))
		return
	})
}

func (c *User) GetOne(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &v1.GetOneReq{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new(user.Controller).GetOne(ctx, req.(*v1.GetOneReq))
		return
	})
}

func (c *User) GetByToken(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &v1.GetByTokenReq{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new(user.Controller).GetByToken(ctx, req.(*v1.GetByTokenReq))
		return
	})
}

func (c *User) Logout(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &v1.LogoutReq{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new(user.Controller).Logout(ctx, req.(*v1.LogoutReq))
		return
	})
}
