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

type SidecarStatusProvider interface {
	Status() agent.Status
}

// AppManagedServer 暴露服务健康、就绪、信息和指标端点。
type AppManagedServer struct {
	lister net.Listener
	srv    *http.Server
}

// NewAppManagedServer 创建管理端口，并对外暴露 sidecar 生命周期状态。
func NewAppManagedServer(bootstrapConfig *conf.BootstrapConfig, provider *telemetry.Providers, sidecarStatusProvider SidecarStatusProvider) *AppManagedServer {
	l, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConfig.ManagedPort), 10)))
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
			"sidecar": buildSidecarStatus(bootstrapConfig, sidecarStatusProvider),
		})
	})
	mux.HandleFunc("/info", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"service_name":        bootstrapConfig.Service.Name,
			"service_namespace":   bootstrapConfig.Service.Namespace,
			"service_instance_id": bootstrapConfig.App.InstanceId,
			"service_endpoint":    bootstrapConfig.ServiceEndpoint(),
			"management_endpoint": bootstrapConfig.ManagementEndpoint(),
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

func (ist *AppManagedServer) Addr() string {
	if ist == nil || ist.lister == nil {
		return ""
	}
	return ist.lister.Addr().String()
}

func (ist *AppManagedServer) Serve() error {
	err := ist.srv.Serve(ist.lister)
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (ist *AppManagedServer) Stop() {
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
func buildSidecarStatus(bootstrapConfig *conf.BootstrapConfig, provider SidecarStatusProvider) map[string]any {
	status := agent.Status{}
	if provider != nil {
		status = provider.Status()
	}

	sidecarConfig := bootstrapConfig.SidecarAgentConfig()
	return map[string]any{
		"base_url":                      sidecarConfig.BaseURL,
		"grace_period":                  sidecarConfig.GracePeriod,
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
