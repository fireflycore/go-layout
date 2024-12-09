package main

import (
	"fmt"
	"github.com/lhdhtrc/func-go/process"
	"go-layout/internal/conf"
	"runtime"
)

func main() {
	bootstrapConf := conf.NewBootstrapConf()

	app, cleanup, err := wireApp(bootstrapConf)
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}

	go app.Start()

	app.Logger.Info(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		cleanup()
	})
}
