package main

import (
	"runtime"
	"time"

	"github.com/fireflycore/go-utils/process"
	"go.uber.org/zap"
)

func init() {
	time.Local = time.UTC
}

func main() {
	app, err := wireApp()
	if err != nil {
		panic(err)
	}

	go app.Start()

	app.Logger.Info("system run address", zap.String("address", app.BootstrapConf.Micro.Network.Internal))
	app.Logger.Info("system self check completed", zap.Int("goroutine_num", runtime.NumGoroutine()))
	process.Watcher(func() {
		app.Logger.Info("uninstall all service for this node from the register")
		app.Stop()
	})
}
