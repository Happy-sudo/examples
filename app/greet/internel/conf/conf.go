package conf

import (
	kitexConf "github.com/Happy-sudo/pkg/conf/kitex_conf"
)

type Config struct {
	Service       kitexConf.Service       `json:"service"`                             //服务名称配置
	ClientConnect kitexConf.ClientConnect `json:"client_connect" json:"clientConnect"` //客户端服务发现配置
	Server        kitexConf.Server        `json:"server"`                              //服务端配置
	Client        kitexConf.Client        `json:"client"`                              //客户端配置
	Logger        kitexConf.Logger        `json:"logger"`                              //日志配置
}
