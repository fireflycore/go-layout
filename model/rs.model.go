package model

import serverLogger "microservice-go/dep/protobuf/gen/acme/logger/server/v1"

type RemoteServiceEntity struct {
	LoggerServer serverLogger.ServerLoggerServiceClient
}
