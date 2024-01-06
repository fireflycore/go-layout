package config

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
)

type EntranceEntity struct {
	System SystemConfigEntity  `json:"system" bson:"system" yaml:"system" mapstructure:"system"`
	Grpc   grpc.ConfigEntity   `json:"grpc" bson:"grpc" yaml:"grpc" mapstructure:"grpc"`
	Logger logger.ConfigEntity `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Micro  micro.ConfigEntity  `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
}
