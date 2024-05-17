package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"web/internal/model"
)

type UserCreateReq struct {
	g.Meta   `path:"/user/create" tags:"User" method:"post" summary:"用户/注册"`
	Nickname string `json:"nickname" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}
type UserCreateRes struct {
}

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"用户/登录"`
	Nickname string `json:"nickname" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}
type UserLoginRes struct {
	Token string `json:"token" dc:"token"`
}

type UserDetailReq struct {
	g.Meta `path:"/user/detail" tags:"User" method:"post" summary:"用户/详情"`
	Id     string `json:"id" v:"required#用户ID不能为空" dc:"ID"`
}
type UserDetailRes struct {
	*model.UserDetailOutput
}
