package data

import (
	gorme "github.com/lhdhtrc/gorm/pkg"
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"
	"go-layout/internal/depend"
	"gorm.io/gorm"
)

type Data struct {
	etcd *clientv3.Client
	rdb  *redis.Client
	db   *gorm.DB
}

func NewData(etcd *clientv3.Client, rdb *redis.Client, db *gorme.MysqlDB) (*Data, error) {
	return &Data{
		etcd: etcd,
		rdb:  rdb,
		db:   db.DB,
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
