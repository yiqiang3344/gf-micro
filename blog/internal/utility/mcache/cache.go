package mcache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

func RedisCache() *gcache.Cache {
	return gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis()))
}
