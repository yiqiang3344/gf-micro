package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/imports"
	"strings"
)

var (
	GenHttpForGrpc = &gcmd.Command{
		Name:  "genHttpForGrpc",
		Usage: "./main genHttpForGrpc",
		Brief: "gen http for grpc cmd and controller",
		Arguments: append(CommonArguments, []gcmd.Argument{
			{
				Name:   "source",
				Short:  "s",
				Brief:  "grpc的控制器文件路径，可以是基于根目录的相对路径，或绝对路径",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "控制器文件输出目录，默认为`internal/controller/http/`",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "cmd",
				Short:  "c",
				Brief:  "cmd文件输出目录，默认为`internal/cmd/`",
				IsArg:  false,
				Orphan: false,
			},
		}...),
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			controllerPath := parser.GetOpt("source", "").String()
			folderPath := parser.GetOpt("output", "internal/controller/http/").String()
			cmdFolderPath := parser.GetOpt("cmd", "internal/cmd/").String()

			if strings.TrimSpace(controllerPath) == "" {
				return fmt.Errorf("output is empty")
			}

			//生成controller
			controllers, err := GenHttpForGrpcControllers(controllerPath, folderPath)
			if err != nil {
				return
			}
			//生成cmd
			err = GenHttpForGrpcCmd(controllers, cmdFolderPath)
			if err != nil {
				return
			}
			return
		},
	}
	controllerTemp = `
package {Package}

type {StructName} struct {
}
`
	controllerMethodTemp = `
func (c *{StructName}) {MethodName}(r *ghttp.Request) {
	cmd.HttpForGrpcFunc(r, &{ReqName}{}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		res, err = new({controllerName}).{MethodName}(ctx, req.(*{ReqName}))
		return
	})
}
`
	httpForGrpcCmdTemp = `
package cmd

var (
	httpForGrpc = &gcmd.Command{
		Name:      "httpForGrpc",
		Usage:     "./main httpForGrpc",
		Brief:     "start http for grpc server",
		Arguments: append(cmd.CommonArguments, []gcmd.Argument{}...),
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(g.Cfg().MustGet(ctx, cfg.APPNAME).String())
			s.SetAddr(cmd.GetCommonArguments(ctx, parser, cmd.Port).String())
			s.Use(
				logging.HttpLogFormatJsonMiddleware,
				logging.HttpAccessLogMiddleware,
				logging.HttpErrorLogMiddleware,
				response.HttpForGrpcResponseMiddleware,
			)
			{binds}
			s.Run()
			return
		},
	}
)
`
	httpForGrpcCmdBindTemp = `
s.BindObject("/{.struct}/{.method}", new({controllerName}))
`
)

type GrpcController struct {
	ControllerName string        //原始控制器完整包及名称
	Package        string        //包名
	Name           string        //控制器名称
	OriginName     string        //原始结构体名称
	Methods        []*GrpcMethod //方法列表
}

type GrpcMethod struct {
	Name    string
	ReqName string
}

func GoFmt(path string) {
	replaceFunc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			fmt.Printf(`error format "%s" go files: %v`, path, err)
			return content
		}
		str := string(res)
		//替换可能不合理的http包
		str = gstr.Replace(str, "\"github.com/gogf/gf/net/ghttp\"", "\"github.com/gogf/gf/v2/net/ghttp\"")
		str = gstr.Replace(str, "\"github.com/gogf/gf/frame/g\"", "\"github.com/gogf/gf/v2/frame/g\"")
		return str
	}

	var err error
	if gfile.IsFile(path) {
		// File format.
		if gfile.ExtName(path) != "go" {
			return
		}
		err = gfile.ReplaceFileFunc(replaceFunc, path)
	} else {
		// Folder format.
		err = gfile.ReplaceDirFunc(replaceFunc, path, "*.go", true)
	}
	if err != nil {
		fmt.Printf(`error format "%s" go files: %v`, path, err)
	}
}

