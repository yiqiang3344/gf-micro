// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Blog is the golang structure for table blog.
type Blog struct {
	Id       uint        `json:"id"       description:"Blog ID"`
	Title    string      `json:"title"    description:"Title"`
	Content  string      `json:"content"  description:"Content"`
	Nickname string      `json:"nickname" description:"Nickname"`
	CreateAt *gtime.Time `json:"createAt" description:"Created Time"`
	UpdateAt *gtime.Time `json:"updateAt" description:"Updated Time"`
	DeleteAt *gtime.Time `json:"deleteAt" description:"Deleted Time"`
}
