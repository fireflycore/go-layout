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
	etcdCli, ee := etcd.New(dc.Etcd)
	if ee != nil {
		return nil, ee
	}

	mysqlCli, oe := NewMysql(bc.Logger.Console, dc.Mysql, logger)
	if oe != nil {
		return nil, oe
	}

	return &Data{
		Etcd:  etcdCli,
		Mysql: mysqlCli,
	}, nil
}
