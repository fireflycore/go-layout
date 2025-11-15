package depend

import (
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/lhdhtrc/logger-go/pkg"
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/internal/conf"
	"go.uber.org/zap"
	"time"
)

type AccessLogger func(b []byte, msg string)
type ServerLogger func(b []byte)
type OperationLogger func(b []byte)

func NewLogger(bc *conf.BootstrapConf, handle ServerLogger) *zap.Logger {
	return logger.New(bc.Logger, handle)
}

func NewAccessLogger(bc *conf.BootstrapConf, service accessLogger.AccessLoggerServiceClient) AccessLogger {
	async := logger.NewAsyncLogger(1000, func(b []byte) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		var row accessLogger.CreateRequest
		if err := json.Unmarshal(b, &row); err == nil {
			_, _ = service.Create(ctx, &row)
		} else {
			fmt.Println(err)
		}
	})

	return func(b []byte, msg string) {
		if bc.Logger.Console {
			fmt.Print(msg)
		}

		if !bc.Logger.Remote {
			return
		}

		async.Logger(b)
	}
}

func NewServerLogger(bc *conf.BootstrapConf, service serverLogger.ServerLoggerServiceClient) ServerLogger {
	async := logger.NewAsyncLogger(1000, func(b []byte) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		var row serverLogger.CreateRequest
		if err := json.Unmarshal(b, &row); err == nil {
			row.AppId = bc.AppId
			_, _ = service.Create(ctx, &row)
		} else {
			fmt.Println(err)
		}
	})

	return func(b []byte) {
		if !bc.Logger.Remote {
			return
		}
		async.Logger(b)
	}
}

func NewOperationLogger(bc *conf.BootstrapConf, service operationLogger.OperationLoggerServiceClient) OperationLogger {
	async := logger.NewAsyncLogger(1000, func(b []byte) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		var row operationLogger.CreateRequest
		if err := json.Unmarshal(b, &row); err == nil {
			row.InvokeAppId = bc.AppId
			_, _ = service.Create(ctx, &row)
		} else {
			fmt.Println(err)
		}
	})

	return func(b []byte) {
		if !bc.Logger.Remote {
			return
		}
		async.Logger(b)
	}
}
