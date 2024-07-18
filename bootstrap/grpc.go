package bootstrap

import (
	"fmt"
	"google.golang.org/grpc"
	"microservice-go/store"
	"net"
)

func GrpcServer(handle func(server *grpc.Server), address string) *grpc.Server {
	logPrefix := "install grpc server"
	store.Use.Logger.Info(fmt.Sprintf("%s %s %s", logPrefix, address, "start ->"))

	listen, err := net.Listen("tcp", address)
	if err != nil {
		store.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
		return nil
	}
	server := grpc.NewServer()

	/*-------------------------------------Register Microservice---------------------------------*/
	handle(server)
	/*-------------------------------------Register Microservice---------------------------------*/

	store.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "register server done ->"))
	go func() {
		sErr := server.Serve(listen)
		if sErr != nil {
			store.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, sErr.Error()))
			return
		}
	}()

	return server
}
