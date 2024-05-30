package test

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	v1 "yijunqiang/gf-micro/user/api/user/v1"
)

func Test_Create(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("user")
			user = v1.NewUserClient(conn)
		)
		for i := 1; i <= 10; i++ {
			_, err := user.Create(ctx, &v1.CreateReq{
				Password: "123456",
				Nickname: fmt.Sprintf(`nickname-%d`, i),
			})
			if err != nil {
				g.Log().Fatalf(ctx, `create user failed: %+v`, err)
			}
		}
		fmt.Println("test:Create success")
	})
}

func Test_GetOne(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("user")
			user = v1.NewUserClient(conn)
			res  *v1.GetOneRes
			err  error
		)
		res, err = user.GetOne(ctx, &v1.GetOneReq{
			Id: "1",
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get user failed: %+v`, err)
		}
		fmt.Printf("test:GetOne success，result: %+v \n", res.User)

		res, err = user.GetOne(ctx, &v1.GetOneReq{
			Id: "100",
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get user failed: %+v`, err)
		}
		if res.GetUser() == nil {
			fmt.Println("test:GetOne valid success")
		}
	})
}

func Test_Login(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("user")
			user = v1.NewUserClient(conn)
		)
		ret, err := user.Login(ctx, &v1.LoginReq{
			Nickname: "nickname-1",
			Password: "123456",
		})
		if err != nil {
			g.Log().Fatalf(ctx, `login user failed: %+v`, err)
		}
		if ret.GetToken() != "nickname-1" {
			g.Log().Fatalf(ctx, `login user token valid: %+v`, ret.GetToken())
		}
		fmt.Printf("test:Login success，result: %+v \n", ret)
	})
}

func Test_Validation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("user")
			user = v1.NewUserClient(conn)
		)
		_, err := user.Login(ctx, &v1.LoginReq{})
		if err != nil {
			fmt.Printf("test:Login valid success，err: %+v \n", err.Error())
		}
	})
}
