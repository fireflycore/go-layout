package data

import (
	gorme "github.com/lhdhtrc/gorm/pkg"
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"
	"go-layout/internal/depend"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

func NewData(mysql *gorme.MysqlDB) (*Data, error) {
	return &Data{
		db: mysql.DB,
	}, nil
}

func NewEtcd(etcdConf *etcd.Conf) (*clientv3.Client, error) {
	return etcd.New(etcdConf)
}

func NewRedis(redisConf *redise.Conf) (*redis.Client, error) {
	return redise.New(redisConf)
}

func NewMysql(bootstrapConf *conf.BootstrapConf, mysqlConf *gorme.MysqlConf, logger dep.OperationLogger) (*gorme.MysqlDB, error) {
	mysqlConf.WithAutoMigrate(false)
	mysqlConf.WithLoggerHandle(logger)
	mysqlConf.WithLoggerConsole(bootstrapConf.Logger.Console)

	return gorme.NewMysql(mysqlConf, []interface{}{
		&entity.Demo{},
	})
}
