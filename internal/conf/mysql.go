package conf

import (
	"errors"
	"github.com/fireflycore/gormx"
	"go-layout/internal/biz/repo"
)

type MysqlConfLoader struct {
	key   string
	group string

	utils         *Utils
	bootstrapConf *BootstrapConf

	configRepo repo.ConfigRepo
}

func NewMysqlConfLoader(utils *Utils, bootstrapConf *BootstrapConf, configRepo repo.ConfigRepo) *MysqlConfLoader {
	return &MysqlConfLoader{
		key:           "mysql",
		group:         "database",
		utils:         utils,
		bootstrapConf: bootstrapConf,
		configRepo:    configRepo,
	}
}

func (load *MysqlConfLoader) Load() (*gormx.MysqlConf, error) {
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

func (load *MysqlConfLoader) Local() (*gormx.MysqlConf, error) {
	var dst gormx.MysqlConf

	filePath := load.utils.GetConfigFilePath(load.key + ".json")
	if err := load.utils.LoadJSONConfig(filePath, &dst); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (load *MysqlConfLoader) Remote() (*gormx.MysqlConf, error) {
	data, err := load.configRepo.GetConfig(load.bootstrapConf.AppId, load.group, load.key)

	if err != nil {
		return nil, err
	}

	var dst gormx.MysqlConf

	if err = load.utils.AnalyzeData(data.Content, []byte(load.bootstrapConf.AppSecret), &dst); err != nil {
		return nil, err
	}

	if err = load.utils.AnalyzeTlsData(load.key, dst.Tls); err != nil {
		return nil, err
	}

	return &dst, nil
}

func NewMysqlConf(loader *MysqlConfLoader) (*gormx.MysqlConf, error) {
	return loader.Load()
}
