package dep

import (
	"context"
	"strings"

	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/constant"
	"github.com/fireflycore/go-micro/invocation"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// NewInvocationDNSManager 创建 config 服务统一复用的 DNS 管理器。
//
// 这里的职责只有一个：
// - 给业务侧补齐标准服务 DNS 的默认配置。
//
// 它不做服务发现，也不做实例选择。
func NewInvocationDNSManager(bootstrapConf *conf.BootstrapConf) *invocation.DNSManager {
	return invocation.NewDNSManager(&invocation.DNSConfig{
		// 默认命名空间来自当前服务启动配置。
		DefaultNamespace: bootstrapConf.GetServiceNamespace(),
		// 默认服务端口统一使用当前新模型的 9090。
		DefaultPort: invocation.DefaultServicePort,
	})
}

// NewInvocationConnectionManager 创建统一的出站连接管理器。
//
// 这里本质上只做三件事：
// - 基于标准 DNS target 建立连接；
// - 复用 grpc.ClientConn；
// - 统一挂接 OTel 与服务身份 metadata。
func NewInvocationConnectionManager(bootstrapConf *conf.BootstrapConf) (*invocation.ConnectionManager, error) {
	return invocation.NewConnectionManager(invocation.ConnectionManagerOptions{
		DNSManager: NewInvocationDNSManager(bootstrapConf),
		DialOptions: []grpc.DialOption{
			// 当前阶段仍然采用明文 gRPC，后续由 sidecar / mesh 处理链路安全。
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// 统一挂上 gRPC client 侧 OTel 采集。
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
			// 统一补服务身份 token，避免 repo 层重复处理。
			grpc.WithUnaryInterceptor(serviceAuthUnaryInterceptor(bootstrapConf.GetServiceAuthToken())),
		},
	})
}

// NewUnaryInvoker 把连接管理器封装成统一的 unary 调用入口。
func NewUnaryInvoker(manager *invocation.ConnectionManager) *invocation.UnaryInvoker {
	return &invocation.UnaryInvoker{Dialer: manager}
}

// serviceAuthUnaryInterceptor 在所有出站调用上统一补 Authorization。
//
// 这样 repo 层只表达“调哪个业务服务、调哪个 method”，
// 不需要每次都重复拼接服务身份 token。
func serviceAuthUnaryInterceptor(token string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 没有 token 时，直接透传调用。
		if strings.TrimSpace(token) == "" {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		// 优先复用已有 outgoing metadata。
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		// 仅在上游没有显式设置时补默认 Authorization。
		if len(md.Get(constant.Authorization)) == 0 {
			md.Set(constant.Authorization, token)
		}

		// 带着新的 outgoing metadata 继续发起调用。
		return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
	}
}
