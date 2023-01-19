package main

import (
	"context"
	"encoding/json"
	"examples/app/hello/internel/conf"
	"examples/pkg/logger"
	"examples/pkg/polaris"
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexZap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
)

var (
	namespace = "examples"                 //空间名称
	fileGroup = "microservices"            //分组名称
	fileName  = "helloService/config.json" //文件名称
)

func main() {
	flag.Parse()

	klog.SetLogger(kitexZap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	configFile := polaris.ConfigApi(namespace, fileGroup, fileName)
	//解析远程配置文件
	config := new(conf.Config)
	err := json.Unmarshal([]byte(configFile.GetContent()), &config)
	if err != nil {
		klog.CtxErrorf(context.Background(), "json 反序列化失败 error：%v", err)
		panic(err)
	}

	// 自定义日志配置
	if config.Logger.Enable {
		klog.SetOutput(logger.CuttingLogWriter(config))
	}

	//wire 依赖注入
	svr, cleanup, err := initApp(klog.DefaultLogger(), config)
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := svr.Run(); err != nil {
		panic(err)
	}

}
