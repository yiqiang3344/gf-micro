package client

type opts struct {
	log           bool
	httpDiscovery bool
}

type OptsFunc func(o *opts)

// WithLog 是否打印client日志
func WithLog(log bool) OptsFunc {
	return func(o *opts) {
		o.log = log
	}
}

// WithHttpDiscovery 是否开启http服务发现，只有http客户端有用
func WithHttpDiscovery(httpDiscovery bool) OptsFunc {
	return func(o *opts) {
		o.httpDiscovery = httpDiscovery
	}
}
