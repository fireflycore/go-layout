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
// 这里维护的是“当前服务依赖了哪些远程业务服务”的总表，后续新增
// 下游时继续在 dep 包集中扩展；`internal/data/rs_*.go` 只负责按业务服务
// 名绑定 caller，不再承载多服务注册表本身。
func NewRemoteServiceManaged(invoker *invocation.UnaryInvoker, bootstrapConfig *conf.BootstrapConfig) *invocation.RemoteServiceManaged {
	return invocation.NewRemoteServiceManaged(
		invoker,
		invocation.DNS{
			Service:   "auth",
			Namespace: bootstrapConfig.Service.Namespace,
		},
	)
}

// NewUnaryInvoker 把连接管理器封装成统一的 unary 调用入口。
//
// 当前 go-micro/invocation 负责：
// - 透传用户 authority 和短 TTL authz sign；
// - 清理上一跳普通身份 metadata；
// - 在配置 provider 后覆盖当前服务的 service authority。
func NewUnaryInvoker(manager *invocation.ConnectionManager, provider authz.ServiceAuthorityProvider) *invocation.UnaryInvoker {
	// ServiceAuthorityProvider 当前默认为 nil，表示模板只清理旧上下文，不伪造服务身份。
	//
	// 具体业务服务接入 auth token proto 后，应在 NewServiceAuthorityProvider 中
	// 调用 auth 服务 GenerateServiceToken，并把 provider 注入到这里。
	return invocation.NewUnaryInvoker(manager, invocation.DefaultInvokeTimeout).
		WithServiceAuthorityProvider(provider)
}

// NewServiceAuthorityProvider 是业务服务接入 service token 的标准装配点。
//
// go-layout 模板不生成 auth token proto，因此这里不伪造 token，也不把 app.instance_id
// 写进任何身份头。实际服务需要生成 acme.auth.token.v1 后，在此处构造：
//
//	authz.NewServiceAuthorityProvider(&authz.ServiceAuthorityConfig{...}, fetchFunc)
//
// fetchFunc 内部调用 auth 服务 GenerateServiceToken(app_id, app_secret)，并返回
// authz.NewServiceAuthorityToken(resp.Data.Token, resp.Data.Expired)。
func NewServiceAuthorityProvider(bootstrapConfig *conf.BootstrapConfig) authz.ServiceAuthorityProvider {
	// 保留 bootstrapConfig 参数，便于业务服务在接入时直接读取 app.id/app.secret。
	_ = bootstrapConfig
	// nil 表示当前模板不自动注入 x-firefly-service-authority。
	return nil
}
