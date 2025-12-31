package conf

import (
	"context"
	"errors"
	"fmt"
	gorm "github.com/lhdhtrc/gorm/pkg"
	"go-layout/internal/biz/repo"
	"time"
)

type MysqlConfLoader struct {
	key   string
	group string

	bootstrapConf *BootstrapConf
	configRepo    repo.ConfigRepo
	utils         *Utils
}

func NewMysqlConfLoader(bootstrapConf *BootstrapConf, utils *Utils, configRepo repo.ConfigRepo) *MysqlConfLoader {
	return &MysqlConfLoader{
		key:           "mysql",
		group:         "database",
		bootstrapConf: bootstrapConf,
		configRepo:    configRepo,
		utils:         utils,
	}
}

func (load *MysqlConfLoader) Load() (*gorm.MysqlConf, error) {
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

func (load *MysqlConfLoader) Local() (*gorm.MysqlConf, error) {
	var dst gorm.MysqlConf

	filePath := load.utils.GetConfigFilePath(fmt.Sprintf("%s.json", load.key))
	if err := load.utils.LoadJSONConfig(filePath, &dst); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (load *MysqlConfLoader) Remote() (*gorm.MysqlConf, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := load.configRepo.GetConfig(ctx, load.bootstrapConf.AppId, load.group, load.key)

	if err != nil {
		return nil, err
	}

	var dst gorm.MysqlConf

	if err = load.utils.AnalyzeData(data.Content, []byte(load.bootstrapConf.AppSecret), &dst); err != nil {
		return nil, err
	}

	if err = load.utils.AnalyzeTlsData(load.key, dst.Tls); err != nil {
		return nil, err
	}

	return &dst, nil
}

func NewMysqlConf(loader *MysqlConfLoader) (*gorm.MysqlConf, error) {
	return loader.Load()
}
