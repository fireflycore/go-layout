package conf

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	gorm "github.com/lhdhtrc/gorm/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
)

type DataConf struct {
	Etcd  *etcd.Config
	Mysql *gorm.Config
}

func NewDataConf(ist *task.Instance) *DataConf {
	result := &DataConf{
		Etcd:  &etcd.Config{},
		Mysql: &gorm.Config{},
	}
	micro.InitConfigTask(ist, []string{
		"https://demo.com/config/etcd.config.json",
		"https://demo.com/config/mysql.config.json",
	}, []interface{}{
		&result.Etcd,
		&result.Mysql,
	})
	ist.Await()
	micro.InitCertTask(ist, "etcd", &result.Etcd.Tls)
	micro.InitCertTask(ist, "mysql", &result.Mysql.Tls)
	ist.Await()

	return result
}
