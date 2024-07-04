package test

import (
	"context"
	"encoding/json"
	"flag"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"strings"
	"time"

	"fmt"
	"github.com/yiqiang3344/gf-micro/cmd"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/yiqiang3344/gf-micro/example/user/api/user/v1"
)

var (
	userClient v1.UserClient
)

func init() {
	ctx := gctx.GetInitCtx()

	//初始化grpc全局中间件
	cmd.GetGrpcMiddleware()(ctx)

	// 客户端初始化
	userClient = v1.NewUserClient(grpcx.Client.MustNewGrpcClientConn("user"))
}

type testInfo struct {
	filePath string
	route    string
}

// TestUnit 单元测试，可通过-args传参测试某些接口，也可以按顺序全量测试
// 跑之前需要先清空数据库和缓存，不能和全流程测试一起跑，会有数据冲突
// 跑之前需要启动web,user,blog,blogConsume几个服务
func TestUnit(t *testing.T) {
	funcListMap := gmap.NewListMap()

	funcListMap.Set("create", testInfo{filePath: "./unitExcel/注册.xlsx"})
	funcListMap.Set("login", testInfo{filePath: "./unitExcel/登录.xlsx"})
	funcListMap.Set("detail", testInfo{filePath: "./unitExcel/详情.xlsx"})
	funcListMap.Set("logout", testInfo{filePath: "./unitExcel/登出.xlsx"})

	if !flag.Parsed() {
		flag.Parse()
	}
	//读取要运行的方法列表，参数逗号分割
	funcsStr := flag.Arg(0)
	var funcs []string
	if funcsStr != "" {
		funcs = strings.Split(flag.Arg(0), ",")
	}
	if len(funcs) > 0 {
		for _, v := range funcs {
			info := funcListMap.Get(v).(testInfo)
			info.route = v
			do(t, info)
		}
	} else {
		funcListMap.Iterator(func(k, v interface{}) bool {
			info := v.(testInfo)
			info.route = k.(string)
			do(t, info)
			return true
		})
	}
}

func do(t *testing.T, info testInfo) {
	o, err := testWithExcel.New(t, info.filePath,
		testWithExcel.WithFailfast(false),
		testWithExcel.WithCaseHandleFunc(func(ctx context.Context, t *testing.T, caseInfo testWithExcel.CaseInfo) (ret interface{}, err error) {
			if caseInfo.Extend["delayTime"] != "" {
				time.Sleep(gconv.Duration(caseInfo.Extend["delayTime"]) * time.Second)
			}
			var (
				res  any
				err1 error
			)
			switch info.route {
			case "create":
				req := &v1.CreateReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = userClient.Create(ctx, req)
			case "login":
				req := &v1.LoginReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = userClient.Login(ctx, req)
			case "detail":
				req := &v1.GetOneReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = userClient.GetOne(ctx, req)
			case "logout":
				req := &v1.LogoutReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = userClient.Logout(ctx, req)
			}
			ret1, _ := json.Marshal(map[string]interface{}{
				"res": res,
				"err": err1,
			})
			ret = string(ret1)
			return
		}),
	)
	if err != nil {
		panic(fmt.Errorf("接口%s运行异常:%v", info.route, err))
	}
	o.Run(context.Background())
}
