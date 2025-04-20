package data

import (
	"github.com/google/wire"
	gorme "github.com/lhdhtrc/gorm/pkg"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewMysql, NewData, NewDemoRepo)

type Data struct {
	db *gorm.DB
}

func NewMysql(bc *conf.BootstrapConf, mc *gorme.MysqlConf, logger biz.OperationLogger) (*gorme.MysqlDB, error) {
	//mc.Conf.WithAutoMigrate(true)
	mc.Conf.WithLoggerHandle(logger)
	mc.Conf.WithLoggerConsole(bc.Logger.Console)

	return gorme.NewMysql(mc, []interface{}{
		&biz.Demo{},
	})
}

func NewData(mysql *gorme.MysqlDB) (*Data, error) {
	return &Data{
		db: mysql.DB,
	}, nil
}
