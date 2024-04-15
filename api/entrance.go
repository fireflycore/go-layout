package api

import (
	"google.golang.org/grpc"
	"microservice-go/store"
)

func ServiceInstance() {
	store.Use.Grpc.CreateServer(func(server *grpc.Server) {
		// register microservice
	}, store.Use.Config.System.RunPort)
}
