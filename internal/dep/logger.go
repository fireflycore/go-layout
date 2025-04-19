package dep

import (
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/lhdhtrc/logger-go/pkg"
	accessLogger "go-layout/dep/protobuf/gen/acme/logger/access/v1"
	operationLogger "go-layout/dep/protobuf/gen/acme/logger/operation/v1"
	serverLogger "go-layout/dep/protobuf/gen/acme/logger/server/v1"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"go.uber.org/zap"
	"time"
)

func NewLogger(bc *conf.BootstrapConf, handle biz.ServerLogger) *zap.Logger {
	return logger.New(bc.Logger, handle)
}

func NewAccessLogger(bc *conf.BootstrapConf, service accessLogger.AccessLoggerServiceClient) biz.AccessLogger {
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

func NewServerLogger(bc *conf.BootstrapConf, service serverLogger.ServerLoggerServiceClient) biz.ServerLogger {
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

func NewOperationLogger(bc *conf.BootstrapConf, service operationLogger.OperationLoggerServiceClient) biz.OperationLogger {
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
