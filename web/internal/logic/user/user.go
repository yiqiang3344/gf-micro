package user

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "web/api/user/v1"
	"web/internal/logging"
	"web/internal/model"
	"web/internal/service"
	userMicroV1 "yijunqiang/gf-micro/user/api/user/v1"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

var userClient userMicroV1.UserClient

func getUserClient() userMicroV1.UserClient {
	if userClient == nil {
		userClient = userMicroV1.NewUserClient(grpcx.Client.MustNewGrpcClientConn("user"))
	}
	return userClient
}

func (s *sUser) UserCreate(ctx context.Context, req *v1.UserCreateReq) (res *v1.UserCreateRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "UserCreate",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "UserCreate",
				Message: "success",
			}.Log(ctx)
		}
	}()
	res = &v1.UserCreateRes{}
	_, err = getUserClient().Create(ctx, &userMicroV1.CreateReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	return
}

func (s *sUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "UserLogin",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "UserLogin",
				Message: "success",
			}.Log(ctx)
		}
	}()
	res = &v1.UserLoginRes{}
	ret, err := getUserClient().Login(ctx, &userMicroV1.LoginReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	res.Token = ret.Token
	return
}

func (s *sUser) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			logging.BizLog{
				Tag:     "UserLogout",
				Message: "failed",
			}.Log(ctx)
		} else {
			logging.BizLog{
				Tag:     "UserLogout",
				Message: "success",
			}.Log(ctx)
		}
	}()
	res = &v1.UserLogoutRes{}
	currentUser, err := service.Context().GetUserWithCheck(ctx)
	if err != nil {
		return
	}

	_, err = getUserClient().Logout(ctx, &userMicroV1.LogoutReq{
		Id: gconv.String(currentUser.Id),
	})
	if err != nil {
		return
	}
	return
}

func (s *sUser) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	res = &v1.UserDetailRes{}
	currentUser, err := service.Context().GetUserWithCheck(ctx)
	if err != nil {
		return
	}

	ret, err := getUserClient().GetOne(ctx, &userMicroV1.GetOneReq{
		Id: gconv.String(currentUser.Id),
	})
	if err != nil {
		return
	}
	if ret.User == nil {
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户不存在")
		return
	}
	res.UserDetailOutput = &model.UserDetailOutput{
		Id:       ret.User.Id,
		Nickname: ret.User.Nickname,
	}
	return
}

func (s *sUser) LoginByToken(ctx context.Context, token string) (user *model.User, err error) {
	if token == "" {
		err = gerror.NewCode(gcode.CodeNotAuthorized, "登陆失败")
		return
	}
	ret, err := getUserClient().GetByToken(ctx, &userMicroV1.GetByTokenReq{
		Token: token,
	})
	if err != nil {
		return
	}
	user = &model.User{
		Id:       ret.User.Id,
		Nickname: ret.User.Nickname,
	}
	return
}
