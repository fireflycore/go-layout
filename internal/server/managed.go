package server

import (
	"context"
	"encoding/json"
	"errors"
	"go-layout/internal/conf"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/fireflycore/go-consul/agent"
	"github.com/fireflycore/go-micro/telemetry"
)

// SidecarStatusProvider 抽象 sidecar 生命周期状态读取能力。
type SidecarStatusProvider interface {
	Status() agent.Status
}

// AppManagedServer 表示服务管理端口 HTTP 服务器。
type AppManagedServer struct {
	lister net.Listener
	srv    *http.Server
}

// NewAppManagedServer 创建管理端口，并把 sidecar 生命周期状态暴露到 ready/info 端点。
func NewAppManagedServer(bootstrapConfig *conf.BootstrapConfig, provider *telemetry.Providers, sidecarStatusProvider SidecarStatusProvider) *AppManagedServer {
	l, e := net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConfig.ManagedPort), 10)))
	if e != nil {
		panic(e)
	}

	mux := http.NewServeMux()

	if provider != nil && provider.MetricsHandler != nil {
		mux.Handle("/metrics", provider.MetricsHandler)
	}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		// ready 端点除了返回服务就绪，还附带当前 sidecar 接管状态摘要。
		writeJSON(w, http.StatusOK, map[string]any{
			"status":  "ready",
			"sidecar": buildSidecarStatus(bootstrapConfig, sidecarStatusProvider),
		})
	})
	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		// info 端点汇总服务元信息、telemetry 配置和 sidecar 运行快照。
		writeJSON(w, http.StatusOK, map[string]any{
			"service_name":        bootstrapConfig.Service.Name,
			"service_namespace":   bootstrapConfig.Service.Namespace,
			"service_instance_id": bootstrapConfig.App.InstanceId,
			"service_endpoint":    net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConfig.ServerPort), 10)),
			"management_endpoint": net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConfig.ManagedPort), 10)),
			"app_id":              bootstrapConfig.App.Id,
			"app_name":            bootstrapConfig.App.Name,
			"version":             bootstrapConfig.App.Version,
			"build_info":          buildInfo(),
			"telemetry": map[string]any{
				"endpoint": bootstrapConfig.Telemetry.OTLPEndpoint,
				"insecure": bootstrapConfig.Telemetry.Insecure,
				"traces":   bootstrapConfig.Telemetry.Traces,
				"metrics":  bootstrapConfig.Telemetry.Metrics,
				"logs":     bootstrapConfig.Telemetry.Logs,
			},
			"sidecar": buildSidecarStatus(bootstrapConfig, sidecarStatusProvider),
		})
	})

	return &AppManagedServer{
		lister: l,
		srv: &http.Server{
			Handler: mux,
		},
	}
}

// Addr 返回管理端口实际监听地址。
func (ist *AppManagedServer) Addr() string {
	if ist == nil || ist.lister == nil {
		return ""
	}
	return ist.lister.Addr().String()
}

// Serve 启动管理端口 HTTP 服务。
func (ist *AppManagedServer) Serve() error {
	err := ist.srv.Serve(ist.lister)
	// 主动关闭管理端口时的 ErrServerClosed 不视为异常退出。
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

// Stop 关闭管理端口 HTTP 服务。
func (ist *AppManagedServer) Stop() {
	// 关闭阶段给管理端口一个固定宽限期，避免请求被立即切断。
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ist.srv.Shutdown(ctx); err != nil {
		// 优雅关闭失败时直接关闭 listener，避免端口悬挂。
		_ = ist.lister.Close()
	}
}

// writeJSON 统一输出 JSON 响应。
func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// buildSidecarStatus 把生命周期状态快照转换成管理接口使用的稳定 JSON 结构。
func buildSidecarStatus(bootstrapConfig *conf.BootstrapConfig, provider SidecarStatusProvider) map[string]any {
	status := agent.Status{}
	if provider != nil {
		// 若已接入 sidecar 生命周期，则直接复用其最新状态快照。
		status = provider.Status()
	}
	return map[string]any{
		// 输出当前 sidecar-agent 基础配置，便于现场排障。
		"base_url":     bootstrapConfig.SidecarAgent.BaseURL,
		"grace_period": bootstrapConfig.SidecarAgent.GracePeriod,
		// 输出连接和注册主状态，便于快速判断接管是否成功。
		"connected":  status.Connected,
		"registered": status.Registered,
		// 输出最近一次成功注册的服务信息。
		"last_service_name": status.LastServiceName,
		"last_service_port": status.LastServicePort,
		// 输出最近事件、连接时间和断连时间，便于恢复时序排查。
		"last_event_type":      status.LastEventType,
		"last_event_id":        status.LastEventId,
		"last_event_at":        status.LastEventAt,
		"last_connected_at":    status.LastConnectedAt,
		"last_disconnected_at": status.LastDisconnectedAt,
		// 输出累计断连和 register replay 统计。
		"disconnect_count":              status.DisconnectCount,
		"register_replay_count":         status.RegisterReplayCount,
		"register_replay_failure_count": status.RegisterReplayFailureCount,
		// 输出最近一次错误信息，便于无需翻日志也能定位最近失败。
		"last_error_kind": status.LastErrorKind,
		"last_error":      status.LastError,
		"last_error_at":   status.LastErrorAt,
	}
}
