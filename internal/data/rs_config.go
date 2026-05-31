package data

import (
	"context"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"

	microConfig "github.com/fireflycore/go-micro/config"
)

// configRepo 负责从统一配置 Store 读取并解码对象存储配置。
type configRepo struct {
	store  microConfig.Store
	client microConfig.Client

	confUtils       *conf.Utils
	bootstrapConfig *conf.BootstrapConfig
}

// NewConfigRepo 创建当前服务使用的运行时配置读取仓库。
//
// 这里优先复用 `go-consul/config.Client`：
// - 把本地 cache 统一下沉到基础库；
// - 避免业务侧继续维护额外 `sync.Map`；
// - 让后续 watch / TTL 策略继续留在基础库演进。
func NewConfigRepo(store microConfig.Store, client microConfig.Client, confUtils *conf.Utils, bootstrapConfig *conf.BootstrapConfig) repo.ConfigRepo {
	return &configRepo{
		store:           store,
		client:          client,
		confUtils:       confUtils,
		bootstrapConfig: bootstrapConfig,
	}
}

// loadRaw 优先通过 Client.Get 读取配置；无客户端时退回 Store.Get。
func (c *configRepo) loadRaw(ctx context.Context, key microConfig.Key) (*microConfig.Raw, error) {
	if c.client != nil {
		return c.client.Get(ctx, key)
	}

	return c.store.Get(ctx, key)
}

// GetConfig 读取指定配置键的原始 payload。
func (r *configRepo) GetConfig(ctx context.Context, appId, group, key string) (*microConfig.Raw, error) {
	// 配置环境固定使用当前服务运行环境，调用方只选择 app/group/key。
	k := microConfig.Key{
		AppId: appId,
		Env:   r.bootstrapConfig.App.Env,
		Group: group,
		Key:   key,
	}

	raw, err := r.loadRaw(ctx, k)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
