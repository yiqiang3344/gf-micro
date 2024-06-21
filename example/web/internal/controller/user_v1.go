package controller

import (
	"context"
	v1 "web/api/user/v1"
	"web/internal/service"
)

func (c *cUser) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	res, err = service.User().UserCreate(ctx, req)
	return
}

func (c *cUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res, err = service.User().UserLogin(ctx, req)
	return
}

func (c *cUser) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	res, err = service.User().UserLogout(ctx, req)
	return
}

func (c *cUser) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	res, err = service.User().UserDetail(ctx, req)
	return
}
