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
// 再按真实下游服务补充 DNS，避免模板默认制造 auth 依赖。
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
// fetchFunc 内部应使用 manager.Dial(...) 创建 auth token client，再调用
// GenerateServiceToken(app_id, app_secret)，最后返回
// authz.NewServiceAuthorityToken(resp.Data.Token, resp.Data.Expired)。
//
// 注意不要通过 UnaryInvoker 或 RemoteServiceManaged 获取服务 token：
// 它们会依赖当前 provider 注入 x-firefly-service-authority，容易形成递归依赖。
func NewServiceAuthorityProvider(bootstrapConfig *conf.BootstrapConfig, manager *invocation.ConnectionManager) authz.ServiceAuthorityProvider {
	// 保留 bootstrapConfig 参数，便于业务服务在接入时直接读取 app.id/app.secret。
	_ = bootstrapConfig
	// 保留 manager 参数，便于业务服务接入时直接拨 auth 服务，避免依赖 UnaryInvoker。
	_ = manager
	// nil 表示当前模板不自动注入 x-firefly-service-authority。
	return nil
}
