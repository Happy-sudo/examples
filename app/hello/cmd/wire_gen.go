// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"xxx/internal/biz"
	"xxx/internal/conf"
	"xxx/internal/data"
	server2 "xxx/internal/server"
	"xxx/internal/service"
)

// Injectors from wire.go:

//*polaris.Registry, *registry.Info
func initApp(ctxLogger klog.CtxLogger, config *conf.Config) (server.Server, func(), error) {
	client := data.NewDBClient(ctxLogger, config)
	redisClient := data.NewRedisClient(ctxLogger, config)
	dataData, cleanup, err := data.NewData(ctxLogger, client, redisClient)
	if err != nil {
		return nil, nil, err
	}
	xxxRepo := data.NewXXXRepo(dataData)
	xxxUseCase := biz.NewXXXUseCase(xxxRepo, ctxLogger)
	xxxService := service.NewXXXService(xxxUseCase, ctxLogger)
	serverServer := server2.NewRPCServer(xxxService, config, ctxLogger)
	return serverServer, func() {
		cleanup()
	}, nil
}
