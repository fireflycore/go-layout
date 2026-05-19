package conf

import (
	"context"
	"time"

	"github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/gormx"
)

// NewMysqlConfig 从数据面 Store 读取当前生效的 MySQL 配置。
func NewMysqlConfig(utils *Utils, bootstrapConfig *BootstrapConfig, store config.Store) (*gormx.MysqlConfig, error) {
	// 启动配置读取允许从根上下文派生超时，避免启动阶段无边界阻塞。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 统一走 LoadStoreConfig，还原 Base64 / 解密 / 解压后的目标配置结构。
	dst, err := config.LoadStoreConfig[gormx.MysqlConfig](
		ctx,
		store,
		config.StoreParams{
			Key: config.Key{
				AppId: bootstrapConfig.App.Id,
				Env:   bootstrapConfig.App.Env,
				Group: "database",
				Key:   "mysql",
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
