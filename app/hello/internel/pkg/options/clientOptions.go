package options

import (
	"context"
	"examples/app/hello/internel/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/polaris"
	"time"
)

type ctxKey int

const (
	ctxConsistenKey ctxKey = iota
)

func ClientOptions(c *conf.Config, suite *polaris.ClientSuite, log klog.CtxLogger) []client.Option {
	ctx := context.Background()
	var options []client.Option
	//超时配置
	if len(c.Client.TimeoutControl.RpcTimeout.Timeout) > 1 && c.Client.TimeoutControl.RpcTimeout.Enable {
		duration, err := time.ParseDuration(c.Client.TimeoutControl.RpcTimeout.Timeout)
		if err != nil {
			log.CtxErrorf(context.Background(), "ParseDuration RpcTimeout duration:%v error：%v", duration, err)
			return nil
		}
		options = append(options, client.WithRPCTimeout(duration))
		log.CtxInfof(ctx, "客户端配置RPC超时配置已配置成功 %v", duration)
	} else {
		// 未配置超时，则默认 1s
		options = append(options, client.WithRPCTimeout(time.Second*1))
		log.CtxInfof(ctx, "客户端配置RPC超时配置已启用默认配置 %v", time.Second*1)
	}

	//连接超时
	if len(c.Client.TimeoutControl.ConnectTimeOut.TimeOut) > 1 && c.Client.TimeoutControl.ConnectTimeOut.Enable {
		duration, err := time.ParseDuration(c.Client.TimeoutControl.ConnectTimeOut.TimeOut)
		if err != nil {
			log.CtxErrorf(ctx, "ParseDuration ConnectTimeOut duration:%v error：%v", duration, err)
			return nil
		}
		options = append(options, client.WithConnectTimeout(duration))
		log.CtxInfof(ctx, "客户端配置连接超时已配置成功 %v", duration)
	} else {
		// 未配置超时，则默认 1s
		options = append(options, client.WithConnectTimeout(time.Millisecond*50))
		log.CtxInfof(ctx, "客户端配置连接超时配置已启用默认配置 %v", time.Millisecond*50)
	}

	//客户端配置jaeger
	if c.Server.Jaeger.Enable {
		provider.NewOpenTelemetryProvider(
			provider.WithServiceName(c.Service.ClientName),
			provider.WithExportEndpoint(c.Server.Jaeger.Endpoint),
			provider.WithEnableTracing(true),
			provider.WithInsecure(),
		)
		//defer p.Shutdown(ctx)
		options = append(options, client.WithSuite(tracing.NewClientSuite()))
		options = append(options, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.Service.ClientName}))
		log.CtxInfof(ctx, "客户端配置链路已配置成功 ClientName：%v，Endpoint:%v", c.Service.ClientName, c.Server.Jaeger.Endpoint)
	}

	// 客户端连接服务中心
	if c.Server.Polaris.Enable {
		options = append(options, client.WithSuite(suite))
		log.CtxInfof(ctx, "客户端配置连接配置中心已配置成功 ClientName：%v，Endpoint:%v", c.Service.ClientName, c.Server.Jaeger.Endpoint)
	}

	//连接类型
	if c.Client.ConnectionType.ShortConnection.Enable {
		//短链接
		options = append(options, client.WithShortConnection())
		log.CtxInfof(ctx, "客户端配置短链接配置已启用 %v", c.Client.ConnectionType.ShortConnection.Enable)
	}

	if c.Client.ConnectionType.LongConnection.Enable {
		duration, err := time.ParseDuration(c.Client.ConnectionType.LongConnection.MaxIdleTimeOut)
		if err != nil {
			log.CtxErrorf(ctx, "ParseDuration LongConnection MaxIdleTimeOut duration:%v error：%v", duration, err)
			return nil
		}
		pool := connpool.IdleConfig{
			MinIdlePerAddress: c.Client.ConnectionType.LongConnection.MinIdlePerAddress,
			MaxIdlePerAddress: c.Client.ConnectionType.LongConnection.MaxIdlePerAddress,
			MaxIdleGlobal:     c.Client.ConnectionType.LongConnection.MaxIdleGlobal,
			MaxIdleTimeout:    duration,
		}
		//长链接
		options = append(options, client.WithLongConnection(pool))
		log.CtxInfof(ctx, "客户端配置长链接配置已启用 连接池：%v", pool)

	}

	//客户端配置多路复用
	if c.Client.ConnectionType.ClientTransport.Enable {
		options = append(options, client.WithMuxConnection(c.Client.ConnectionType.ClientTransport.MuxConnection))
		log.CtxInfof(ctx, "客户端配置多路复用 已启用 ：%v", c.Client.ConnectionType.ClientTransport.MuxConnection)
	}

	//请求重试机制
	if c.Client.FailureRetry.Enable {
		failurePolicy := retry.NewFailurePolicy()
		failurePolicy.WithMaxRetryTimes(c.Client.FailureRetry.MaxRetryTimes)
		options = append(options, client.WithFailureRetry(failurePolicy))
		log.CtxInfof(ctx, "客户端配置请求重试机制 已启用 重试次数：%v", c.Client.FailureRetry.MaxRetryTimes)
	}

	//负载均衡
	if c.Client.LoadBalancer.Enable {
		options = append(options, client.WithLoadBalancer(loadbalance.NewConsistBalancer(
			loadbalance.NewConsistentHashOption(func(ctx context.Context, request interface{}) string {
				return ctx.Value(ctxConsistenKey).(string)
			}))))
		log.CtxInfof(ctx, "客户端配置负载均衡 已启用 ：%v", c.Client.LoadBalancer.Enable)
	}
	//熔断器
	if c.Client.CBSuite.Enable {
		//options = append(options, client.WithCircuitBreaker(circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		//	return ""
		//})))
	}
	//埋点策略&埋点粒度
	if c.Client.StatsLevel.LevelBase {
		options = append(options, client.WithStatsLevel(stats.LevelBase))
		log.CtxInfof(ctx, "客户端配置启用基本埋点 已启用 LevelBase：%v", stats.LevelBase)
	}
	if c.Client.StatsLevel.LevelDetailed {
		options = append(options, client.WithStatsLevel(stats.LevelDetailed))
		log.CtxInfof(ctx, "客户端配置启用基本埋点和细粒度埋点 已启用 LevelDetailed：%v", stats.LevelDetailed)
	}
	if c.Client.StatsLevel.LevelDisabled {
		options = append(options, client.WithStatsLevel(stats.LevelDisabled))
		log.CtxInfof(ctx, "客户端配置禁用埋点 已禁用 LevelDisabled：%v", stats.LevelDisabled)
	}
	return options
}
