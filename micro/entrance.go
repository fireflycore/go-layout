package micro

import (
	"microservice-go/micro/etcd"
	"microservice-go/micro/grpc"
)

type _Entrance struct {
	Etcd etcd.Entrance
	Grpc grpc.Entrance
}

var Use = new(_Entrance)
