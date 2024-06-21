// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Stats is the golang structure for table stats.
type Stats struct {
	Id       uint        `json:"id"       orm:"id"        description:"Blog ID"`
	Nickname string      `json:"nickname" orm:"nickname"  description:"Nickname"`
	BlogCnt  uint        `json:"blogCnt"  orm:"blog_cnt"  description:"Blog Count"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"Created Time"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"Updated Time"`
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" description:"Deleted Time"`
}
