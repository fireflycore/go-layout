package data

import (
	"context"
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"

	micro "github.com/fireflycore/go-micro/rpc"
)

type configRepo struct {
	configService config.ConfigServiceClient

	bootstrapConf *conf.BootstrapConf
	confUtils     *conf.Utils
}

func NewConfigRepo(
	configService config.ConfigServiceClient,

	bootstrapConf *conf.BootstrapConf,
	confUtils *conf.Utils,
) repo.ConfigRepo {
	return &configRepo{
		configService: configService,

		bootstrapConf: bootstrapConf,
		confUtils:     confUtils,
	}
}

func (repo *configRepo) GetConfig(ctx context.Context, appId, group, key string) (*config.Config, error) {
	return micro.WithRemoteInvoke[*config.Config, *config.GetResponse](func() (*config.GetResponse, error) {
		return repo.configService.Get(ctx, &config.GetRequest{
			AppId: appId,
			Group: group,
			Env:   repo.bootstrapConf.Env,
			Key:   key,
		})
	})
}
