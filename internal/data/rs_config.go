package data

import (
	"context"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"
	"sync"
)

type configRepo struct {
	configService config.ConfigServiceClient

	bc *conf.BootstrapConf
	cu *conf.Utils

	// 缓存配置
	tokenCache sync.Map
}

func NewConfigRepo(
	configService config.ConfigServiceClient,

	bc *conf.BootstrapConf,
	cu *conf.Utils,
) repo.ConfigRepo {
	return &configRepo{
		configService: configService,

		bc: bc,
		cu: cu,
	}
}

func (repo *configRepo) GetConfig(ctx context.Context, appId, group, key string) (*config.Config, error) {
	return micro.WithRemoteInvoke[*config.Config, *config.GetResponse](func() (*config.GetResponse, error) {
		return repo.configService.Get(ctx, &config.GetRequest{
			AppId: appId,
			Group: group,
			Env:   repo.bc.Env,
			Key:   key,
		})
	})
}
