package conf

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewDataConf, NewEtcdConf, NewMysqlConf)
