package context

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"web/internal/model"
	"web/internal/service"
)

const (
	ContextKey = "ContextKey"
)

type sContext struct{}

func New() *sContext {
	return &sContext{}
}

func init() {
	service.RegisterContext(New())
}

func (s *sContext) InitRUser(r *ghttp.Request, user *model.User) {
	customCtx := model.MContext{
		User: user,
	}
	r.SetCtxVar(ContextKey, &customCtx)
}

func (s *sContext) InitUser(ctx context.Context, user *model.User) context.Context {
	customCtx := model.MContext{
		User: user,
	}
	ctx = context.WithValue(ctx, ContextKey, &customCtx)
	return ctx
}

func (s *sContext) Get(ctx context.Context) *model.MContext {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.MContext); ok {
		return localCtx
	}
	return nil
}

func (s *sContext) GetUser(ctx context.Context) *model.User {
	save := s.Get(ctx)
	if save == nil {
		return nil
	}
	return save.User
}

func (s *sContext) GetUserWithCheck(ctx context.Context) (ret *model.User, err error) {
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
