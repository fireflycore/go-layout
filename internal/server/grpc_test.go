package server

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"go-layout/internal/conf"
	"go-layout/internal/service"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/fireflycore/go-micro/authz"
	"github.com/fireflycore/go-micro/constant"
	"github.com/fireflycore/go-micro/logger"
	gm "github.com/fireflycore/go-micro/middleware/grpc"
	framework "github.com/fireflycore/go-micro/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func contextConfig(t *testing.T) (*conf.BootstrapConfig, ed25519.PrivateKey) {
	t.Helper()
	public, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKIXPublicKey(public)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "authz-public.pem")
	if err := os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: der}), 0600); err != nil {
		t.Fatal(err)
	}
	cfg := &conf.BootstrapConfig{AuthzVerification: &authz.VerificationConfig{PublicKeyPath: path}}
	cfg.App.Id, cfg.App.InstanceId = "local-service", "local-instance"
	cfg.Service.Name = "go-layout"
	return cfg, private
}

// 测试证明只使用临时随机密钥，与生产签发端的跨仓签名测试互相独立。
func contextProof(t *testing.T, key ed25519.PrivateKey, app, method, path string) string {
	t.Helper()
	now := time.Now().Unix()
	head, err := json.Marshal(map[string]string{"alg": "EdDSA", "typ": "JWT", "kid": "default"})
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(framework.AuthzSign{AuthzSignContext: framework.AuthzSignContext{
		Issuer: authz.DefaultIssuer, SubjectId: "service-A", SubjectType: "service", SubjectAppId: "service-A", InvokeServiceAppId: "service-A",
		Method: "GET", Path: "/demo", ServiceAppId: app, ServiceMethod: method, ServicePath: path,
		Decision: "allow", DecisionId: "test-decision", IssuedAt: now, NotBefore: now, ExpiresAt: now + 60,
	}, UserContext: &framework.UserContext{UserId: "user-1", AppId: "login-app", TenantId: "user-tenant", Session: "private-session"}})
	if err != nil {
		t.Fatal(err)
	}
	data := base64.RawURLEncoding.EncodeToString(head) + "." + base64.RawURLEncoding.EncodeToString(body)
	return data + "." + base64.RawURLEncoding.EncodeToString(ed25519.Sign(key, []byte(data)))
}

func TestBusinessVerificationRequiredAndBoundToActualEntry(t *testing.T) {
	if _, err := serviceContextOptions(&conf.BootstrapConfig{}); err == nil {
		t.Fatal("missing verifier accepted")
	}
	cfg, key := contextConfig(t)
	opts, err := serviceContextOptions(cfg)
	if err != nil {
		t.Fatal(err)
	}
	interceptor := gm.NewServiceContextUnaryInterceptor(opts)
	const path = "/acme.demo.v1.DemoService/GetDemoInfo"
	for _, binding := range []string{"valid", "missing", "wrong-service", "wrong-method", "wrong-path"} {
		t.Run(binding, func(t *testing.T) {
			app, method, target := cfg.App.Id, "GRPC", path
			switch binding {
			case "wrong-service":
				app = "other"
			case "wrong-method":
				method = "GET"
			case "wrong-path":
				target = "/other"
			}
			md := metadata.Pairs(constant.UserAuthority, "private-user-token", constant.UserId, "forged", "x-firefly-service-app-id", "forged", constant.InvokeMethod, "GRPC", constant.InvokePath, "/demo.A/Entry")
			if binding != "missing" {
				md.Set(constant.AuthzSign, contextProof(t, key, app, method, target))
			}
			called := false
			core, observed := observer.New(zap.InfoLevel)
			_, err := interceptor(metadata.NewIncomingContext(context.Background(), md), nil, &grpc.UnaryServerInfo{FullMethod: path}, func(ctx context.Context, _ any) (any, error) {
				called = true
				c := framework.MustFromContext(ctx)
				if c.ServiceAppId != cfg.App.Id || c.ServiceInstanceId != cfg.App.InstanceId || c.ServicePath != path || c.ServiceMethod != "GRPC" {
					t.Fatal("actual entry overwritten by signature or header")
				}
				if c.UserContext.UserId != "user-1" || c.InvokeAppId != "service-A" || c.InvokePath != "/demo.A/Entry" {
					t.Fatal("trusted identity lost")
				}
				public, _ := metadata.FromIncomingContext(ctx)
				if len(public.Get(constant.UserAuthority)) != 0 || len(public.Get(constant.AuthzSign)) != 0 {
					t.Fatal("credential exposed in business metadata")
				}
				// 与生产日志保持一致，过滤仅供 OTel 日志桥接器使用的原生上下文。
				return gm.NewAccessLogger(logger.NewAccessLogger(zap.New(logger.NewContextOmittingCore(core))))(ctx, nil, &grpc.UnaryServerInfo{FullMethod: path}, func(context.Context, any) (any, error) { return nil, nil })
			})
			if binding == "valid" {
				if err != nil || !called {
					t.Fatalf("valid proof failed: %v", err)
				}
				if observed.Len() != 1 {
					t.Fatal("access log missing")
				}
				for _, entry := range observed.All() {
					text := fmt.Sprint(entry.ContextMap())
					if strings.Contains(text, "private-user-token") || strings.Contains(text, "private-session") || strings.Contains(text, md.Get(constant.AuthzSign)[0]) {
						t.Fatal("default access log exposed credentials or session")
					}
				}
			} else if err == nil || called {
				t.Fatal("missing/mismatched proof reached business handler")
			}
		})
	}
	// 跳过列表是 FullMethod 精确匹配，不能变成健康方法前缀放行。
	_, err = interceptor(context.Background(), nil, &grpc.UnaryServerInfo{FullMethod: health.Health_Check_FullMethodName + "Extra"}, func(context.Context, any) (any, error) { t.Fatal("prefix bypassed verifier"); return nil, nil })
	if status.Code(err) != codes.Unauthenticated {
		t.Fatal("unregistered prefix not rejected")
	}
	cfg.AuthzVerification.PublicKeyPath = "missing-file"
	if _, err := serviceContextOptions(cfg); err == nil {
		t.Fatal("missing key accepted")
	}
}

func TestTemplateHealthUnaryAndStreamRemainIndependent(t *testing.T) {
	cfg, _ := contextConfig(t)
	server := NewGrpcServer(logger.NewAccessLogger(zap.NewNop()), cfg, &service.DemoService{})
	done := make(chan error, 1)
	go func() { done <- server.Serve() }()
	t.Cleanup(func() {
		server.Stop()
		if err := <-done; err != nil {
			t.Error(err)
		}
	})
	// 真实 TCP 回环验证生产构造函数，而不是仅检查 options 字段。
	address := strings.Replace(server.Addr(), "[::]", "127.0.0.1", 1)
	address = strings.Replace(address, "0.0.0.0", "127.0.0.1", 1)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client := health.NewHealthClient(conn)
	if _, err := client.Check(ctx, &health.HealthCheckRequest{Service: cfg.Service.Name}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.List(ctx, &health.HealthListRequest{}); err != nil {
		t.Fatal(err)
	}
	stream, err := client.Watch(ctx, &health.HealthCheckRequest{Service: cfg.Service.Name})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := stream.Recv(); err != nil {
		t.Fatal(err)
	}
	// 收到首条健康状态后取消长连接，避免服务优雅关闭等待到超时。
	cancel()
}
