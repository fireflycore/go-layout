package data

import (
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"

	consul "github.com/fireflycore/go-consul"
	consulConfig "github.com/fireflycore/go-consul/config"
	microConfig "github.com/fireflycore/go-micro/config"
	redisx "github.com/fireflycore/go-redis"
	"github.com/fireflycore/gormx"
	"github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client

	store  microConfig.Store
	consul *api.Client
}

func NewData(db *gorm.DB, rdb *redis.Client, store microConfig.Store, consul *api.Client) (*Data, error) {
	return &Data{
		db:  db,
		rdb: rdb,

		store:  store,
		consul: consul,
	}, nil
}

func NewConsul(consulConf *consul.Config) (*api.Client, error) {
	return consul.New(consulConf)
}

func NewRedis(redisConf *redisx.Conf) (*redis.Client, error) {
	return redisx.New(redisConf)
}

// NewMysql 初始化模板库默认 MySQL 连接，并保持示例实体自动迁移关闭。
func NewMysql(bootstrapConf *conf.BootstrapConf, mysqlConf *gormx.MysqlConf) (*gormx.MysqlDB, error) {
	mysqlConf.WithLoggerConsole(bootstrapConf.Logger.Console)
	mysqlConf.WithAutoMigrate(false)

	return gormx.NewMysql(mysqlConf, []interface{}{
		&entity.Demo{},
	})
}

func NewConfigStore(client *api.Client) (microConfig.Store, error) {
	store, err := consulConfig.NewStore(client, nil)
	if err != nil {
		return nil, err
	}
	return store, nil
}
