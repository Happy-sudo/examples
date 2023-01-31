package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/kitex-contrib/polaris"
	"hello/internel/biz"
	"hello/internel/conf"
	"hello/internel/pkg/options"
	"hello/kitex_gen/hello/v1/hello"
)

var ProviderSet = wire.NewSet(NewHelloService)

type HelloService struct {
	biz *biz.HelloUseCase
	log klog.CtxLogger
}

func NewHelloService(biz *biz.HelloUseCase, log klog.CtxLogger) *HelloService {
	return &HelloService{
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

//NewExamplesClient 客户端连接模板 可根据业务自行调整
func NewExamplesClient(c *conf.Config, suite *polaris.ClientSuite, log klog.CtxLogger) hello.Client {
	return hello.MustNewClient(c.Service.ServerName, options.ClientOptions(c, suite, log)...)
}
