package config

import (
	loger "github.com/lhdhtrc/logger-go/pkg"
	micro "github.com/lhdhtrc/micro-go/pkg"
	task "github.com/lhdhtrc/task-go/pkg"
)

type CoreEntity struct {
	System SystemConfigEntity `json:"system" bson:"system" yaml:"system" mapstructure:"system"`
	Logger loger.ConfigEntity `json:"logger" bson:"logger" yaml:"logger" mapstructure:"logger"`
	Task   task.ConfigEntity  `json:"task" bson:"task" yaml:"task" mapstructure:"task"`
	Micro  micro.ConfigEntity `json:"micro" bson:"micro" yaml:"micro" mapstructure:"micro"`
}
