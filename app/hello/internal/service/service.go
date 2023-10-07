package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
	"github.com/kitex-contrib/polaris"
	"xxx/internal/biz"
	"xxx/internal/conf"
)

var ProviderSet = wire.NewSet(NewXXXService)

type XXXService struct {
	biz *biz.XXXUseCase
	log klog.CtxLogger
}

func NewXXXService(biz *biz.XXXUseCase, log klog.CtxLogger) *XXXService {
	return &XXXService{
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

////NewExamplesClient 客户端连接模板 可根据业务自行调整
//func NewExamplesClient(c *conf.Config, suite *polaris.ClientSuite, log klog.CtxLogger) xxx.Client {
//
//	newClient, err := xxx.NewClient(c.ClientConnect.XXXService, options.ClientOptions(c, suite, log)...)
//	if err != nil {
//		log.CtxErrorf(context.Background(), "%s 客户端连接失败 err：%s", c.ClientConnect.XXXClient, err)
//	}
//	log.CtxInfof(context.Background(), "%s 客户端连接成功 ", c.ClientConnect.XXXClient)
//
//	return newClient
//}
