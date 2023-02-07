package conf

import (
	kitexConf "github.com/Happy-sudo/pkg/conf/kitex_conf"
	hertzConf "github.com/baoyxing/hertz-contrib/pkg/config"
)

type Config struct {
	Service       kitexConf.Service `yaml:"service"`                             //服务名称配置
	ClientConnect ClientConnect     `yaml:"client_connect" yaml:"clientConnect"` //客户端服务发现配置
	Server        kitexConf.Server  `yaml:"server"`                              //服务端配置
	Client        kitexConf.Client  `yaml:"client"`                              //客户端配置
	Logger        kitexConf.Logger  `yaml:"logger"`                              //日志配置
	MysqlOptions  MysqlOptions      `yaml:"mysql_options" yaml:"mysqlOptions"`   // 数据库配置
	RedisOptions  RedisOptions      `yaml:"redis_options" yaml:"redisOptions"`   // redis配置

}

// ClientConnect 客户端服务发现配置
type ClientConnect struct {
	HelloService string `yaml:"hello_service" yaml:"helloService"`
	HelloClient  string `yaml:"Hello_client" yaml:"helloClient"`
}

type MysqlOptions struct {
	Enable bool `yaml:"enable" ` //是否启用数据库配置
	hertzConf.MysqlOptions
	ConnMaxIdleTime int `yaml:"conn_max_idle_time"`
}

type RedisOptions struct {
	Enable bool `yaml:"enable" ` //是否启用缓存配置
	hertzConf.RedisOptions
	Username        string `yaml:"user_name"`          //用户名
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"` //连接可能空闲的最长时间
	ConnMaxLifetime int    `yaml:"conn_max_life_time"` //可以增加连接的最大时间
	MaxIdleConns    int    `yaml:"max_idle_conns"`
}
