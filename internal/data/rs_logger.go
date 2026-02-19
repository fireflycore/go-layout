package data

import (
	"context"
	"encoding/json"
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/internal/biz/repo"
	"time"

	"github.com/fireflycore/go-micro/rpc"
)

type loggerRepo struct {
	ctx context.Context

	accessLogger    accessLogger.AccessLoggerServiceClient
	serverLogger    serverLogger.ServerLoggerServiceClient
	operationLogger operationLogger.OperationLoggerServiceClient
}

func NewLoggerRepo(
	ctx context.Context,

	serverLogger serverLogger.ServerLoggerServiceClient,
	accessLogger accessLogger.AccessLoggerServiceClient,
	operationLogger operationLogger.OperationLoggerServiceClient,
) repo.LoggerRepo {
	return &loggerRepo{
		ctx: ctx,

		serverLogger:    serverLogger,
		accessLogger:    accessLogger,
		operationLogger: operationLogger,
	}
}

func (ur *loggerRepo) CreateAccessLog(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(ur.ctx, time.Second*5)
	defer cancel()

	var row accessLogger.CreateLogRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.AppId = appId

	_, _ = rpc.WithRemoteInvoke[string, *accessLogger.CreateLogResponse](func() (*accessLogger.CreateLogResponse, error) {
		return ur.accessLogger.CreateLog(ctx, &row)
	})
}

func (ur *loggerRepo) CreateServerLog(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(ur.ctx, time.Second*5)
	defer cancel()

	var row serverLogger.CreateLogRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.AppId = appId

	_, _ = rpc.WithRemoteInvoke[string, *serverLogger.CreateLogResponse](func() (*serverLogger.CreateLogResponse, error) {
		return ur.serverLogger.CreateLog(ctx, &row)
	})
}

func (ur *loggerRepo) CreateOperationLog(appId string, raw []byte) {
	ctx, cancel := context.WithTimeout(ur.ctx, time.Second*5)
	defer cancel()

	var row operationLogger.CreateLogRequest
	if err := json.Unmarshal(raw, &row); err != nil {
		return
	}
	row.TargetAppId = appId

	_, _ = rpc.WithRemoteInvoke[string, *operationLogger.CreateLogResponse](func() (*operationLogger.CreateLogResponse, error) {
		return ur.operationLogger.CreateLog(ctx, &row)
	})
}
