package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"web/internal/model"
)

type BlogCreateReq struct {
	g.Meta  `path:"/blog/create" tags:"Blog" method:"post" summary:"博客/创建"`
	Title   string `json:"title" v:"required#标题不能为空" dc:"标题"`
	Content string `json:"content" v:"required#内容不能为空" dc:"内容"`
}
type BlogCreateRes struct {
}

type BlogEditReq struct {
	g.Meta  `path:"/blog/edit" tags:"Blog" method:"post" summary:"博客/编辑"`
	Id      string `json:"id" v:"required#ID不能为空" dc:"ID"`
	Title   string `json:"title" v:"required#标题不能为空" dc:"标题"`
	Content string `json:"content" v:"required#内容不能为空" dc:"内容"`
}
type BlogEditRes struct {
}

type BlogDetailReq struct {
	g.Meta `path:"/blog/detail" tags:"Blog" method:"post" summary:"博客/详情"`
	Id     string `json:"id" v:"required#博客ID不能为空" dc:"ID"`
}
type BlogDetailRes struct {
	*model.BlogDetailOutput
}

type BlogListReq struct {
	g.Meta `path:"/blog/list" tags:"Blog" method:"post" summary:"博客/列表"`
}
type BlogListRes struct {
	List []*model.BlogDetailOutput `json:"list" dc:"博客列表"`
}

type BlogDeleteReq struct {
	g.Meta `path:"/blog/delete" tags:"Blog" method:"post" summary:"博客/删除"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"ID"`
}
type BlogDeleteRes struct {
}

type BlogBatDeleteReq struct {
	g.Meta `path:"/blog/bat-delete" tags:"Blog" method:"post" summary:"博客/批量删除"`
	Ids    string `json:"ids" v:"required#ID列表不能为空" dc:"ID列表,逗号分割"`
}
type BlogBatDeleteRes struct {
	BatNo string `json:"batNo" dc:"批次号"`
}

type BlogGetBatDeleteStatusReq struct {
	g.Meta `path:"/blog/get-bat-delete-status" tags:"Blog" method:"post" summary:"博客/获取批量删除状态"`
	BatNo  string `json:"batNo" v:"required#批次号不能为空" dc:"批次号"`
}
type BlogGetBatDeleteStatusRes struct {
	Status string `json:"status" dc:"状态"`
}
