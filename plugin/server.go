package plugin

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	serverLogger "microservice-go/dep/protobuf/gen/acme/logger/server/v1"
	"microservice-go/model"
)

func SetupRemoteService(remote string) *model.RemoteServiceEntity {
	RemoteService := &model.RemoteServiceEntity{}
	if cli, err := grpc.Dial(remote, grpc.WithTransportCredentials(insecure.NewCredentials())); err == nil {
		RemoteService.LoggerServer = serverLogger.NewServerLoggerServiceClient(cli)
	}
	return RemoteService
}
