package plugin

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	serverLogger "microservice-go/dep/protobuf/gen/acme/logger/server/v1"
	"microservice-go/model"
)

func InstallRemoteService(remote string) *model.RemoteServiceEntity {
	RemoteService := &model.RemoteServiceEntity{}
	if cli, err := grpc.NewClient(remote, grpc.WithTransportCredentials(insecure.NewCredentials())); err == nil {
		RemoteService.ServerLogger = serverLogger.NewServerLoggerServiceClient(cli)
	}
	return RemoteService
}
