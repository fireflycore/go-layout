package store

import (
	"github.com/lhdhtrc/microservice-go/logger"
	"github.com/lhdhtrc/microservice-go/micro"
)

type _Entrance struct {
	Logger logger.Abstraction
	Micro  micro.Abstraction
}

var Use = new(_Entrance)
