package conf

import (
	gorm "github.com/lhdhtrc/gorm/pkg"
	configCenter "go-layout/dep/protobuf/gen/acme/config/v1"
)

// MysqlLoader 实现 ConfigLoader 接口
type MysqlLoader struct{}

func NewMysqlConf(dc *DataConf) *gorm.MysqlConf {
	return dc.Mysql
}

func (ist *MysqlLoader) LoadLocal(dc *DataConf) error {
	filePath := getConfigFilePath("mysql.json")

	if err := loadJSONConfig(filePath, &dc.Mysql.Conf); err != nil {
		return err
	}

	return nil
}

func (ist *MysqlLoader) LoadRemote(dc *DataConf, bc *BootstrapConf, cc configCenter.ConfigCenterServiceClient) error {
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
