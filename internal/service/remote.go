package service

import (
	config "go-layout/dep/protobuf/gen/acme/config/v1"
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"google.golang.org/grpc"
)

func NewConfigCenterRemoteService(client *grpc.ClientConn) config.ConfigServiceClient {
	return config.NewConfigServiceClient(client)
}

func NewAccessLoggerRemoteService(client *grpc.ClientConn) accessLogger.AccessLoggerServiceClient {
	return accessLogger.NewAccessLoggerServiceClient(client)
}

func NewOperationLoggerRemoteService(client *grpc.ClientConn) operationLogger.OperationLoggerServiceClient {
	return operationLogger.NewOperationLoggerServiceClient(client)
}

func NewServerLoggerRemoteService(client *grpc.ClientConn) serverLogger.ServerLoggerServiceClient {
	return serverLogger.NewServerLoggerServiceClient(client)
}
