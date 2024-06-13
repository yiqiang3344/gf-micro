package mcache

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func RedisCache() *gcache.Cache {
	return gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis()))
}

func GetDbCacheFullKey(table string, key string) string {
	//gf/database/gdb/gdb.go:386中有定义缓存前缀，但没开放出来
	return fmt.Sprintf("SelectCache:%s@%s", table, key)
}

func GetDbCacheKey(table string, key string) string {
	return fmt.Sprintf("%s@%s", table, key)
}
