package test

import (
	"context"
	"fmt"
	"github.com/Happy-sudo/pkg/polaris"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexZap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"testing"
	"time"
	"xxx/internal/conf"
	"xxx/internal/service"
	v1 "xxx/kitex_gen/xxx/v1"
)

var (
	namespace = "examples"               //空间名称
	fileGroup = "microservices"          //分组名称
	fileName  = "XXXService/config.json" //文件名称
)

func TestClientXXX(t *testing.T) {

	klog.SetLogger(kitexZap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	ctx := context.Background()

	configFile := polaris.ConfigApi(namespace, fileGroup, fileName)
	//解析远程配置文件
	config := new(conf.Config)
	err := sonic.Unmarshal([]byte(configFile.GetContent()), &config)
	if err != nil {
		klog.CtxErrorf(ctx, "json 反序列化失败 error：%v", err)
		panic(err)
	}
	newClient := service.NewExamplesClient(config, service.NewDiscover(config, klog.DefaultLogger()), klog.DefaultLogger())
	//option := polaris.ClientOptions{}
	//r, err := polaris.NewPolarisResolver(option)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//pb, err := polaris.NewPolarisBalancer()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//suite := &polaris.ClientSuite{
	//	DstNameSpace:       XXXConfig.Service.NameSpace,
	//	Resolver:           r,
	//	Balancer:           pb,
	//	ReportCallResultMW: polaris.NewUpdateServiceCallResultMW(),
	//}
	//
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(XXXConfig.Service.ClientName),
	//	// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
	//	provider.WithExportEndpoint(XXXConfig.Server.Jaeger.Endpoint),
	//	provider.WithInsecure(),
	//)
	//defer p.Shutdown(context.Background())
	//
	//var options []client.Option
	//options = append(options, client.WithHostPorts("127.0.0.1:4441"))
	//options = append(options, client.WithSuite(tracing.NewClientSuite()))
	//options = append(options, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: XXXConfig.Service.ClientName}))
	//options = append(options, client.WithSuite(suite))
	//options = append(options, client.WithRPCTimeout(time.Second*1))
	//newClient, _ := XXX.NewClient(XXXConfig.Service.ServerName,
	//	//client.WithSuite(tracing.NewClientSuite()),
	//	//client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: XXXConfig.Service.ClientName}),
	//	//client.WithSuite(suite),
	//	//client.WithRPCTimeout(time.Second*1),
	//	options...,
	//)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		resp, err := newClient.XXX(ctx, &v1.Request{Message: "Hi,polaris!"})
		fmt.Println(resp, err)
		if err != nil {
			t.Log(err)
		}
		t.Log(resp)
		cancel()
		time.Sleep(1 * time.Second)
	}

}
