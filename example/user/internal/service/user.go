// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/yiqiang3344/gf-micro/example/user/internal/model/entity"
)

type (
	IUser interface {
		Create(ctx context.Context, nickname string, password string) (user *entity.User, err error)
		Login(ctx context.Context, nickname string, password string) (token string, user *entity.User, err error)
		Logout(ctx context.Context, uid string) (err error)
		GetById(ctx context.Context, uid string) (user *entity.User, err error)
		GetByToken(ctx context.Context, token string) (ret *entity.User, err error)
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
