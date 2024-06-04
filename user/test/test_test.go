package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/v2/os/gcfg"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"testing"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	v1 "yijunqiang/gf-micro/user/api/user/v1"
)

var (
	userClient v1.UserClient
)

func init() {
	//接入配置中心
	ctx := gctx.GetInitCtx()
	if gcfg.Instance().MustGet(ctx, "apollo") != nil {
		adapter, err := gcfg_apollo.CreateAdapterApollo(ctx)
		if err != nil {
			panic(err)
		}
		gcfg.Instance().SetAdapter(adapter)
	}
	// grpc服务注册发现
	grpcx.Resolver.Register(etcd.New(gcfg.Instance().MustGet(ctx, "registry.etcd").String()))

	// 客户端初始化
	userClient = v1.NewUserClient(grpcx.Client.MustNewGrpcClientConn("user"))
}

func TestCreate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
			err error
		)
		for i := 1; i <= 10; i++ {
			_, err = userClient.Create(ctx, &v1.CreateReq{
				Password: "123456",
				Nickname: fmt.Sprintf(`nickname-%d`, i),
			})
			if err != nil {
				break
			}
		}
		gtest.Assert(err, "")
	})
}

func TestGetOne(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
			res *v1.GetOneRes
			err error
		)
		res, err = userClient.GetOne(ctx, &v1.GetOneReq{
			Id: "1",
		})
		gtest.Assert(err, "")
		gtest.Assert(res.User.Id, 1)

		res, err = userClient.GetOne(ctx, &v1.GetOneReq{
			Id: "100",
		})
		gtest.Assert(err, "")
		gtest.Assert(res.User, "")
	})
}

func TestLogin(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		ret, err := userClient.Login(ctx, &v1.LoginReq{
			Nickname: "nickname-1",
			Password: "123456",
		})
		gtest.Assert(err, "")
		gtest.AssertNE(ret.GetToken(), "")
		gtest.AssertNE(ret.GetUser(), "")
	})
}

func TestLogout(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		_, err := userClient.Logout(ctx, &v1.LogoutReq{
			Id: "1",
		})
		gtest.Assert(err, "")
	})
}

func TestValidation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		_, err := userClient.Login(ctx, &v1.LoginReq{})
		gtest.Assert(err, "The Nickname field is required")
	})
}
