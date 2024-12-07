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
		fmt.Println(err)
		return
	}

	go app.Start()

	fmt.Println(fmt.Sprintf("system self check completed，current goroutine num - %d", runtime.NumGoroutine()))
	process.Watcher(func() {
		cleanup()
	})
}
