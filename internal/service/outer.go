package service

import (
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/internal/biz"
)

func NewAccessLoggerOuterService(client biz.OuterGrpcClient) accessLogger.AccessLoggerServiceClient {
	return accessLogger.NewAccessLoggerServiceClient(client)
}

func NewOperationLoggerOuterService(client biz.OuterGrpcClient) operationLogger.OperationLoggerServiceClient {
	return operationLogger.NewOperationLoggerServiceClient(client)
}

func NewServerLoggerOuterService(client biz.OuterGrpcClient) serverLogger.ServerLoggerServiceClient {
	return serverLogger.NewServerLoggerServiceClient(client)
}
