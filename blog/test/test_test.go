package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcfg"
	gcfg_apollo "github.com/yiqiang3344/gcfg-apollo"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	v1 "yijunqiang/gf-micro/blog/api/blog/v1"
)

var (
	blogClient v1.BlogClient
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
	blogClient = v1.NewBlogClient(grpcx.Client.MustNewGrpcClientConn("blog"))
}

func _TestCreate(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
			err error
		)
		for i := 1; i <= 10; i++ {
			_, err = blogClient.Create(ctx, &v1.CreateReq{
				Title:    fmt.Sprintf(`blog-%d`, i),
				Content:  fmt.Sprintf(`content-%d`, i),
				Nickname: fmt.Sprintf(`nickname-%d`, i),
			})
			if err != nil {
				break
			}
		}
		gtest.Assert(err, "")
	})
}

func _TestGetOne(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
			res *v1.GetOneRes
			err error
		)
		res, err = blogClient.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		gtest.Assert(err, "")

		res, err = blogClient.GetOne(ctx, &v1.GetOneReq{
			Id: 100,
		})
		gtest.Assert(err, "")
		gtest.Assert(res.GetBlog(), "")
	})
}

func _TestEdit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
			err error
		)
		_, err = blogClient.Edit(ctx, &v1.EditReq{
			Id:       1,
			Title:    "title-1-edit",
			Content:  "content-1-edit",
			Nickname: "nickname-1",
		})
		gtest.Assert(err, "")
		ret1, err := blogClient.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		gtest.Assert(err, "")
		gtest.Assert(ret1.GetBlog().Title, "title-1-edit")
		gtest.Assert(ret1.GetBlog().Content, "content-1-edit")
	})
}

func _TestGetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		ret, err := blogClient.GetList(ctx, &v1.GetListReq{})
		gtest.Assert(err, "")
		gtest.AssertGT(len(ret.List), 0)
	})
}

func _TestDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		_, err := blogClient.Delete(ctx, &v1.DeleteReq{
			Id: 2,
		})
		gtest.Assert(err, "")

		ret1, err := blogClient.GetOne(ctx, &v1.GetOneReq{
			Id: 2,
		})
		gtest.Assert(err, "")
		gtest.AssertGT(ret1.GetBlog(), "")
	})
}

func TestBatDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		ret, err := blogClient.BatDelete(ctx, &v1.BatDeleteReq{
			Ids: []uint64{3, 4},
		})
		gtest.Assert(err, "")
		gtest.AssertNE(ret.BatNo, "")
	})
}

func _TestGetBatDeleteStatus(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		ret, err := blogClient.GetBatDeleteStatus(ctx, &v1.GetBatDeleteStatusReq{
			BatNo: "bat1",
		})
		gtest.Assert(err, "")
		gtest.Assert(ret.Status, "success")
	})
}

func _TestValidation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx = gctx.GetInitCtx()
		)
		_, err := blogClient.Create(ctx, &v1.CreateReq{})
		gtest.Assert(err, "The Title field is required")
	})
}
