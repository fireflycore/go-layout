package data

import (
	"fmt"
	"go-layout/internal/conf"
	"go-layout/internal/data/entity"

	"github.com/fireflycore/go-consul"
	consulConfig "github.com/fireflycore/go-consul/config"
	"github.com/fireflycore/go-micro/config"
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

	store  config.Store
	consul *api.Client
}

// NewData 统一向各 Repo 暴露数据库、缓存与配置读取能力。
func NewData(db *gorm.DB, rdb *redis.Client, store config.Store, consul *api.Client) *Data {
	return &Data{
		db:     db,
		rdb:    rdb,
		store:  store,
		consul: consul,
	}
}

// NewMysql 初始化当前服务使用的 Mysql 连接，并保持自动迁移关闭。
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

// NewRedis 创建当前服务使用的 Redis 客户端。
func NewRedis(redisConfig *redisx.Config) (*redis.Client, error) {
	return redisx.New(redisConfig)
}

// NewConfigStore 基于 Consul 客户端构造统一配置 Store。
func NewConfigStore(client *api.Client, bootstrapConfig *conf.BootstrapConfig) (config.Store, error) {
	store, err := consulConfig.NewStore(client, nil, config.WithNamespace(fmt.Sprintf("%s/config", bootstrapConfig.Service.Namespace)))
	if err != nil {
		return nil, err
	}
	return store, nil
}

// NewConfigClient 基于 go-consul/config.Client 构造运行时配置客户端。
func NewConfigClient(store config.Store) config.Client {
	consulStore, ok := store.(*consulConfig.StoreInstance)
	if !ok || consulStore == nil {
		return nil
	}

	// 构造失败时返回 nil，让上层退回到直接走 Store.Get 的兜底路径。
	client, err := consulConfig.NewClient(consulStore)
	if err != nil {
		return nil
	}

	return client
}

// NewConsul 创建当前服务复用的 Consul 客户端。
func NewConsul(consulConfig *consul.Config) (*api.Client, error) {
	return consul.New(consulConfig)
}