func ParseGrpcControllerFile(filePath string) (controllers []*GrpcController, err error) {
	var (
		controllerPackage = "controller"
	)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return
	}
	packageName := file.Name.Name //包名
	//获取结构体map
	structMap := make(map[string]*GrpcController)
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			if _, ok := typeSpec.Type.(*ast.StructType); !ok {
				continue
			}
			structV := &GrpcController{
				Package:        controllerPackage,
				OriginName:     typeSpec.Name.Name,
				Name:           gstr.UcWords(packageName), //默认使用包名作为控制器名称
				ControllerName: fmt.Sprintf("%s.%s", packageName, typeSpec.Name.Name),
			}
			structMap[typeSpec.Name.Name] = structV
			controllers = append(controllers, structV)
		}
	}
	//获取每个结构体对应的方法信息
	for _, decl := range file.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		//只处理包括receiver的函数
		if funcDecl.Recv == nil || len(funcDecl.Recv.List) != 1 {
			continue
		}
		receiverName := funcDecl.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
		//只处理上面结构体定义的方法
		if _, ok := structMap[receiverName]; !ok {
			continue
		}
		if len(funcDecl.Type.Params.List) != 2 {
			continue
		}
		vv := funcDecl.Type.Params.List[1].Type.(*ast.StarExpr).X.(*ast.SelectorExpr)
		structMap[receiverName].Methods = append(structMap[receiverName].Methods, &GrpcMethod{
			Name:    funcDecl.Name.Name,
			ReqName: fmt.Sprintf("%s.%s", vv.X.(*ast.Ident).Name, vv.Sel.Name),
		})
	}

	//如果控制器数量大于1，则需要对控制器名称做特殊处理，使用原始结构体名称作为控制器名称
	if len(controllers) > 1 {
		for _, v := range controllers {
			v.Name = v.OriginName
		}
	}

	return
}

func GenHttpForGrpcControllers(controllerPath string, folderPath string) (controllers []*GrpcController, err error) {
	var (
		isDirty          bool
		generatedContent string
	)

	controllers, err = ParseGrpcControllerFile(controllerPath)
	if err != nil {
		return
	}

	if !gfile.Exists(folderPath) {
		if err = gfile.Mkdir(folderPath); err != nil {
			return
		}
	}

	for _, controllerV := range controllers {
		isDirty = false
		filePath := gfile.Join(folderPath, gstr.LcFirst(controllerV.Name)+".go")
		templateContent := gstr.ReplaceByMap(controllerTemp, g.MapStrStr{
			"{Package}":    controllerV.Package,
			"{StructName}": controllerV.Name,
		})
		if err = gfile.PutContents(filePath, templateContent); err != nil {
			return
		}
		isDirty = true

		for _, method := range controllerV.Methods {
			if generatedContent != "" {
				generatedContent += "\n"
			}
			generatedContent += gstr.ReplaceByMap(controllerMethodTemp, g.MapStrStr{
				"{StructName}":     controllerV.Name,
				"{MethodName}":     method.Name,
				"{ReqName}":        method.ReqName,
				"{controllerName}": controllerV.ControllerName,
			})
		}

		if generatedContent != "" {
			err = gfile.PutContentsAppend(filePath, generatedContent)
			if err != nil {
				return
			}
			isDirty = true
		}
		if isDirty {
			GoFmt(filePath)
		}
	}
	return
}

func GenHttpForGrpcCmd(controllers []*GrpcController, folderPath string) (err error) {
	var (
		generatedContent = ""
		filePath         = gfile.Join(folderPath, "httpForGrpc.go")
	)

	if !gfile.Exists(folderPath) {
		if err = gfile.Mkdir(folderPath); err != nil {
			return err
		}
	}

	strs := make([]string, 0)
	for _, controllerV := range controllers {
		strs = append(strs, gstr.ReplaceByMap(httpForGrpcCmdBindTemp, g.MapStrStr{
			"{controllerName}": controllerV.Package + "." + controllerV.Name,
		}))
	}
	generatedContent = gstr.ReplaceByMap(httpForGrpcCmdTemp, g.MapStrStr{
		"{binds}": strings.Join(strs, "\n"),
	})
	if err = gfile.PutContents(filePath, generatedContent); err != nil {
		return err
	}
	GoFmt(filePath)
	return
}

func HttpForGrpcFunc(r *ghttp.Request, req interface{}, f func(ctx context.Context, req interface{}) (res interface{}, err error)) {
	var err1 error
	reqJson, err1 := r.GetJson()
	if err1 != nil {
		r.SetError(err1)
		return
	}
	err1 = reqJson.Scan(req)
	if err1 != nil {
		r.SetError(err1)
		return
	}
	//参数校验
	if err := g.Validator().Data(req).Run(r.GetCtx()); err != nil {
		r.SetError(err)
		return
	}

	ret, err := f(r.GetCtx(), req)
	if err != nil {
		r.SetError(err)
		return
	}
	r.Response.WriteJson(ret)
}
