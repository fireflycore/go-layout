package config

import "microservice-go/model/config"

type _LoggerConfigEntrance struct {
	Mode   bool                   `json:"mode" yaml:"mode" mapstructure:"mode"`       // Mode is true enabling remote logging, is false console local(zap is available)
	Enable bool                   `json:"enable" yaml:"enable" mapstructure:"enable"` // Enable is logger
	Remote *config.DBConfigEntity `json:"remote" yaml:"remote" mapstructure:"remote"` // Remote database config
}
