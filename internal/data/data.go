package data

import (
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"
	"go-layout/internal/dep"

	etcd "github.com/fireflycore/go-etcd"
	redisx "github.com/fireflycore/go-redis"
	gorme "github.com/fireflycore/gormx"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
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

func NewRedis(redisConf *redisx.Conf) (*redis.Client, error) {
	return redisx.New(redisConf)
}

func NewMysql(bootstrapConf *conf.BootstrapConf, mysqlConf *gorme.MysqlConf, logger dep.OperationLogger) (*gorme.MysqlDB, error) {
	mysqlConf.WithAutoMigrate(false)
	mysqlConf.WithLoggerHandle(logger)
	mysqlConf.WithLoggerConsole(bootstrapConf.Logger.Console)

	return gorme.NewMysql(mysqlConf, []interface{}{
		&entity.Demo{},
	})
}
