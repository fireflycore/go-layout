package service

import (
	"microservice-go/service/logger"
)

type _Entrance struct {
	Logger logger.EntranceEntity
}

var Use = new(_Entrance)
