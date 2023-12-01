package config

import "microservice-go/utils/logger"

type Entrance struct {
	Logger logger.ConfigEntity
	Micro  _MicroConfigEntrance
	System _SystemConfigEntrance
}
