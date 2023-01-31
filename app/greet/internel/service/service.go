package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/kitex-contrib/polaris"
	"greet/internel/biz"
	"greet/internel/conf"
	"greet/internel/pkg/options"
	"greet/kitex_gen/greet/v1/greet"
	"hello/kitex_gen/hello/v1/hello"
)

var ProviderSet = wire.NewSet(NewGreetService, NewDiscover, NewHelloClient, NewGreetClient)

type GreetService struct {
	biz *biz.GreetUseCase
	log klog.CtxLogger
}

func NewGreetService(biz *biz.GreetUseCase, log klog.CtxLogger) *GreetService {
	return &GreetService{
		biz: biz,
		log: log,
	}
}

//NewDiscover 服务发现
func NewDiscover(c *conf.Config, log klog.CtxLogger) *polaris.ClientSuite {

	ctx := context.Background()
	resolver, err := polaris.NewPolarisResolver(polaris.ClientOptions{})
	if err != nil {
		log.CtxErrorf(ctx, "NewPolarisResolver creates a polaris based resolver:%v error：%v", resolver, err)
		return nil
	}

	balancer, err := polaris.NewPolarisBalancer()
	if err != nil {
		log.CtxErrorf(ctx, "NewPolarisBalancer creates a polaris based balancer:%v error：%v", balancer, err)
		return nil
	}

	return &polaris.ClientSuite{
		DstNameSpace:       c.Service.NameSpace,
		Resolver:           resolver,
		Balancer:           balancer,
		ReportCallResultMW: polaris.NewUpdateServiceCallResultMW(),
	}
}

//NewHelloClient 客户端连接模板 可根据业务自行调整
func NewHelloClient(c *conf.Config, suite *polaris.ClientSuite, log klog.CtxLogger) hello.Client {
	return hello.MustNewClient(c.ClientConnect.HelloService, options.ClientOptions(c, suite, log)...)
}

//NewGreetClient 客户端连接模板 可根据业务自行调整
func NewGreetClient(c *conf.Config, suite *polaris.ClientSuite, log klog.CtxLogger) greet.Client {
	return greet.MustNewClient(c.Service.ServerName, options.ClientOptions(c, suite, log)...)
}
