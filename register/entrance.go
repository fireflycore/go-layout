package register

import (
	"google.golang.org/grpc"
	"microservice-go/store"
)

func ServiceInstance() {
	store.Use.Grpc.Server(func(server *grpc.Server) {
		// register microservice
	}, store.Use.Config.Grpc.Address)
}
