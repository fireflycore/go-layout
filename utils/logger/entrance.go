package logger

import (
	"microservice-go/utils/logger/internal"
)

type Interface interface {
	Info(message string)
	Success(message string)
	Warning(message string)
	Error(message string)
}

func New(config ConfigEntity) Interface {
	if len(config.Dictionary) != 4 {
		config.Dictionary = []uint32{10, 11, 12, 13}
	}
	return &entrance{config}
}

type ConfigEntity struct {
	Enable     bool     `json:"enable" bson:"enable" yaml:"enable" mapstructure:"enable"`                 // Enable is remote
	Console    bool     `json:"console" bson:"console" yaml:"console" mapstructure:"console"`             // Console is console
	Dictionary []uint32 `json:"dictionary" bson:"dictionary" yaml:"dictionary" mapstructure:"dictionary"` // Dictionary remote storage dictionary code
	WithRemote func(level uint32, message string)
}

type entrance struct {
	Config ConfigEntity
}

func (c *entrance) Info(message string) {
	internal.Logger(&internal.ConfigEntity{
		Console: c.Config.Console,
		Enable:  c.Config.Enable,
		Ltd:     c.Config.Dictionary[0],
		Level:   "info",
		Message: message,
		Remote:  c.Config.WithRemote,
	})
}
func (c *entrance) Success(message string) {
	internal.Logger(&internal.ConfigEntity{
		Console: c.Config.Console,
		Enable:  c.Config.Enable,
		Ltd:     c.Config.Dictionary[1],
		Level:   "success",
		Message: message,
		Remote:  c.Config.WithRemote,
	})
}
func (c *entrance) Warning(message string) {
	internal.Logger(&internal.ConfigEntity{
		Console: c.Config.Console,
		Enable:  c.Config.Enable,
		Ltd:     c.Config.Dictionary[2],
		Level:   "warning",
		Message: message,
		Remote:  c.Config.WithRemote,
	})
}
func (c *entrance) Error(message string) {
	internal.Logger(&internal.ConfigEntity{
		Console: c.Config.Console,
		Enable:  c.Config.Enable,
		Ltd:     c.Config.Dictionary[3],
		Level:   "error",
		Message: message,
		Remote:  c.Config.WithRemote,
	})
}
