package utils

import (
	"microservice-go/utils/logger"
)

type _Entrance struct {
	Logger logger.Entrance
}

var Use = new(_Entrance)
