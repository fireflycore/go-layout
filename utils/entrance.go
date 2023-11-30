package utils

import (
	"microservice-go/utils/logger"
	"microservice-go/utils/mongo"
	"microservice-go/utils/time"
)

type _Entrance struct {
	Logger logger.Entrance
	Time   time.Entrance
	Mongo  mongo.Entrance
}

var Use = new(_Entrance)
