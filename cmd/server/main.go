package main

import (
	"context"
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

	// 创建应用根上下文，统一接收系统退出信号并驱动托管运行结束。
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听系统信号，收到退出信号时只需取消上下文，由托管器统一完成收尾。
	go process.Watcher(func() {
		app.Logger.Info("received shutdown signal")
		cancel()
	})

	// 启动前输出当前 goroutine 数，便于基础运行态观测。
	app.Logger.Info("go-layout service runtime prepared",
		zap.Int("goroutine_num", runtime.NumGoroutine()),
	)

	// 进入统一托管运行入口，由 App.Run 负责 sidecar 生命周期和本地服务协同运行。
	if err = app.Run(ctx); err != nil && err != context.Canceled {
		app.Logger.Error("go-layout service exited with error", zap.Error(err))
		panic(err)
	}
}
