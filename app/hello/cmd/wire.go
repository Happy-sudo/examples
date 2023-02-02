//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kserver "github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"hello/internal/biz"
	"hello/internal/conf"
	"hello/internal/data"
	"hello/internal/server"
	"hello/internal/service"
)

//*polaris.Registry, *registry.Info
func initApp(klog.CtxLogger, *conf.Config) (kserver.Server, func(), error) {
	panic(wire.Build(service.ProviderSet, biz.ProviderSet, server.ProviderSet, data.ProviderSet))
}
