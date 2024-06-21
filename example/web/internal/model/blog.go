package model

type BlogDetailOutput struct {
	Id       uint32 `json:"id" dc:"ID"`
	Title    string `json:"title" dc:"标题"`
	Content  string `json:"content" dc:"内容"`
	Nickname string `json:"nickname" dc:"作者"`
}
