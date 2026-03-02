package dep

import (
	"go-layout/internal/biz/repo"
	"go-layout/internal/conf"
	"os"

	"github.com/fireflycore/go-micro/logger"
)

const CacheSize = 1000

type AccessLogger func(b []byte, msg string)
type ServerLogger func(b []byte)
type OperationLogger func(b []byte)

func NewLogger(bootstrapConf *conf.BootstrapConf, handle ServerLogger) *logger.Core {
	return logger.NewLogger(logger.NewZapLogger(bootstrapConf.Logger, handle))
}

func NewAccessLogger(bootstrapConf *conf.BootstrapConf, loggerRepo repo.LoggerRepo) AccessLogger {
	async := logger.NewAsyncLogger(CacheSize, func(b []byte) {
		loggerRepo.CreateAccessLog(bootstrapConf.AppId, b)
	})

	return func(b []byte, msg string) {
		if bootstrapConf.Logger.Console {
			_, _ = os.Stdout.WriteString(msg)
		}

		if !bootstrapConf.Logger.Remote {
			return
		}

		async.Logger(b)
	}
}

func NewServerLogger(bootstrapConf *conf.BootstrapConf, loggerRepo repo.LoggerRepo) ServerLogger {
	async := logger.NewAsyncLogger(CacheSize, func(b []byte) {
		loggerRepo.CreateServerLog(bootstrapConf.AppId, b)
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
		loggerRepo.CreateOperationLog(bootstrapConf.AppId, b)
	})

	return func(b []byte) {
		if !bootstrapConf.Logger.Remote {
			return
		}
		async.Logger(b)
	}
}
