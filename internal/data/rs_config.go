package data

import (
	"context"
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"
	"sync"
	"time"

	"github.com/fireflycore/go-micro/rpc"
)

type configRepo struct {
	// 缓存配置
	cache sync.Map

	ctx context.Context

	confUtils     *conf.Utils
	bootstrapConf *conf.BootstrapConf

	configService config.ConfigServiceClient
}

func NewConfigRepo(
	ctx context.Context,

	confUtils *conf.Utils,
	bootstrapConf *conf.BootstrapConf,

	configService config.ConfigServiceClient,
) repo.ConfigRepo {
	return &configRepo{
		ctx: ctx,

		confUtils:     confUtils,
		bootstrapConf: bootstrapConf,

		configService: configService,
	}
}

func (ur *configRepo) GetConfig(appId, group, key string) (*config.Config, error) {
	ctx, cancel := context.WithTimeout(ur.ctx, time.Second*5)
	defer cancel()

	return rpc.WithRemoteInvoke[*config.Config, *config.GetConfigResponse](func() (*config.GetConfigResponse, error) {
		return ur.configService.GetConfig(ctx, &config.GetConfigRequest{
			AppId: appId,
			Group: group,
			Env:   ur.bootstrapConf.Env,
			Key:   key,
		})
	})
}
