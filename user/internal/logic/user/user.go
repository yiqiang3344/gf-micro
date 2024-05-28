package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"yijunqiang/gf-micro/user/internal/model/entity"

	"yijunqiang/gf-micro/user/api/pbentity"
	"yijunqiang/gf-micro/user/internal/dao"
	"yijunqiang/gf-micro/user/internal/model/do"
	"yijunqiang/gf-micro/user/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(&sUser{})
}

func (s *sUser) Create(ctx context.Context, nickname string, password string) (user *entity.User, err error) {
	user = &entity.User{
		Nickname: nickname,
		Password: password,
	}
	_, err = dao.User.Ctx(ctx).Data(user).Insert()
	if err != nil {
		return
	}
	return
}

func (s *sUser) Login(ctx context.Context, nickname string, password string) (token string, err error) {
	user := (*entity.User)(nil)
	err = dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
		Password: password,
	}).Scan(&user)
	if err != nil {
		return
	}
	if user == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "账号或密码错误")
		return
	}
	token = user.Nickname
	return
}

func (s *sUser) GetById(ctx context.Context, uid string) (*pbentity.User, error) {
	var user *pbentity.User
	err := dao.User.Ctx(ctx).Where(do.User{
		Id: uid,
	}).Scan(&user)
	return user, err
}
