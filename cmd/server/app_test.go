package main

import (
	"context"
	"errors"
	"go-layout/internal/conf"
	"go-layout/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fireflycore/go-consul/agent"
	microapp "github.com/fireflycore/go-micro/app"
	"github.com/fireflycore/go-micro/kernel"
	"github.com/fireflycore/go-micro/logger"
	"github.com/fireflycore/go-micro/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

// 使用 Agent 和 HTTP 客户端复现关闭阶段，确认 app 的退出策略与本地清理结果。
func TestAppRunShutdownErrors(t *testing.T) {
	tests := []struct {
		name          string
		cancel        bool
		sidecarOnline bool
		shutdownErr   error
		wantErr       bool
		wantWarning   bool
	}{
		{name: "canceled with unavailable sidecar", cancel: true, wantWarning: true},
		{name: "ordinary exit with unavailable sidecar", wantErr: true},
		{name: "canceled with rejected drain", cancel: true, sidecarOnline: true, wantErr: true},
		{name: "canceled with local shutdown failure", cancel: true, shutdownErr: errors.New("local shutdown failed"), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sidecar := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte(`{"success":false,"code":"drain_failed","message":"rejected"}`))
			}))
			t.Cleanup(sidecar.Close)
			if !tt.sidecarOnline {
				sidecar.Close()
			}
			node := &agent.ServiceNode{
				ServiceOptions: &agent.ServiceOptions{
					App: microapp.Config{
						Id: "app", InstanceId: "test-instance", Name: "app", Env: "test", Version: "1.0.0",
					},
					Kernel:     kernel.Config{Language: "go", Version: "1.25.12"},
					Service:    service.Config{Name: "app", Namespace: "test", Weight: 100},
					Protocol:   "grpc",
					ServerPort: 10503,
				},
				DNS: "app.test.svc.cluster.local",
			}
			client := agent.NewApiClient(agent.NewHttpClient(sidecar.URL, 100*time.Millisecond))
			controller, err := agent.NewController(client, node)
			if err != nil {
				t.Fatal(err)
			}
			runner, err := agent.NewRunner(idleSidecarEvents{}, controller, nil)
			if err != nil {
				t.Fatal(err)
			}
			sidecarAgent := &agent.Agent{Controller: controller, Runner: runner}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			shutdownCalled := false
			sidecarAgent.ConfigureRun(agent.SidecarAgentConfig{
				GracePeriod: "20s",
				Serve: func(context.Context) error {
					if tt.cancel {
						cancel()
					}
					return nil
				},
				Shutdown: func(context.Context) error {
					shutdownCalled = true
					return tt.shutdownErr
				},
			})
			core, logs := observer.New(zap.WarnLevel)
			app := &App{
				BootstrapConfig: &conf.BootstrapConfig{},
				Server:          &server.AppServer{SidecarAgent: sidecarAgent},
				Logger:          logger.NewServerLogger(zap.New(core)),
			}
			err = app.Run(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Run() error = %v, want error: %v", err, tt.wantErr)
			}
			if !shutdownCalled {
				t.Fatal("local shutdown callback was not called")
			}
			if tt.shutdownErr != nil && !errors.Is(err, tt.shutdownErr) {
				t.Fatalf("local shutdown error was lost: %v", err)
			}
			warnings := logs.FilterMessage("failed to notify sidecar during service shutdown").All()
			if (len(warnings) == 1) != tt.wantWarning {
				t.Fatalf("shutdown warnings = %d, want warning: %v", len(warnings), tt.wantWarning)
			}
			if tt.wantWarning && warnings[0].ContextMap()["error"] == nil {
				t.Fatal("shutdown warning omitted the original error")
			}
		})
	}
}

type idleSidecarEvents struct{}

func (idleSidecarEvents) Subscribe(context.Context) (<-chan agent.ConnectionEvent, error) {
	return make(chan agent.ConnectionEvent), nil
}
