package model

import serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"

type RemoteServiceEntity struct {
	ServerLogger serverLogger.ServerLoggerServiceClient
}
