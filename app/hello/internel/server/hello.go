package server

import (
	"examples/app/hello/internel/conf"
	"examples/app/hello/internel/pkg/options"
	"examples/app/hello/internel/service"
	"examples/kitex_gen/hello/v1/hello"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func NewRPCServer(s *service.HelloService, c *conf.Config, log klog.CtxLogger) server.Server {
	return hello.NewServer(s, options.ServerOptions(c, log)...)
}
