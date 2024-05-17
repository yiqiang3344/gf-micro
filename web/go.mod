module web

go 1.15

require (
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.7.1
	github.com/gogf/gf/v2 v2.7.1
	github.com/xuri/excelize/v2 v2.7.1 // indirect
	golang.org/x/net v0.24.0 // indirect
	yijunqiang/gf-micro/user v0.0.0
)

replace yijunqiang/gf-micro/user v0.0.0 => ../user

replace yijunqiang/gf-micro/blog v0.0.0 => ../blog
