package core

import (
	"microservice-go/core/db"
)

type _Entrance struct {
	DB db.Entrance
}

var Use = new(_Entrance)
