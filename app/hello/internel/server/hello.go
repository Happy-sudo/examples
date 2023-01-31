package server

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"hello/internel/conf"
	"hello/internel/pkg/options"
	"hello/internel/service"
	"hello/kitex_gen/hello/v1/hello"
)

func NewRPCServer(s *service.HelloService, c *conf.Config, log klog.CtxLogger) server.Server {
	return hello.NewServer(s, options.ServerOptions(c, log)...)
}
