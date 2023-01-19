//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"examples/app/hello/internel/biz"
	"examples/app/hello/internel/conf"
	"examples/app/hello/internel/data"
	"examples/app/hello/internel/server"
	"examples/app/hello/internel/service"
	"github.com/cloudwego/kitex/pkg/klog"
	kserver "github.com/cloudwego/kitex/server"
	"github.com/google/wire"
)

//*polaris.Registry, *registry.Info
func initApp(klog.CtxLogger, *conf.Config) (kserver.Server, func(), error) {
	panic(wire.Build(service.ProviderSet, biz.ProviderSet, server.ProviderSet, data.ProviderSet))
}
