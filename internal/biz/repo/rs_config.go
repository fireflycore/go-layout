package repo

import (
	"context"

	"github.com/fireflycore/go-micro/config"
)

// ConfigRepo 定义读取运行时配置的仓储接口。
type ConfigRepo interface {
	// GetConfig 读取指定 group/key 下的原始配置 payload。
	GetConfig(ctx context.Context, appId, group, key string) (*config.Raw, error)
}
