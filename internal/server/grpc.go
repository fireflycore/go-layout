package server

import (
	"errors"
	demo "go-layout/dep/protobuf/gen/acme/demo/v1"
	"go-layout/internal/conf"
	"go-layout/internal/service"
	"net"
	"strconv"
	"time"

	"github.com/fireflycore/go-micro/logger"
	gm "github.com/fireflycore/go-micro/middleware/grpc"
	ggm "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
)

// GrpcServer 表示业务 gRPC 服务实例。
type GrpcServer struct {
	lister net.Listener
	srv    *grpc.Server
}

// NewGrpcServer 创建 gRPC 服务实例。
func NewGrpcServer(
	log *logger.AccessLogger,
	bootstrapConfig *conf.BootstrapConfig,

	demoService *service.DemoService,
) *GrpcServer {
	// gRPC 监听地址固定绑定所有网卡，具体暴露端口由 bootstrap 控制。
	listen, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConfig.ServerPort), 10)))

	if err != nil {
		panic(err)
	}

	// 中间件顺序保持固定：先 recovery，再建立 ServiceContext，最后做校验错误映射和访问日志。
	srv := grpc.NewServer(
		grpc.StatsHandler(gm.NewOtelServerStatsHandler()),
		grpc.UnaryInterceptor(ggm.ChainUnaryServer(
			recovery.UnaryServerInterceptor(),
			// 在服务入口先建立统一的 ServiceContext，供日志与后续业务链路复用。
			gm.NewServiceContextUnaryInterceptor(gm.ServiceContextInterceptorOptions{
				ServiceAppId:      bootstrapConfig.App.Id,
				ServiceInstanceId: bootstrapConfig.App.InstanceId,
			}),

			gm.ValidationErrorToInvalidArgument(),
			gm.NewAccessLogger(log),
		)),
	)

	// 注册服务级 health check，供 sidecar 与管理链路探活。
	healthServer := health.NewServer()
	healthServer.SetServingStatus(bootstrapConfig.Service.Name, grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(srv, healthServer)

	// 显式注册本服务暴露的业务服务。
	demo.RegisterDemoServiceServer(srv, demoService)

	return &GrpcServer{
		lister: listen,
		srv:    srv,
	}
}

// Serve 启动 gRPC 服务。
func (ist *GrpcServer) Serve() error {
	err := ist.srv.Serve(ist.lister)
	// 优雅关闭时的 ErrServerStopped 不视为异常退出。
	if errors.Is(err, grpc.ErrServerStopped) {
		return nil
	}
	return err
}

// Addr 返回 gRPC 实际监听地址。
func (ist *GrpcServer) Addr() string {
	if ist == nil || ist.lister == nil {
		return ""
	}
	return ist.lister.Addr().String()
}

// Stop 优先尝试优雅关闭，超时后再强制停止。
func (ist *GrpcServer) Stop() {
	done := make(chan struct{})
	go func() {
		ist.srv.GracefulStop()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		ist.srv.Stop()
	}
}
