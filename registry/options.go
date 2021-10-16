package registry

import "time"

type Options struct {
	Addrs        []string
	Timeout      time.Duration
	RegistryPath string
	HeartBeat    int64
}

type Operation func(opt *Options)

//func NewOptions(opts ...Operation) {
//	options := &Options{}
//	for _, opt := range opts {
//		opt(options)
//	}
//}

func WithAddrs(addrs []string) Operation {
	return func(opt *Options) {
		opt.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Operation {
	return func(opt *Options) {
		opt.Timeout = timeout
	}
}

func WithRegistryPath(registryPath string) Operation {
	return func(opt *Options) {
		opt.RegistryPath = registryPath
	}
}

func WithHeartBeat(heartBeat int64) Operation {
	return func(opt *Options) {
		opt.HeartBeat = heartBeat
	}
}
