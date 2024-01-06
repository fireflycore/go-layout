package service

import (
	"microservice-go/service/db"
	"microservice-go/service/logger"
)

type _Entrance struct {
	DB     db.EntranceEntity
	Logger logger.EntranceEntity
}

var Use = new(_Entrance)
