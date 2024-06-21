package model

type UserDetailOutput struct {
	Id       uint32 `json:"id" dc:"ID"`
	Nickname string `json:"nickname" dc:"昵称"`
}
