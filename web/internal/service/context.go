// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"web/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		InitRUser(r *ghttp.Request, user *model.User)
		InitUser(ctx context.Context, user *model.User) context.Context
		Get(ctx context.Context) *model.MContext
		GetUser(ctx context.Context) *model.User
		GetUserWithCheck(ctx context.Context) (ret *model.User, err error)
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
