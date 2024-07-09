package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserGetOneReq struct {
	g.Meta `path:"/user/detail" tags:"User" method:"post,get" summary:"用户/详情"`
	Id     string `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}
type UserGetOneRes struct {
	Id       uint32 `json:"id" dc:"ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	CreateAt string `json:"create_at" dc:"创建时间"`
}
