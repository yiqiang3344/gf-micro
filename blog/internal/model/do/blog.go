// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Blog is the golang structure of table blog for DAO operations like Where/Data.
type Blog struct {
	g.Meta   `orm:"table:blog, do:true"`
	Id       interface{} // Blog ID
	Title    interface{} // Title
	Content  interface{} // Content
	Nickname interface{} // Nickname
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
	DeleteAt *gtime.Time // Deleted Time
}
