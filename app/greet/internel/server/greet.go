package server

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"greet/internel/conf"
	"greet/internel/pkg/options"
	"greet/internel/service"
	"greet/kitex_gen/greet/v1/greet"
)

func NewRPCServer(s *service.GreetService, c *conf.Config, log klog.CtxLogger) server.Server {
	return greet.NewServer(s, options.ServerOptions(c, log)...)
}
