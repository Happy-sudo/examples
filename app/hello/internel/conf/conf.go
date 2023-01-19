package conf

type Config struct {
	Service service `json:"service"` //服务名称配置
	Server  server  `json:"server"`  //服务端配置
	Client  client  `json:"client"`  //客户端配置
	Logger  logger  `json:"logger"`  //日志配置
}

type logger struct {
	Enable     bool   `json:"enable" `                       //是否启用自定义日志配置
	Filename   string `json:"file_name" json:"fileName"`     //路径
	MaxSize    int    `json:"max_size" json:"maxSize"`       //日志的最大大小（M）
	MaxBackups int    `json:"max_backups" json:"maxBackups"` //日志的最大保存数量
	MaxAge     int    `json:"max_age" json:"maxAge"`         //日志文件存储最大天数
	Compress   bool   `json:"compress"`                      //是否执行压缩
	LocalTime  bool   `json:"local_time" json:"localTime"`   //是否使用格式化时间辍
}

//服务端配置
type server struct {
	Rpc        rpc        `json:"rpc"`                           //服务ip配置
	Polaris    polaris    `json:"polaris"`                       //北极星注册中心配置
	Jaeger     jaeger     `json:"jaeger"`                        //链路配置
	Transport  transport  `json:"transport"`                     //多路复用配置
	Limit      limit      `json:"limit"`                         //限流器
	StatsLevel statsLevel `json:"stats_level" json:"statsLevel"` //埋点策略&埋点粒度
}

//服务名称配置
type service struct {
	NameSpace  string `json:"namespace"`                     //服务空间名称
	ServerName string `json:"server_name" json:"serverName"` //服务名称
	ClientName string `json:"client_name" json:"clientName"` //客户端名称
	Version    string `json:"version"`                       //版本信息
}

//服务地址端口配置
type rpc struct {
	Enable  bool   `json:"enable" `                 //是否启用rpc自定义配置
	Address string `json:"address"`                 //地址
	Network string `json:"net_work" json:"netWork"` //连接方式 (tcp udp)
}

// 注册中心配置
type polaris struct {
	Enable bool `json:"enable"` //是否启用注册中心，默认开启
}

//链路追踪配置
type jaeger struct {
	Enable   bool   `json:"enable"`   //是否启用链路追踪
	Endpoint string `json:"endpoint"` //地址
}

//多路复用配置
type transport struct {
	Enable bool `json:"enable"` //是否启用多路复用
}

//限流器配置
type limit struct {
	Enable         bool `json:"enable"`          //是否启用多路复用
	MaxConnections int  `json:"max_connections"` // 最大连接数
	MaxQPS         int  `json:"max_qps"`         //最大qps
}

// **********************************公共对象*******************************

type statsLevel struct {
	LevelDisabled bool `json:"level_disabled"`
	LevelBase     bool `json:"level_base"`
	LevelDetailed bool `json:"level_detailed"`
}

// **********************************客户端对象******************************
//客户端配置
type client struct {
	TimeoutControl timeOutControl `json:"timeout_control" json:"timeoutControl"` //超时控制
	ConnectionType connectionType `json:"connection_type" json:"connectionType"` // 连接类型
	FailureRetry   failureRetry   `json:"failure_retry" json:"failureRetry"`     //请求重试
	LoadBalancer   loadBalancer   `json:"load_balancer" json:"loadBalancer"`     //负载均衡
	CBSuite        cbsuite        `json:"cbsuite"`                               //熔断器
	StatsLevel     statsLevel     `json:"stats_level"`                           //埋点策略&埋点粒度
}

//超时控制
type timeOutControl struct {
	RpcTimeout     rpcTimeout     `json:"rpc_timeout" json:"rpcTimeout"`
	ConnectTimeOut connectTimeOut `json:"connect_time_out" json:"connectTimeOut"`
}

//连接类型（长链接 短链接）
type connectionType struct {
	ShortConnection shortConnection `json:"short_connection" json:"shortConnection"` //短链接
	LongConnection  longConnection  `json:"long_connection" json:"longConnection"`   //长链接
	ClientTransport clientTransport `json:"transport"`                               //客户端多路复用

}

//rpc超时控制
type rpcTimeout struct {
	Enable  bool   `json:"enable"`                  //是否启用rpc超时
	Timeout string `json:"time_out" json:"timeout"` //超时时间 （默认 1s 单位："ns", "us" (or "µs"), "ms", "s", "m", "h"）
}

//connect超时控制
type connectTimeOut struct {
	Enable  bool   `json:"enable"`                  //是否启用rpc超时
	TimeOut string `json:"time_out" json:"timeOut"` //连接超时 （默认：50ms）
}

//短链接
type shortConnection struct {
	Enable bool `json:"enable"` //是否启用短链接
}

//长链接
type longConnection struct {
	Enable            bool   `json:"enable"`                                        //是否启用长链接
	MaxIdlePerAddress int    `json:"max_idle_per_address" json:"maxIdlePerAddress"` //最大空闲地址
	MinIdlePerAddress int    `json:"min_idle_per_address" json:"minIdlePerAddress"` //最小空闲地址
	MaxIdleGlobal     int    `json:"max_idle_global" json:"maxIdleGlobal"`          //最大空闲数
	MaxIdleTimeOut    string `json:"max_idle_time_out" json:"maxIdleTimeOut"`       //最大空闲超时
}

// 客户端多路复用
type clientTransport struct {
	Enable        bool `json:"enable"`                              //是否启用多路复用
	MuxConnection int  `json:"mux_connection" json:"muxConnection"` //连接数
}

//重试机制
type failureRetry struct {
	Enable        bool `json:"enable"`                                 //是否启用请求重试机制
	MaxRetryTimes int  `json:"max_retry_times" json:"max_retry_times"` //重试次数
}

//负载均衡
type loadBalancer struct {
	Enable bool `json:"enable"` //是否启用负载均衡
}

//熔断器
type cbsuite struct {
	Enable bool `json:"enable"` //是否启用熔断器
}
