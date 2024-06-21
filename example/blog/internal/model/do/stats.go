// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Stats is the golang structure of table stats for DAO operations like Where/Data.
type Stats struct {
	g.Meta   `orm:"table:stats, do:true"`
	Id       interface{} // Blog ID
	Nickname interface{} // Nickname
	BlogCnt  interface{} // Blog Count
	CreateAt *gtime.Time // Created Time
	UpdateAt *gtime.Time // Updated Time
	DeleteAt *gtime.Time // Deleted Time
}
