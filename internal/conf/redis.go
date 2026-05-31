package conf

import (
	"context"
	"time"

	"github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/go-redis"
)

// NewRedisConfig 从 Consul Store 读取当前生效的 Redis 配置。
func NewRedisConfig(bootstrapConfig *BootstrapConfig, utils *Utils, store config.Store) (*redis.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dst, err := config.LoadStoreConfig[redis.Config](
		ctx,
		store,
		config.StoreParams{
			Key: config.Key{
				AppId: bootstrapConfig.App.Id,
				Env:   bootstrapConfig.App.Env,
				Group: "database",
				Key:   "redis",
			},
			AppSecret:  []byte(bootstrapConfig.App.Secret),
			Compressor: utils.Compressor(),
			Encryptor:  utils.Encryptor(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &dst, nil
}
