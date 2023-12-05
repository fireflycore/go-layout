package internal

import (
	"fmt"
	"gorm.io/gorm/logger"
	"microservice-go/store"
)

type GormWriter struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *GormWriter {
	return &GormWriter{Writer: w}
}

// Printf 格式化打印日志
func (w *GormWriter) Printf(message string, data ...interface{}) {
	store.Use.Logger.Func.Info(fmt.Sprintf(message+"\n", data...))
}
