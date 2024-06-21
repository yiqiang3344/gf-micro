package cache

import (
	"fmt"
)

// GetDbCacheFullKey 获取数据库缓存完整的key
func GetDbCacheFullKey(table string, key string) string {
	//gf/database/gdb/gdb.go:386中有定义缓存前缀，但没开放出来
	return fmt.Sprintf("SelectCache:%s@%s", table, key)
}

// GetDbCacheKey 获取数据库缓存标准的key
func GetDbCacheKey(table string, key string) string {
	return fmt.Sprintf("%s@%s", table, key)
}
