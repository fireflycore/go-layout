package conf

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	gorm "github.com/lhdhtrc/gorm/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
)

type DataConf struct {
	Etcd  *etcd.Config
	Mysql *gorm.Config
}

func NewDataConf(bc *BootstrapConf, ist *task.Instance) *DataConf {
	result := new(DataConf)
	return result
}

func NewMysqlConf(dc *DataConf) *gorm.Config {
	return dc.Mysql
}

func NewEtcdConf(dc *DataConf) *etcd.Config {
	return dc.Etcd
}
