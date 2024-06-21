// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Blog is the golang structure for table blog.
type Blog struct {
	Id       uint        `json:"id"       orm:"id"        description:"Blog ID"`
	Title    string      `json:"title"    orm:"title"     description:"Title"`
	Content  string      `json:"content"  orm:"content"   description:"Content"`
	Nickname string      `json:"nickname" orm:"nickname"  description:"Nickname"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"Created Time"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"Updated Time"`
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" description:"Deleted Time"`
}
