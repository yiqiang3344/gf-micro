// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "web/api/user/v1"
	"web/internal/model"
)

type (
	IUser interface {
		UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error)
		UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
		UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
		UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error)
		LoginByToken(ctx context.Context, token string) (user *model.User, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
