package config

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
)

type Entrance struct {
	Logger logger.ConfigEntity `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Micro  micro.ConfigEntity  `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
}
