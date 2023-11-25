package core

import (
	"microservice-go/core/db"
	"microservice-go/core/micro"
)

type _Entrance struct {
	DB    db.Entrance
	Micro micro.Entrance
}

var Use = new(_Entrance)
