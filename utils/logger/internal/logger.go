package internal

import (
	"fmt"
	"time"
)

type ConfigEntity struct {
	Console bool
	Enable  bool
	Ltd     uint32
	Level   string
	Message string
	Remote  func(level uint32, message string)
}

func Logger(config *ConfigEntity) {
	if config.Console {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), config.Level, config.Message)
	}

	if config.Enable && config.Remote != nil {
		config.Remote(config.Ltd, config.Message)
	}
}
