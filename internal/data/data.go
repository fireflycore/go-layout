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

// NewData 汇总数据层基础依赖，保持与 config 服务一致的集中注入方式。
func NewData(db *gorm.DB, rdb *redis.Client, store microConfig.Store, consul *api.Client) *Data {
	return &Data{
		db:  db,
		rdb: rdb,

		store:  store,
		consul: consul,
	}
}

// NewConsul 创建数据层复用的 Consul 客户端，供 Store 与运维场景统一复用。
func NewConsul(consulConf *consul.Config) (*api.Client, error) {
	return consul.New(consulConf)
}

// NewRedis 创建模板默认 Redis 客户端，供 Repo 与缓存场景直接注入使用。
func NewRedis(redisConf *redisx.Conf) (*redis.Client, error) {
	return redisx.New(redisConf)
}

// NewMysql 初始化模板库默认 MySQL 连接，并保持示例实体自动迁移关闭。
func NewMysql(bootstrapConf *conf.BootstrapConf, mysqlConf *gormx.MysqlConf) (*gorm.DB, error) {
	mysqlConf.WithLoggerConsole(bootstrapConf.Logger.Console)
	mysqlConf.WithAutoMigrate(false)

	db, err := gormx.NewMysql(mysqlConf, []any{
		&entity.Demo{},
	})

	if err != nil {
		return nil, err
	}

	return db.DB, nil
}

// NewConfigStore 基于 Consul 客户端构造运行期配置 Store。
func NewConfigStore(client *api.Client) (microConfig.Store, error) {
	store, err := consulConfig.NewStore(client, nil)
	if err != nil {
		return nil, err
	}
	return store, nil
}
