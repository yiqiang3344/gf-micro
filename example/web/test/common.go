package test

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"github.com/yiqiang3344/gf-micro/logging"
	"golang.org/x/net/context"
	v1 "web/api/user/v1"
)

var (
	Port int
)

func init() {
	//接入配置中心
	ctx := gctx.New()
	if gcfg.Instance().MustGet(ctx, "apollo") != nil {
		adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
		if err != nil {
			panic(err)
		}
		gcfg.Instance().SetAdapter(adapter)
	}
	glog.SetDefaultHandler(logging.HandlerJson)
}

func GetClient(ctx context.Context, token ...string) *gclient.Client {
	if Port == 0 {
		Port = gconv.Int(gstr.Split(g.Cfg().MustGet(ctx, "server.address").String(), ":")[1])
	}

	prefix := fmt.Sprintf("http://127.0.0.1:%d", Port)
	client := g.Client()
	client.Use(logging.MiddlewareClientLog)
	client.SetPrefix(prefix)

	if len(token) > 0 {
		client.SetHeader("token", token[0])
	}

	return client
}

func Login(ctx context.Context, nickname, Password string, client *gclient.Client) (err error) {
	retTmp := client.PostContent(ctx, "/user/login", v1.UserLoginReq{
		Nickname: nickname,
		Password: Password,
	})
	j, err := gjson.DecodeToJson(retTmp)
	if err != nil {
		return
	}
	if j.Get("code").String() != "0" {
		err = errors.New("登录失败")
		return
	}
	client.SetHeader("token", j.Get("data.token").String())
	return
}
