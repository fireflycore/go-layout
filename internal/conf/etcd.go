package conf

import (
	"context"
	"errors"
	"go-layout/internal/biz/repo"
	"time"

	etcd "github.com/fireflycore/go-etcd"
)

type EtcdConfLoader struct {
	key   string
	group string

	bootstrapConf *BootstrapConf
	configRepo    repo.ConfigRepo
	utils         *Utils
}

func NewEtcdConfLoader(bootstrapConf *BootstrapConf, utils *Utils, configRepo repo.ConfigRepo) *EtcdConfLoader {
	return &EtcdConfLoader{
		key:           "etcd",
		group:         "database",
		bootstrapConf: bootstrapConf,
		configRepo:    configRepo,
		utils:         utils,
	}
}

func (load *EtcdConfLoader) Load() (*etcd.Conf, error) {
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

func (load *EtcdConfLoader) Local() (*etcd.Conf, error) {
	var dst etcd.Conf

	filePath := load.utils.GetConfigFilePath(load.key + ".json")
	if err := load.utils.LoadJSONConfig(filePath, &dst); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (load *EtcdConfLoader) Remote() (*etcd.Conf, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := load.configRepo.GetConfig(ctx, load.bootstrapConf.AppId, load.group, load.key)
	if err != nil {
		return nil, err
	}

	var dst etcd.Conf

	if err = load.utils.AnalyzeData(data.Content, []byte(load.bootstrapConf.AppSecret), &dst); err != nil {
		return nil, err
	}

	if err = load.utils.AnalyzeTlsData(load.key, dst.Tls); err != nil {
		return nil, err
	}

	return &dst, nil
}

func NewEtcdConf(loader *EtcdConfLoader) (*etcd.Conf, error) {
	return loader.Load()
}
