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

func NewMysql(bc *conf.BootstrapConf, mc *gorme.MysqlConf, logger depend.OperationLogger) (*gorme.MysqlDB, error) {
	mc.Conf.WithAutoMigrate(false)
	mc.Conf.WithLoggerHandle(logger)
	mc.Conf.WithLoggerConsole(bc.Logger.Console)

	return gorme.NewMysql(mc, []interface{}{
		&entity.Demo{},
	})
}
