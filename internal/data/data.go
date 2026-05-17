package data

import (
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"

	consul "github.com/fireflycore/go-consul"
	consulConfig "github.com/fireflycore/go-consul/config"
	microConfig "github.com/fireflycore/go-micro/config"
	"github.com/fireflycore/go-micro/constant"
	redisx "github.com/fireflycore/go-redis"
	"github.com/fireflycore/gormx"
	"github.com/fireflycore/gormx/logger"
	"github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Data 汇总 data 层共享基础依赖。
type Data struct {
	db  *gorm.DB
	rdb *redis.Client

	store  microConfig.Store
	consul *api.Client
}

// NewData 汇总模板服务数据层基础依赖，统一向 Repo 暴露数据库、缓存与 Store 能力。
func NewData(db *gorm.DB, rdb *redis.Client, store microConfig.Store, consul *api.Client) *Data {
	return &Data{
		db:  db,
		rdb: rdb,

		store:  store,
		consul: consul,
	}
}

// NewMysql 初始化模板服务使用的 MySQL 连接，并保持示例实体自动迁移关闭。
func NewMysql(bootstrapConfig *conf.BootstrapConfig, mysqlConfig *gormx.MysqlConfig) (*gorm.DB, error) {
	mysqlConfig.WithTables([]any{
		&entity.Demo{},
	})
	mysqlConfig.WithAutoMigrate(false)
	mysqlConfig.WithLoggerConsole(bootstrapConfig.Logger.Console)
	mysqlConfig.WithUserContextFields(&logger.UserContextFields{
		UserId:  constant.UserId,
		OrgIds:  constant.OrgIds,
		RoleIds: constant.RoleIds,

		AppId:    constant.AppId,
		TenantId: constant.TenantId,

		ServiceAppId:      constant.ServiceAppId,
		ServiceInstanceId: constant.ServiceInstanceId,
	})

	db, err := gormx.NewMysql(mysqlConfig)

	if err != nil {
		return nil, err
	}

	return db.DB, nil
}

// NewRedis 创建模板服务默认 Redis 客户端。
func NewRedis(redisConf *redisx.Config) (*redis.Client, error) {
	return redisx.New(redisConf)
}

// NewConfigStore 基于 Consul 客户端构造统一配置 Store，供运行期配置按当前主线直接从数据面读取。
func NewConfigStore(client *api.Client) (microConfig.Store, error) {
	store, err := consulConfig.NewStore(client, nil)
	if err != nil {
		return nil, err
	}
	return store, nil
}

// NewConsul 创建模板服务复用的 Consul 客户端。
func NewConsul(consulConf *consul.Config) (*api.Client, error) {
	return consul.New(consulConf)
}
