package conf

import (
	"context"
	"time"

	"github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/go-redis"
)

func NewRedisConfig(utils *Utils, bootstrapConfig *BootstrapConfig, store config.Store) (*redis.Conf, error) {
	// 启动配置读取允许从根上下文派生超时，避免启动阶段无边界阻塞。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dst, err := config.LoadStoreConfig[redis.Conf](
		ctx,
		store,
		config.StoreParams{
			AppId:      bootstrapConfig.App.Id,
			Env:        bootstrapConfig.App.Env,
			Group:      "database",
			Name:       "redis",
			AppSecret:  []byte(bootstrapConfig.App.Secret),
			Compressor: utils.Compressor(),
			Encryptor:  utils.Encryptor(),
		},
	)
	if err != nil {
		return nil, err
	}

	// TLS 字段直接使用配置中的本地证书路径，不再接受证书正文落盘的旧模式。
	return &dst, nil
}
