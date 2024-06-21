package auth

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	ContextKey = "StandardAuth"
)

type User struct {
	Id       uint32 `json:"id" dc:"ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	Token    string `json:"token" dc:"token"`
}

type ContextValue struct {
	User *User // 登录用户信息
}

type StandardAuthContext interface {
	InitRUser(r *ghttp.Request, user *User)
	InitUser(ctx context.Context, user *User) context.Context
	Get(ctx context.Context) *ContextValue
	GetUser(ctx context.Context) *User
	GetUserWithCheck(ctx context.Context) (ret *User, err error)
}

func GetStandardAuth() StandardAuthContext {
	return &defaultStandardAuth{}
}

type defaultStandardAuth struct{}

func (s *defaultStandardAuth) InitRUser(r *ghttp.Request, user *User) {
	customCtxVal := ContextValue{
		User: user,
	}
	r.SetCtxVar(ContextKey, &customCtxVal)
}

func (s *defaultStandardAuth) InitUser(ctx context.Context, user *User) context.Context {
	customCtx := ContextValue{
		User: user,
	}
	ctx = context.WithValue(ctx, ContextKey, &customCtx)
	return ctx
}

func (s *defaultStandardAuth) Get(ctx context.Context) *ContextValue {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*ContextValue); ok {
		return localCtx
	}
	return nil
}

func (s *defaultStandardAuth) GetUser(ctx context.Context) *User {
	save := s.Get(ctx)
	if save == nil {
		return nil
	}
	return save.User
}

func (s *defaultStandardAuth) GetUserWithCheck(ctx context.Context) (ret *User, err error) {
	save := s.Get(ctx)
	if save == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "请先登录")
		return
	}
	ret = save.User
	if ret == nil {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "请先登录")
		return
	}
	return
}
