package conf

import (
	"errors"
	"github.com/fireflycore/go-redis"
	"go-layout/internal/biz/repo"
)

type RedisConfLoader struct {
	key   string
	group string

	utils         *Utils
	bootstrapConf *BootstrapConf

	configRepo repo.ConfigRepo
}

func NewRedisConfLoader(utils *Utils, bootstrapConf *BootstrapConf, configRepo repo.ConfigRepo) *RedisConfLoader {
	return &RedisConfLoader{
		key:           "redis",
		group:         "database",
		utils:         utils,
		bootstrapConf: bootstrapConf,
		configRepo:    configRepo,
	}
}

func (load *RedisConfLoader) Load() (*redis.Conf, error) {
	switch load.bootstrapConf.LoadConfMode {
	case "local":
		// 从本地加载
		return load.Local()
	case "remote":
		// 从配置中心加载
		return load.Remote()
	default:
		return nil, errors.New("not found load conf mode")
	}
}

func (load *RedisConfLoader) Local() (*redis.Conf, error) {
	var dst redis.Conf

	filePath := load.utils.GetConfigFilePath(load.key + ".json")
	if err := load.utils.LoadJSONConfig(filePath, &dst); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (load *RedisConfLoader) Remote() (*redis.Conf, error) {
	data, err := load.configRepo.GetConfig(load.bootstrapConf.AppId, load.group, load.key)

	if err != nil {
		return nil, err
	}

	var dst redis.Conf

	if err = load.utils.AnalyzeData(data.Content, []byte(load.bootstrapConf.AppSecret), &dst); err != nil {
		return nil, err
	}

	if err = load.utils.AnalyzeTlsData(load.key, dst.Tls); err != nil {
		return nil, err
	}

	return &dst, nil
}

func NewRedisConf(loader *RedisConfLoader) (*redis.Conf, error) {
	return loader.Load()
}
