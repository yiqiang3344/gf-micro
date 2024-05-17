package test

import (
	"fmt"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/test/gtest"

	v1 "yijunqiang/gf-micro/blog/api/blog/v1"
)

func Test_Create(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		for i := 1; i <= 10; i++ {
			_, err := blog.Create(ctx, &v1.CreateReq{
				Title:    fmt.Sprintf(`blog-%d`, i),
				Content:  fmt.Sprintf(`content-%d`, i),
				Nickname: fmt.Sprintf(`nickname-%d`, i),
			})
			if err != nil {
				g.Log().Fatalf(ctx, `create blog failed: %+v`, err)
			}
		}
		fmt.Println("test:Create success")
	})
}

func Test_GetOne(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
			res  *v1.GetOneRes
			err  error
		)
		res, err = blog.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get blog failed: %+v`, err)
		}
		fmt.Printf("test:GetOne success，result: %+v \n", res.Blog)

		res, err = blog.GetOne(ctx, &v1.GetOneReq{
			Id: 100,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get blog failed: %+v`, err)
		}
		if res.GetBlog() == nil {
			fmt.Println("test:GetOne valid success")
		}
	})
}

func Test_Edit(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		_, err := blog.Edit(ctx, &v1.EditReq{
			Title:    "title-1-edit",
			Content:  "content-1-edit",
			Nickname: "nickname-1",
		})
		if err != nil {
			g.Log().Fatalf(ctx, `edit blog failed: %+v`, err)
		}

		ret1, err := blog.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get blog failed: %+v`, err)
		}
		if ret1.GetBlog().Title != "title-1-edit" {
			g.Log().Fatalf(ctx, `edit blog title failed: %+v`, ret1.GetBlog())
		}

		if ret1.GetBlog().Content != "content-1-edit" {
			g.Log().Fatalf(ctx, `edit blog content failed: %+v`, ret1.GetBlog())
		}

		fmt.Printf("test:Edit success: %+v \n", ret1.GetBlog())
	})
}

func Test_GetList(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		ret, err := blog.GetList(ctx, &v1.GetListReq{})
		if err != nil {
			g.Log().Fatalf(ctx, `getList blog failed: %+v`, err)
		}

		fmt.Printf("test:getList success: %+v \n", ret)
	})
}

func Test_Delete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		_, err := blog.Delete(ctx, &v1.DeleteReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `delete blog failed: %+v`, err)
		}

		ret1, err := blog.GetOne(ctx, &v1.GetOneReq{
			Id: 1,
		})
		if err != nil {
			g.Log().Fatalf(ctx, `get blog failed: %+v`, err)
		}
		if ret1.GetBlog() != nil {
			g.Log().Fatalf(ctx, `delete blog failed: %+v`, ret1.GetBlog())
		}

		fmt.Println("test:delete success")
	})
}

func Test_BatDelete(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		_, err := blog.BatDelete(ctx, &v1.BatDeleteReq{
			Ids: []uint64{2, 3},
		})
		if err != nil {
			g.Log().Fatalf(ctx, `delete blog failed: %+v`, err)
		}

		fmt.Println("test:batDelete success")
	})
}

func Test_GetBatDeleteStatus(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		ret, err := blog.GetBatDeleteStatus(ctx, &v1.GetBatDeleteStatusReq{
			BatNo: "bat12345",
		})
		if err != nil {
			g.Log().Fatalf(ctx, `delete blog failed: %+v`, err)
		}

		fmt.Printf("test:getBatDeleteStatus success: %+v \n", ret)
	})
}

func Test_Validation(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			ctx  = gctx.GetInitCtx()
			conn = grpcx.Client.MustNewGrpcClientConn("blog")
			blog = v1.NewBlogClient(conn)
		)
		_, err := blog.Create(ctx, &v1.CreateReq{})
		if err == nil {
			g.Log().Fatal(ctx, "test: Create valid failed")
		}
		fmt.Printf("test:Create valid success，err: %+v \n", err.Error())
	})
}
