package register

import (
	"fmt"
	"google.golang.org/grpc"
	"microservice-go/store"
	"net"
)

func ServiceInstance() {
	logPrefix := "setup grpc"
	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s %s", logPrefix, store.Use.Config.Local.Micro.Address.Inside, "start ->"))

	listen, err := net.Listen("tcp", store.Use.Config.Local.Micro.Address.Inside)
	if err != nil {
		store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
		return
	}
	server := grpc.NewServer()

	/*-------------------------------------Register Microservice---------------------------------*/
	/*-------------------------------------Register Microservice---------------------------------*/

	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s", logPrefix, "register server done ->"))
	go func() {
		sErr := server.Serve(listen)
		if sErr != nil {
			store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, sErr.Error()))
			return
		}
	}()
}
