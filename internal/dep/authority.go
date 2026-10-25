package dep

import (
	"context"
	authToken "go-layout/dep/protobuf/gen/acme/auth/token/v1"
	"go-layout/internal/conf"

	"github.com/fireflycore/go-micro/authz"
	"github.com/fireflycore/go-micro/invocation"
)

// NewServiceAuthorityFetchFunc 创建当前服务凭证的获取函数，签发仍由 auth 服务负责。
//
// 获取函数只负责向 auth 服务换取凭证，不负责缓存、刷新和出站注入；
// 模板默认提供该装配点，确保新业务服务一开始就具备统一服务身份链路。
func NewServiceAuthorityFetchFunc(bootstrapConfig *conf.BootstrapConfig, manager *invocation.ConnectionManager) authz.ServiceAuthorityFetchFunc {
	// 直接复用普通调用的连接管理器，但获取凭证不经过依赖凭证的业务调用器。
	// 凭证申请只允许该精确方法，并清空业务身份和入站/出站元数据，防止递归与串号。
	bootstrap := invocation.NewBootstrapUnaryInvoker(manager, invocation.DefaultInvokeTimeout, authToken.AuthTokenService_GenerateServiceToken_FullMethodName)

	return func(ctx context.Context) (*authz.ServiceAuthorityToken, error) {
		var resp authToken.GenerateServiceTokenResponse

		dns := &invocation.DNS{
			// auth 是统一的 service token 签发服务。
			Service: "auth",
			// 默认与当前服务处于同一 namespace。
			Namespace: bootstrapConfig.Service.Namespace,
		}

		err := bootstrap.Invoke(ctx, dns, authToken.AuthTokenService_GenerateServiceToken_FullMethodName, &authToken.GenerateServiceTokenRequest{
			// 仅使用当前服务本地配置，不能采用请求中的用户或上一跳应用。
			AppId:     bootstrapConfig.App.Id,
			AppSecret: bootstrapConfig.App.Secret,
		}, &resp)
		if err != nil {
			return nil, err
		}
		if resp.GetData() == nil {
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
