package config

import (
	microModel "github.com/lhdhtrc/microcore-go/model"
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro/grpc"
)

type EntranceEntity struct {
	System SystemConfigEntity      `json:"system" bson:"system" yaml:"system" mapstructure:"system"`
	Grpc   grpc.ConfigEntity       `json:"grpc" bson:"grpc" yaml:"grpc" mapstructure:"grpc"`
	Logger logger.ConfigEntity     `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Micro  microModel.ConfigEntity `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
}
