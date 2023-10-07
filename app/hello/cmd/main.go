package main

import (
	"bytes"
	"context"
	"flag"
	loggerZap "github.com/Happy-sudo/pkg/logger/klogzap"
	"github.com/Happy-sudo/pkg/polaris"
	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/yaml.v3"
	"xxx/internal/conf"
)

var (
	namespace = "examples"               //空间名称
	fileGroup = "microservices"          //分组名称
	fileName  = "xxxService/config.json" //文件名称
)

func main() {

	flag.Parse()

	configFile := polaris.ConfigApi(namespace, fileGroup, fileName)
	//解析远程配置文件
	config := new(conf.Config)
	err := yaml.Unmarshal([]byte(configFile.GetContent()), &config)

	if err != nil {
		klog.CtxErrorf(context.Background(), "yaml 解析失败 error：%v", err)
		panic(err)
	}

	// 自定义日志配置
	if config.Logger.Enable {
		buf := new(bytes.Buffer)
		loggers := loggerZap.NewZapLogger(&loggerZap.Zap{
			Directory:      config.Logger.Directory,      // 目录
			LoggerFileName: config.Logger.LoggerFileName, //文件名
			DirectoryType:  config.Logger.DirectoryType,  //日志类型/等级
			Suffix:         config.Logger.Suffix,         //后缀
			Day:            config.Logger.Day,            // 最大保存天数（天）
			CuttingTime:    config.Logger.CuttingTime,    // 按照时间切割（分钟）
			LoggerType:     config.Logger.LoggerType,     // 输出日志类型 Console/JSON
			ISConsole:      config.Logger.ISConsole,      // 是否输出到系统日志
		})
		klog.SetLogger(loggers)
		klog.SetLevel(klog.LevelTrace)
		klog.SetOutput(buf)
		defer loggers.Sync()
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
