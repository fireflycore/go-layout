package conf

import (
	gorm "github.com/lhdhtrc/gorm/pkg"
	config "go-layout/depend/protobuf/gen/acme/config/v1"
)

// MysqlLoader 实现 ConfigLoader 接口
type MysqlLoader struct{}

func NewMysqlConf(dc *Conf) *gorm.MysqlConf {
	return dc.Mysql
}

func (ist *MysqlLoader) Local(dc *Conf) error {
	filePath := getConfigFilePath("mysql.json")

	if err := loadJSONConfig(filePath, &dc.Mysql.Conf); err != nil {
		return err
	}

	return nil
}

func (ist *MysqlLoader) Remote(dc *Conf, bc *BootstrapConf, cc config.ConfigServiceClient) error {
	content, err := fetchRemoteConfig(bc, cc, "database", "mysql")

	if err != nil {
		return err
	}

	if err = analyzeData(content, []byte(bc.AppSecret), &dc.Mysql.Conf); err != nil {
		return err
	}

	if err = analyzeTlsData("mysql", dc.Mysql.Conf.Tls); err != nil {
		return err
	}

	return nil
}
