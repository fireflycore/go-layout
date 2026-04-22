package data

import (
	"context"
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"
	"sync"
	"time"

	"github.com/fireflycore/go-micro/invocation"
)

type configRepo struct {
	// 缓存配置
	cache sync.Map

	confUtils     *conf.Utils
	bootstrapConf *conf.BootstrapConf

	caller *invocation.RemoteServiceCaller
}

func NewConfigRepo(
	confUtils *conf.Utils,
	bootstrapConf *conf.BootstrapConf,
	invoker *invocation.UnaryInvoker,
) repo.ConfigRepo {
	return &configRepo{
		confUtils:     confUtils,
		bootstrapConf: bootstrapConf,
		caller: invocation.NewRemoteServiceCaller(
			invoker,
			&invocation.ServiceDNS{
				Service:   "config",
				Namespace: bootstrapConf.GetServiceNamespace(),
			},
			invocation.BuildInvocationContextFromContext,
		),
	}
}

func (ur *configRepo) GetConfig(appId, group, key string) (*config.Config, error) {
	// 启动阶段读取配置属于后台初始化流程，允许基于根上下文派生超时。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var res config.GetConfigResponse

	err := ur.caller.Invoke(
		ctx,
		config.ConfigService_GetConfig_FullMethodName,
		&config.GetConfigRequest{
			AppId: appId,
			Group: group,
			Env:   ur.bootstrapConf.Env,
			Key:   key,
		},
		&res,
	)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
