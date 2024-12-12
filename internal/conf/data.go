package conf

import (
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	gorm "github.com/lhdhtrc/gorm/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	task "github.com/lhdhtrc/task-go/pkg"
)

type DataConf struct {
	Etcd  *etcd.Config
	Mysql *gorm.Config
}

func NewDataConf(bc *BootstrapConf, ist *task.Instance) *DataConf {
	result := &DataConf{
		Etcd:  &etcd.Config{},
		Mysql: &gorm.Config{},
	}
	micro.ReadConfigTask(ist, bc.DataConfFile, []interface{}{
		&result.Etcd,
		&result.Mysql,
	})
	ist.Await()
	micro.ReadCertAndWriteLocalTask(ist, "etcd", &result.Etcd.Tls)
	micro.ReadCertAndWriteLocalTask(ist, "mysql", &result.Mysql.Tls)
	ist.Await()

	return result
}
