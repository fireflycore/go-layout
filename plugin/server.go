package plugin

import (
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InstallRemoteService(remote string) *model.RemoteServiceEntity {
	RemoteService := &model.RemoteServiceEntity{}
	if cli, err := grpc.NewClient(remote, grpc.WithTransportCredentials(insecure.NewCredentials())); err == nil {
		RemoteService.ServerLogger = serverLogger.NewServerLoggerServiceClient(cli)
	}
	return RemoteService
}
