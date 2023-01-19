package polaris

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/polarismesh/polaris-go"
	"github.com/polarismesh/polaris-go/pkg/model"
)

func ConfigApi(namespace, fileGroup, fileName string) model.ConfigFile {
	ctx := context.Background()
	//解析配置文件
	configAPI, err := polaris.NewConfigAPI()
	if err != nil {
		klog.CtxErrorf(ctx, "解析本地配置文件错误 ：%v", err)
		panic(err)
	}

	//获取远程配置文件
	configFile, err := configAPI.GetConfigFile(namespace, fileGroup, fileName)
	if err != nil {
		klog.CtxErrorf(ctx, "获取远程配置文件错误 ：%v", err)
		panic(err)
	}

	//监听器
	configFile.AddChangeListener(changeListener)

	return configFile
}

func changeListener(event model.ConfigFileChangeEvent) {
	klog.CtxInfof(context.Background(), "recevied change event. %+v", event)
}
