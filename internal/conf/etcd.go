package conf

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	config "go-layout/depend/protobuf/gen/acme/config/v1"
)

// EtcdLoader 实现 ConfigLoader 接口
type EtcdLoader struct{}

func NewEtcdConf(dc *Conf) *etcd.Config {
	return dc.Etcd
}

func (ist *EtcdLoader) Local(dc *Conf) error {
	filePath := getConfigFilePath("etcd.json")

	if err := loadJSONConfig(filePath, &dc.Etcd); err != nil {
		return err
	}

	return nil
}

func (ist *EtcdLoader) Remote(dc *Conf, bc *BootstrapConf, cc config.ConfigServiceClient) error {
	content, err := fetchRemoteConfig(bc, cc, "database", "etcd")

	if err != nil {
		return err
	}

	if err = analyzeData(content, []byte(bc.AppSecret), &dc.Etcd); err != nil {
		return err
	}

	if err = analyzeTlsData("etcd", dc.Etcd.Tls); err != nil {
		return err
	}

	return nil
}
