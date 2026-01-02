package data

import (
	"context"
	"encoding/json"
	micro "github.com/lhdhtrc/micro-go/pkg/core"
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/internal/biz/repo"
	"time"
)

type loggerRepo struct {
	accessLogger    accessLogger.AccessLoggerServiceClient
	serverLogger    serverLogger.ServerLoggerServiceClient
	operationLogger operationLogger.OperationLoggerServiceClient
}

func NewLoggerRepo(
	serverLogger serverLogger.ServerLoggerServiceClient,
	accessLogger accessLogger.AccessLoggerServiceClient,
	operationLogger operationLogger.OperationLoggerServiceClient,
) repo.LoggerRepo {
	return &loggerRepo{
		serverLogger:    serverLogger,
		accessLogger:    accessLogger,
		operationLogger: operationLogger,
	}
}

func (repo *loggerRepo) CreateAccessLogger(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row accessLogger.CreateRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.AppId = appId

	_, _ = micro.WithRemoteInvoke[string, *accessLogger.CreateResponse](func() (*accessLogger.CreateResponse, error) {
		return repo.accessLogger.Create(ctx, &row)
	})
}

func (repo *loggerRepo) CreateServerLogger(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row serverLogger.CreateRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.AppId = appId

	_, _ = micro.WithRemoteInvoke[string, *serverLogger.CreateResponse](func() (*serverLogger.CreateResponse, error) {
		return repo.serverLogger.Create(ctx, &row)
	})
}

func (repo *loggerRepo) CreateOperationLogger(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row operationLogger.CreateRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.InvokeAppId = appId

	_, _ = micro.WithRemoteInvoke[string, *operationLogger.CreateResponse](func() (*operationLogger.CreateResponse, error) {
		return repo.operationLogger.Create(ctx, &row)
	})
}
