package dep

import (
	"context"
	authToken "go-layout/dep/protobuf/gen/acme/auth/token/v1"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/authz"
	"github.com/fireflycore/go-micro/invocation"
)

// NewServiceAuthorityFetchFunc 创建当前服务获取 service token 的签发函数。
//
// fetch 只负责向 auth 服务换取 token，不负责缓存、刷新和出站注入；
// 模板默认提供该装配点，确保新业务服务一开始就具备统一服务身份链路。
func NewServiceAuthorityFetchFunc(bootstrapConfig *conf.BootstrapConfig, manager *invocation.ConnectionManager) authz.ServiceAuthorityFetchFunc {
	return func(ctx context.Context) (*authz.ServiceAuthorityToken, error) {
		// 获取 service token 不能走 UnaryInvoker，否则 provider 会递归依赖自身。
		conn, err := manager.Dial(ctx, &invocation.DNS{
			// auth 是统一的 service token 签发服务。
			Service: "auth",
			// 默认与当前服务处于同一 namespace。
			Namespace: bootstrapConfig.Service.Namespace,
		})
		if err != nil {
			return nil, err
		}

		// app_id/app_secret 来自当前服务启动配置，由 auth 校验后签发 service token。
		client := authToken.NewAuthTokenServiceClient(conn)
		resp, err := client.GenerateServiceToken(ctx, &authToken.GenerateServiceTokenRequest{
			AppId:     bootstrapConfig.App.Id,
			AppSecret: bootstrapConfig.App.Secret,
		})
		if err != nil {
			return nil, err
		}
		if resp == nil || resp.GetData() == nil {
			return nil, authz.ErrServiceAuthorityTokenMissing
		}

		// 把 auth 返回的 SessionToken 转成 go-micro provider 使用的缓存模型。
		return authz.NewServiceAuthorityToken(resp.GetData().GetToken(), resp.GetData().GetExpired())
	}
}

// NewServiceAuthorityProvider 是业务服务接入 service token 的标准装配点。
//
// manager 负责缓存、后台获取和按过期时间刷新 service token；
// 出站 metadata 清理、用户 authority 透传和 service authority 覆盖由 invocation 统一完成。
func NewServiceAuthorityProvider(bootstrapConfig *conf.BootstrapConfig, fetch authz.ServiceAuthorityFetchFunc) (authz.ServiceAuthorityManager, error) {
	// 使用 go-micro 的缓存型 provider，避免每次远程调用都访问 auth 服务。
	provider, err := authz.NewServiceAuthorityProvider(bootstrapConfig.ServiceAuthority, fetch)
	if err != nil {
		return nil, err
	}

	// 返回 manager 后由 AppServer 在 sidecar 生命周期中调用 Start(ctx)，不阻塞服务启动。
	return provider, nil
}
