package hello

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"svc1/api/hello/v1"
	v12 "yiqiang3344/web/api/user/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	var (
		conn   = grpcx.Client.MustNewGrpcClientConn("demo")
		client = v12.NewUserClient(conn)
	)
	_res, err := client.Create(ctx, &v12.CreateReq{
		Passport: "test",
		Password: "1234",
		Nickname: "test",
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", _res.String())
	return
}
