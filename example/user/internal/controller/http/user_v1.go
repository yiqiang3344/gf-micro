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

func (c *cUser) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	ret, err := service.User().GetById(ctx, req.Id)
	if err != nil {
		return
	}
	if ret != nil {
		res = &v1.UserDetailRes{}
		gconv.ConvertWithRefer(ret, res)
		res.CreateAt = ret.CreateAt.AsTime().Format("2006-01-02 15:04:05")
	}
	return
}
