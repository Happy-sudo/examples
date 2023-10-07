package server

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"xxx/internal/conf"
	"xxx/internal/pkg/options"
	"xxx/internal/service"
	"xxx/kitex_gen/xxx/v1/xxx"
)

func NewRPCServer(s *service.XXXService, c *conf.Config, log klog.CtxLogger) server.Server {
	return xxx.NewServer(s, options.ServerOptions(c, log)...)
}
