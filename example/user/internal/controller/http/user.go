package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yiqiang3344/gf-micro/example/user/api/http/v1"
	"github.com/yiqiang3344/gf-micro/example/user/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) UserGetOne(ctx context.Context, req *v1.UserGetOneReq) (res *v1.UserGetOneRes, err error) {
	ret, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return
	}
	if ret != nil {
		res = &v1.UserGetOneRes{}
		gconv.ConvertWithRefer(ret, res)
	}
	return
}
