// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"yijunqiang/gf-micro/user/api/pbentity"
	"yijunqiang/gf-micro/user/internal/model/entity"
)

type (
	IUser interface {
		Create(ctx context.Context, nickname string, password string) (user *entity.User, err error)
		Login(ctx context.Context, nickname string, password string) (token string, err error)
		GetById(ctx context.Context, uid string) (*pbentity.User, error)
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
