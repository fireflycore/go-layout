package micro

import (
	"fmt"
	"google.golang.org/grpc"
	"microservice-go/store"
	"microservice-go/utils"
	"net"
)

func RegisterServiceInstance() {
	logPrefix := "setup grpc"
	utils.Use.Logger.Info(fmt.Sprintf("%s %s %s", logPrefix, store.Use.Config.Local.Micro.Address.Inside, "start ->"))

	listen, err := net.Listen("tcp", store.Use.Config.Local.Micro.Address.Inside)
	if err != nil {
		utils.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
		return
	}
	server := grpc.NewServer()

	/*-------------------------------------Register Microservice---------------------------------*/
	/*-------------------------------------Register Microservice---------------------------------*/

	utils.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "register server done ->"))
	sErr := server.Serve(listen)
	if sErr != nil {
		utils.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, sErr.Error()))
		return
	}
}
