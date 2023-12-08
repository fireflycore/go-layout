package common

import "microservice-go/model/base"

type DBConfigEntity struct {
	base.AccountPasswordEntity `mapstructure:"normal" yaml:"normal"`
	base.TlsEntity             `mapstructure:"tls" yaml:"tls"`

	Address string `json:"address" bson:"address" yaml:"address" mapstructure:"address"`
	DB      string `json:"db" bson:"db" yaml:"db" mapstructure:"db"`

	Mode bool `json:"mode" bson:"mode" yaml:"mode" mapstructure:"mode"` // Mode is true cluster
	Auth uint `json:"auth" bson:"auth" yaml:"auth" mapstructure:"auth"` // Auth way, account+password / TLS

	MaxOpenConnects        int  `json:"max_open_connect" bson:"max_open_connect" yaml:"max_open_connect" mapstructure:"max_open_connect"`
	MaxIdleConnects        int  `json:"max_idle_connect" bson:"max_idle_connect" yaml:"max_idle_connect" mapstructure:"max_idle_connect"`
	ConnMaxLifeTime        int  `json:"conn_max_life_time" bson:"conn_max_life_time" yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	SkipDefaultTransaction bool `json:"skip_default_transaction" bson:"skip_default_transaction" yaml:"skip_default_transaction" mapstructure:"skip_default_transaction"`
	PrepareStmt            bool `json:"prepare_stmt" bson:"prepare_stmt" yaml:"prepare_stmt" mapstructure:"prepare_stmt"`
}
