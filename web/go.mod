module web

go 1.15

require (
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.7.1
	github.com/gogf/gf/v2 v2.7.1
	github.com/stretchr/testify v1.8.3
	github.com/xuri/excelize/v2 v2.7.1
	golang.org/x/net v0.24.0
	yijunqiang/gf-micro/blog v0.0.0
	yijunqiang/gf-micro/user v0.0.0
)

replace yijunqiang/gf-micro/user v0.0.0 => ../user

replace yijunqiang/gf-micro/blog v0.0.0 => ../blog
