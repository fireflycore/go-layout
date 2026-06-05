package dep

import (
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/authz"
	"github.com/fireflycore/go-micro/invocation"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewInvocationDNSManager 创建服务统一复用的 DNS 管理器。
//
// 这里的职责只有一个：
// - 给业务侧补齐标准服务 DNS 的默认配置。
//
// 它不做服务发现，也不做实例选择。
func NewInvocationDNSManager(bootstrapConfig *conf.BootstrapConfig) *invocation.DNSManager {
	return invocation.NewDNSManager(&invocation.DNSConfig{
		// 默认命名空间来自当前服务启动配置。
		DefaultNamespace: bootstrapConfig.Service.Namespace,
		// 默认服务端口统一使用当前新模型的 9090。
		DefaultPort: invocation.DefaultServicePort,
	})
}

// NewInvocationConnectionManager 创建统一的出站连接管理器。
//
// 这里本质上只做三件事：
// - 基于标准 DNS target 建立连接；
// - 复用 grpc.ClientConn；
// - 统一挂接 OTel client 侧采集。
func NewInvocationConnectionManager(bootstrapConfig *conf.BootstrapConfig) (*invocation.ConnectionManager, error) {
	return invocation.NewConnectionManager(invocation.ConnectionManagerOptions{
		DNSManager: NewInvocationDNSManager(bootstrapConfig),
		DialOptions: []grpc.DialOption{
			// 当前阶段仍然采用明文 gRPC，后续由 sidecar / mesh 处理链路安全。
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			// 统一挂上 gRPC client 侧 OTel 采集。
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		},
	})
}

// NewRemoteServiceManaged 在启动期集中登记本服务依赖的远程业务服务。
//
// 默认 demo 不登记任何远程业务服务；后续业务服务新增 rs_*.go 时，
// 再按真实下游服务补充 DNS。service token 获取由 authority.go 直连 auth，不放入远程业务注册表。
func NewRemoteServiceManaged(invoker *invocation.UnaryInvoker, bootstrapConfig *conf.BootstrapConfig) *invocation.RemoteServiceManaged {
	// 保留 bootstrapConfig 参数，便于新增下游服务时直接读取 namespace。
	_ = bootstrapConfig
	return invocation.NewRemoteServiceManaged(invoker)
}

// NewUnaryInvoker 把连接管理器封装成统一的 unary 调用入口。
//
// 当前 go-micro/invocation 负责：
// - 透传用户 authority 和短 TTL authz sign；
// - 清理上一跳普通身份 metadata；
// - 在配置 manager 后覆盖当前服务的 service authority。
func NewUnaryInvoker(manager *invocation.ConnectionManager, provider authz.ServiceAuthorityManager) *invocation.UnaryInvoker {
	// manager 来自 NewServiceAuthorityProvider，负责读取后台刷新的当前服务 service token。
	return invocation.NewUnaryInvoker(manager, invocation.DefaultInvokeTimeout).
		WithServiceAuthorityProvider(provider)
}
