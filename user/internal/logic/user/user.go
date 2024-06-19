package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"yijunqiang/gf-micro/user/internal/logging"
	"yijunqiang/gf-micro/user/internal/model/entity"
	"yijunqiang/gf-micro/user/internal/utility/mcache"

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
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "Create",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Create",
				Message: "success",
			}.Log(ctx)
		}
	}()
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

func (s *sUser) Login(ctx context.Context, nickname string, password string) (token string, ret *pbentity.User, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "Login",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Login",
				Message: "success",
			}.Log(ctx)
		}
	}()
	var user *entity.User
	err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     mcache.GetDbCacheKey(dao.User.Table(), fmt.Sprintf("LoginByPassword:%s", gmd5.MustEncryptString(nickname+password))),
	}).Where(do.User{
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
	if err != nil {
		return
	}
	if user != nil {
		ret = &pbentity.User{}
		gconv.ConvertWithRefer(user, ret)
		ret.CreateAt = timestamppb.New(user.CreateAt.Time)
		ret.UpdateAt = timestamppb.New(user.UpdateAt.Time)
	}

	//缓存的token已存在则直接使用，不存在则生成新的
	token, err = GetCacheToken(ctx, gconv.String(user.Id))
	if err != nil {
		return
	}
	if !g.IsEmpty(token) {
		return
	}
	token, err = Token(ctx, gconv.String(user.Id))
	if err != nil {
		return
	}
	return
}

func (s *sUser) Logout(ctx context.Context, uid string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "Logout",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "Logout",
				Message: "success",
			}.Log(ctx)
		}
	}()
	err = Clear(ctx, uid)
	return
}

func (s *sUser) GetById(ctx context.Context, uid string) (ret *pbentity.User, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour,
		Name:     mcache.GetDbCacheKey(dao.User.Table(), fmt.Sprintf("GetById:%s", uid)),
	}).Where(do.User{
		Id: uid,
	}).Scan(&user)
	if err != nil {
		return
	}
	if user != nil {
		ret = &pbentity.User{}
		gconv.ConvertWithRefer(user, ret)
		ret.CreateAt = timestamppb.New(user.CreateAt.Time)
		ret.UpdateAt = timestamppb.New(user.UpdateAt.Time)
	}
	return
}

func (s *sUser) GetByToken(ctx context.Context, token string) (ret *pbentity.User, err error) {
	uid, err := Parse(ctx, token)
	if err != nil {
		return
	}
	ret, err = s.GetById(ctx, uid)
	if err != nil {
		return
	}
	return
}
