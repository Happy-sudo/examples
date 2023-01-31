package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/Happy-sudo/pkg/logger"
	"github.com/Happy-sudo/pkg/polaris"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexZap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"hello/internel/conf"
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
		var cuttingLogConfig = new(logger.CuttingLogConfig)
		cuttingLogConfig.Filename = config.Logger.Filename
		cuttingLogConfig.MaxSize = config.Logger.MaxSize
		cuttingLogConfig.MaxBackups = config.Logger.MaxBackups
		cuttingLogConfig.MaxAge = config.Logger.MaxAge
		cuttingLogConfig.Compress = config.Logger.Compress
		cuttingLogConfig.LocalTime = config.Logger.LocalTime
		klog.SetOutput(cuttingLogConfig.CuttingLogWriter())
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
