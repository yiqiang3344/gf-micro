// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StatsDao is the data access object for table stats.
type StatsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns StatsColumns // columns contains all the column names of Table for convenient usage.
}

// StatsColumns defines and stores column names for table stats.
type StatsColumns struct {
	Id       string // Blog ID
	Nickname string // Nickname
	BlogCnt  string // Blog Count
	CreateAt string // Created Time
	UpdateAt string // Updated Time
	DeleteAt string // Deleted Time
}

// statsColumns holds the columns for table stats.
var statsColumns = StatsColumns{
	Id:       "id",
	Nickname: "nickname",
	BlogCnt:  "blog_cnt",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewStatsDao creates and returns a new DAO object for table data access.
func NewStatsDao() *StatsDao {
	return &StatsDao{
		group:   "default",
		table:   "stats",
		columns: statsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StatsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StatsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StatsDao) Columns() StatsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StatsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StatsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StatsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
