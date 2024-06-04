package model

type User struct {
	Id       uint32 `json:"id" dc:"ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	Token    string `json:"token" dc:"token"`
}

type UserDetailOutput struct {
	Id       uint32 `json:"id" dc:"ID"`
	Nickname string `json:"nickname" dc:"昵称"`
}
