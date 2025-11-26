package dep

import (
	"fmt"
	logger "github.com/lhdhtrc/logger-go/pkg"
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"
	"go.uber.org/zap"
)

const CacheSize = 1000

type AccessLogger func(b []byte, msg string)
type ServerLogger func(b []byte)
type OperationLogger func(b []byte)

func NewLogger(bootstrapConf *conf.BootstrapConf, handle ServerLogger) *zap.Logger {
	return logger.New(bootstrapConf.Logger, handle)
}

func NewAccessLogger(bootstrapConf *conf.BootstrapConf, loggerRepo repo.LoggerRepo) AccessLogger {
	async := logger.NewAsyncLogger(CacheSize, func(b []byte) {
		loggerRepo.CreateAccessLogger(bootstrapConf.AppId, b)
	})

	return func(b []byte, msg string) {
		if bootstrapConf.Logger.Console {
			fmt.Print(msg)
		}

		if !bootstrapConf.Logger.Remote {
			return
		}

		async.Logger(b)
	}
}

func NewServerLogger(bootstrapConf *conf.BootstrapConf, loggerRepo repo.LoggerRepo) ServerLogger {
	async := logger.NewAsyncLogger(CacheSize, func(b []byte) {
		loggerRepo.CreateServerLogger(bootstrapConf.AppId, b)
	})

	return func(b []byte) {
		if !bootstrapConf.Logger.Remote {
			return
		}
		async.Logger(b)
	}
}

func NewOperationLogger(bootstrapConf *conf.BootstrapConf, loggerRepo repo.LoggerRepo) OperationLogger {
	async := logger.NewAsyncLogger(CacheSize, func(b []byte) {
		loggerRepo.CreateOperationLogger(bootstrapConf.AppId, b)
	})

	return func(b []byte) {
		if !bootstrapConf.Logger.Remote {
			return
		}
		async.Logger(b)
	}
}
