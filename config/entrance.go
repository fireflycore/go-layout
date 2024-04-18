package config

import (
	loggerModel "github.com/lhdhtrc/logger-go/model"
	microModel "github.com/lhdhtrc/microcore-go/model"
	taskModel "github.com/lhdhtrc/task-go/model"
)

type EntranceEntity struct {
	System SystemConfigEntity       `json:"system" bson:"system" yaml:"system" mapstructure:"system"`
	Logger loggerModel.ConfigEntity `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Micro  microModel.ConfigEntity  `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
	Task   taskModel.ConfigEntity   `json:"task" bson:"task" yaml:"task" mapstructure:"task"`
}
