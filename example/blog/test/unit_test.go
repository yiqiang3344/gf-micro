package test

import (
	"context"
	"encoding/json"
	"flag"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	logging2 "github.com/yiqiang3344/gf-micro/logging"
	"github.com/yiqiang3344/gf-micro/testWithExcel"
	"strings"
	"time"

	"fmt"
	"github.com/yiqiang3344/gf-micro/cmd"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	v1 "github.com/yiqiang3344/gf-micro/example/blog/api/blog/v1"
)

var (
	blogClient v1.BlogClient
)

func init() {
	ctx := gctx.GetInitCtx()

	//初始化grpc全局中间件
	cmd.GetGrpcMiddleware()(ctx, &gcmd.Parser{})

	// 客户端初始化
	blogClient = v1.NewBlogClient(grpcx.Client.MustNewGrpcClientConn("blog", grpcx.Client.ChainUnary(
		logging2.GrpcClientLoggerUnary,
	)))
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

	funcListMap.Set("create", testInfo{filePath: "./unitExcel/创建.xlsx"})
	funcListMap.Set("detail", testInfo{filePath: "./unitExcel/详情.xlsx"})
	funcListMap.Set("edit", testInfo{filePath: "./unitExcel/编辑.xlsx"})
	funcListMap.Set("list", testInfo{filePath: "./unitExcel/列表.xlsx"})
	funcListMap.Set("delete", testInfo{filePath: "./unitExcel/删除.xlsx"})
	funcListMap.Set("batDelete", testInfo{filePath: "./unitExcel/批量删除.xlsx"})
	funcListMap.Set("batDeleteStatus", testInfo{filePath: "./unitExcel/批量删除状态.xlsx"})

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
				res, err1 = blogClient.Create(ctx, req)
			case "detail":
				req := &v1.GetOneReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = blogClient.GetOne(ctx, req)
			case "edit":
				req := &v1.EditReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = blogClient.Edit(ctx, req)
			case "list":
				req := &v1.GetListReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = blogClient.GetList(ctx, req)
			case "delete":
				req := &v1.DeleteReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = blogClient.Delete(ctx, req)
			case "batDelete":
				req := &v1.BatDeleteReq{}
				if caseInfo.Body["ids"] != "" {
					req.Ids = strings.Split(caseInfo.Body["ids"], ",")
				}
				res, err1 = blogClient.BatDelete(ctx, req)
			case "batDeleteStatus":
				req := &v1.GetBatDeleteStatusReq{}
				gconv.ConvertWithRefer(caseInfo.Body, req)
				res, err1 = blogClient.GetBatDeleteStatus(ctx, req)
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
