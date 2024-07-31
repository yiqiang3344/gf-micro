package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/yiqiang3344/gf-micro/cfg"
	"strings"
)

// RedisCache 获取基于redis的缓存实例
func RedisCache() *gcache.Cache {
	return gcache.NewWithAdapter(gcache.NewAdapterRedis(g.Redis()))
}

// GetKeyWithApp 获取带应用标识前缀的缓存key，如 app:str1:str2，没有配置应用标识则为 unkownApp:str1:str2
func GetKeyWithApp(ctx context.Context, keys ...string) string {
	return fmt.Sprintf("%s:%s",
		g.Cfg().MustGet(ctx, cfg.APPNAME, "unkownApp"),
		strings.Join(keys, ":"),
	)
}
