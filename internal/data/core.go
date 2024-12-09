package data

import (
	"github.com/google/wire"
	etcd "github.com/lhdhtrc/etcd-go/pkg"
	gorme "github.com/lhdhtrc/gorm/pkg"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewDemoRepo)

type Data struct {
	Etcd  *clientv3.Client
	Mysql *gorm.DB
}

func NewMysql(console bool, mc *gorme.Config, logger biz.OperationLogger) (*gorm.DB, error) {
	//mc.WithAutoMigrate(true)
	mc.WithLoggerHandle(logger)
	mc.WithLoggerConsole(console)

	return gorme.NewMysql(mc, []interface{}{
		&biz.Demo{},
	})
}

func NewData(bc *conf.BootstrapConf, dc *conf.DataConf, logger biz.OperationLogger) (*Data, error) {
	result := &Data{}
	if cli, err := etcd.New(dc.Etcd); err != nil {
		result.Etcd = cli
		return nil, err
	}
	if orm, err := NewMysql(bc.Logger.Console, dc.Mysql, logger); err != nil {
		result.Mysql = orm
		return nil, err
	}

	return result, nil
}
