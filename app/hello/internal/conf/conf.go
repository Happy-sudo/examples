package conf

import (
	kitexConf "github.com/Happy-sudo/pkg/conf/kitex_conf"
	hertzConf "github.com/baoyxing/hertz-contrib/pkg/config"
)

type Config struct {
	Service       kitexConf.Service `json:"service"`                             //服务名称配置
	ClientConnect ClientConnect     `json:"client_connect" json:"clientConnect"` //客户端服务发现配置
	Server        kitexConf.Server  `json:"server"`                              //服务端配置
	Client        kitexConf.Client  `json:"client"`                              //客户端配置
	Logger        kitexConf.Logger  `json:"logger"`                              //日志配置
	MysqlOptions  MysqlOptions      `json:"mysql_options" json:"mysqlOptions"`   // 数据库配置
	RedisOptions  RedisOptions      `json:"redis_options" json:"redisOptions"`   // redis配置
}

// ClientConnect 客户端服务发现配置
type ClientConnect struct {
	HelloService string `json:"hello_service" json:"helloService"`
	HelloClient  string `json:"Hello_client" json:"helloClient"`
}
type MysqlOptions struct {
	Enable bool `mapstructure:"enable" ` //是否启用数据库配置
	hertzConf.MysqlOptions
	ConnMaxIdleTime int `mapstructure:"conn_max_idle_time"`
}
type RedisOptions struct {
	Enable bool `mapstructure:"enable" ` //是否启用缓存配置
	hertzConf.RedisOptions
	Username        string `mapstructure:"user_name"`          //用户名
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time"` //连接可能空闲的最长时间
	ConnMaxLifetime int    `mapstructure:"conn_max_life_time"` //可以增加连接的最大时间
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`     //最大空闲连接数
}
