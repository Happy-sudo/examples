package options

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/polaris"
	"greet/internel/conf"
	"net"
)

func ServerOptions(c *conf.Config, log klog.CtxLogger) []server.Option {
	var options []server.Option
	ctx := context.Background()
	//是否启用自定义端口
	if c.Server.Rpc.Enable {
		addr, err := net.ResolveTCPAddr(c.Server.Rpc.Network, c.Server.Rpc.Address)
		if err != nil {
			log.CtxErrorf(ctx, "ResolveTCPAddr error addr:%v err:%v", addr, err)
			return nil
		}
		options = append(options, server.WithServiceAddr(addr))
		log.CtxInfof(ctx, "服务端配置自定义端口已配置成功 %v", addr)
	}

	//是否启用北极星注册中心
	if c.Server.Polaris.Enable {
		r, err := polaris.NewPolarisRegistry(polaris.ServerOptions{})
		if err != nil {
			log.CtxErrorf(ctx, "polaris NewPolarisRegistry fatal ：%v", err)
			return nil
		}
		info := &registry.Info{
			ServiceName: c.Service.ServerName,
			Tags: map[string]string{
				polaris.NameSpaceTagKey: c.Service.NameSpace,
			},
		}
		options = append(options, server.WithRegistry(r))
		options = append(options, server.WithRegistryInfo(info))
		log.CtxInfof(ctx, "服务端配置北极星注册中心已配置成功 name：%v，nameSpace:%v", c.Service.ServerName, c.Service.NameSpace)
	}

	//是否启用jaeger链路追踪
	if c.Server.Jaeger.Enable {
		provider.NewOpenTelemetryProvider(
			provider.WithServiceName(c.Service.ServerName),
			provider.WithExportEndpoint(c.Server.Jaeger.Endpoint),
			provider.WithInsecure(),
		)
		//defer p.Shutdown(ctx)
		options = append(options, server.WithSuite(tracing.NewServerSuite()))
		options = append(options, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.Service.ServerName}))
		log.CtxInfof(ctx, "服务端配置链路已配置成功 ServiceName：%v，Endpoint:%v", c.Service.ServerName, c.Server.Jaeger.Endpoint)
	}

	//是否启用多路复用
	if c.Server.Transport.Enable {
		options = append(options, server.WithMuxTransport())
		log.CtxInfof(ctx, "服务端配置多路复用已配置成功 :%v", c.Server.Transport.Enable)
	}
	//是否启用限流器
	if c.Server.Limit.Enable {
		options = append(options, server.WithLimit(&limit.Option{
			MaxConnections: c.Server.Limit.MaxConnections,
			MaxQPS:         c.Server.Limit.MaxQPS,
		}))
		log.CtxInfof(ctx, "服务端配置限流器已配置成功 :%v", c.Server.Transport.Enable)
	}

	//埋点策略&埋点粒度
	if c.Server.StatsLevel.LevelBase {
		options = append(options, server.WithStatsLevel(stats.LevelBase))
		log.CtxInfof(ctx, "客户端配置启用基本埋点 已启用 LevelBase：%v", stats.LevelBase)
	}
	if c.Server.StatsLevel.LevelDetailed {
		options = append(options, server.WithStatsLevel(stats.LevelDetailed))
		log.CtxInfof(ctx, "客户端配置启用基本埋点和细粒度埋点 已启用 LevelDetailed：%v", stats.LevelDetailed)
	}
	if c.Server.StatsLevel.LevelDisabled {
		options = append(options, server.WithStatsLevel(stats.LevelDisabled))
		log.CtxInfof(ctx, "客户端配置禁用埋点 已禁用 LevelDisabled：%v", stats.LevelDisabled)
	}
	return options
}
