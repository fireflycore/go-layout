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

func NewMysql(bc *conf.BootstrapConf, mc *gorme.Config, logger biz.OperationLogger) (*gorm.DB, error) {
	//mc.WithAutoMigrate(true)
	mc.WithLoggerHandle(logger)
	mc.WithLoggerConsole(bc.Logger.Console)

	return gorme.NewMysql(mc, []interface{}{
		&biz.Demo{},
	})
}
func NewData(db *gorm.DB) (*Data, error) {
	return &Data{
		db: db,
	}, nil
}
