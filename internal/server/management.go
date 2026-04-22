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

// SidecarStatusProvider 抽象 sidecar 生命周期状态来源，便于管理端口复用。
type SidecarStatusProvider interface {
	Status() agent.Status
}

// ManagementServer 暴露服务健康、就绪、信息和指标端点。
type ManagementServer struct {
	lister net.Listener
	srv    *http.Server
}

// NewManagementServer 创建管理端口，并对外暴露 sidecar 生命周期状态。
func NewManagementServer(bootstrapConf *conf.BootstrapConf, provider *telemetry.Providers, sidecarStatusProvider SidecarStatusProvider) *ManagementServer {
	l, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConf.GetManagementPort()), 10)))
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	if provider != nil && provider.MetricsHandler != nil {
		mux.Handle("/metrics", provider.MetricsHandler)
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})
	mux.HandleFunc("/ready", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"status":  "ready",
			"sidecar": buildSidecarStatus(bootstrapConf, sidecarStatusProvider),
		})
	})
	mux.HandleFunc("/info", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"service_name":        bootstrapConf.GetServiceName(),
			"service_namespace":   bootstrapConf.GetServiceNamespace(),
			"service_instance_id": bootstrapConf.GetServiceInstanceId(),
			"service_endpoint":    bootstrapConf.GetServiceEndpoint(),
			"management_endpoint": net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConf.GetManagementPort()), 10)),
			"app_id":              bootstrapConf.GetAppId(),
			"app_name":            bootstrapConf.GetAppName(),
			"version":             bootstrapConf.GetAppVersion(),
			"build_info":          buildInfo(),
			"telemetry": map[string]any{
				"endpoint": bootstrapConf.GetOtelEndpoint(),
				"insecure": bootstrapConf.GetOtelInsecure(),
				"traces":   bootstrapConf.GetOtelTraces(),
				"metrics":  bootstrapConf.GetOtelMetrics(),
				"logs":     bootstrapConf.GetOtelLogs(),
			},
			"sidecar": buildSidecarStatus(bootstrapConf, sidecarStatusProvider),
		})
	})

	return &ManagementServer{
		lister: l,
		srv: &http.Server{
			Handler: mux,
		},
	}
}

// Addr 返回管理端口地址。
func (ist *ManagementServer) Addr() string {
	if ist == nil || ist.lister == nil {
		return ""
	}
	return ist.lister.Addr().String()
}

// Serve 阻塞运行管理端口。
func (ist *ManagementServer) Serve() error {
	err := ist.srv.Serve(ist.lister)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

// Stop 优雅关闭管理端口。
func (ist *ManagementServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ist.srv.Shutdown(ctx); err != nil {
		_ = ist.lister.Close()
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

// buildSidecarStatus 把生命周期状态快照转换成稳定的管理接口结构。
func buildSidecarStatus(bootstrapConf *conf.BootstrapConf, provider SidecarStatusProvider) map[string]any {
	status := agent.Status{}
	if provider != nil {
		status = provider.Status()
	}

	return map[string]any{
		"base_url":                      bootstrapConf.GetSidecarAgentBaseURL(),
		"grace_period":                  bootstrapConf.GetSidecarGracePeriod(),
		"connected":                     status.Connected,
		"registered":                    status.Registered,
		"last_service_name":             status.LastServiceName,
		"last_service_port":             status.LastServicePort,
		"last_event_type":               status.LastEventType,
		"last_event_id":                 status.LastEventId,
		"last_event_at":                 status.LastEventAt,
		"last_connected_at":             status.LastConnectedAt,
		"last_disconnected_at":          status.LastDisconnectedAt,
		"disconnect_count":              status.DisconnectCount,
		"register_replay_count":         status.RegisterReplayCount,
		"register_replay_failure_count": status.RegisterReplayFailureCount,
		"last_error_kind":               status.LastErrorKind,
		"last_error":                    status.LastError,
		"last_error_at":                 status.LastErrorAt,
	}
}
