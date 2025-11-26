package main

import (
	"fmt"
	"github.com/lhdhtrc/func-go/process"
	"runtime"
	"time"
)

func init() {
	time.Local = time.UTC
}

func main() {
	app, err := wireApp()
	if err != nil {
		panic(err)
		return
	}

	go app.Start()

	app.Logger.Info(fmt.Sprintf("system run address - %s", app.BootstrapConf.Micro.Network.Internal))
	app.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		app.Logger.Info("uninstall all service for this node from the register")
	})
}
