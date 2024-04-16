package config

import (
	loggerModel "github.com/lhdhtrc/logger-go/model"
	"github.com/lhdhtrc/microcore-go/grpc"
	microModel "github.com/lhdhtrc/microcore-go/model"
)

type EntranceEntity struct {
	System SystemConfigEntity       `json:"system" bson:"system" yaml:"system" mapstructure:"system"`
	Grpc   grpc.ConfigEntity        `json:"grpc" bson:"grpc" yaml:"grpc" mapstructure:"grpc"`
	Logger loggerModel.ConfigEntity `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Micro  microModel.ConfigEntity  `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
}
