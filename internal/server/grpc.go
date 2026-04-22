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
)

// GrpcServer 封装模板库业务 gRPC 服务的监听和优雅停机。
type GrpcServer struct {
	lister net.Listener
	srv    *grpc.Server
}

// NewGrpcServer 创建业务 gRPC 服务，并统一挂接 OTel 与访问日志中间件。
func NewGrpcServer(
	log *logger.AccessLogger,
	bootstrapConf *conf.BootstrapConf,

	demoService *service.DemoService,
) *GrpcServer {
	listen, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConf.GetServerPort()), 10)))
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer(
		grpc.StatsHandler(gm.NewOtelServerStatsHandler()),
		grpc.UnaryInterceptor(ggm.ChainUnaryServer(
			recovery.UnaryServerInterceptor(),

			gm.ValidationErrorToInvalidArgument(),
			gm.NewAccessLogger(log),
		)),
	)

	demo.RegisterDemoServiceServer(srv, demoService)

	return &GrpcServer{
		lister: listen,
		srv:    srv,
	}
}

// Serve 阻塞运行 gRPC 服务。
func (ist *GrpcServer) Serve() error {
	err := ist.srv.Serve(ist.lister)
	if errors.Is(err, grpc.ErrServerStopped) {
		return nil
	}
	return err
}

// Addr 返回 gRPC 监听地址。
func (ist *GrpcServer) Addr() string {
	if ist == nil || ist.lister == nil {
		return ""
	}
	return ist.lister.Addr().String()
}

// Stop 在退出阶段优先走优雅停机，超时后再强制关闭。
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
