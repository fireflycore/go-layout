package main

import (
	"fmt"
	"github.com/lhdhtrc/func-go/process"
	"os"
	"runtime"
	"time"
)

func init() {
	// 1. 设置环境变量（影响系统组件）
	_ = os.Setenv("TZ", "UTC")

	// 2. 设置 Go 运行时（影响标准库）
	time.Local = time.UTC

	// 3. 重新加载确保一致性
	if loc, err := time.LoadLocation(""); err == nil {
		time.Local = loc
	}
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
