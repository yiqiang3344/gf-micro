// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogDao is the data access object for table blog.
type BlogDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns BlogColumns // columns contains all the column names of Table for convenient usage.
}

// BlogColumns defines and stores column names for table blog.
type BlogColumns struct {
	Id       string // Blog ID
	Title    string // Title
	Content  string // Content
	Nickname string // Nickname
	CreateAt string // Created Time
	UpdateAt string // Updated Time
	DeleteAt string // Deleted Time
}

// blogColumns holds the columns for table blog.
var blogColumns = BlogColumns{
	Id:       "id",
	Title:    "title",
	Content:  "content",
	Nickname: "nickname",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewBlogDao creates and returns a new DAO object for table data access.
func NewBlogDao() *BlogDao {
	return &BlogDao{
		group:   "default",
		table:   "blog",
		columns: blogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogDao) Columns() BlogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
