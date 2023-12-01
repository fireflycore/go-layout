package utils

import (
	"microservice-go/utils/mongo"
	"microservice-go/utils/time"
)

type _Entrance struct {
	Time  time.Entrance
	Mongo mongo.Entrance
}

var Use = new(_Entrance)
