package service

import "microservice-go/service/logger"

type _Entrance struct {
	Logger logger.Entrance
}

var Use = new(_Entrance)
