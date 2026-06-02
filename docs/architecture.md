# 架构分层与依赖注入

本文档深入解析 `go-layout` 的分层架构实现细节，特别是 **Wire 依赖注入** 和 **配置管理** 在本模板中的具体落地方式。

## 架构概览

遵循整洁架构 (Clean Architecture) 原则，依赖方向由外向内：
`Server` -> `Service` -> `Biz` (核心) <- `Data`

但为了工程实现的便利性，我们在物理目录结构上做了一些适配，通过 Go 的 `interface` 和 `Wire` 来管理依赖。

## 依赖注入 (Wire) 详解

`go-layout` 深度集成了 Google Wire，通过依赖注入管理各个组件的生命周期。理解这一点对于在模板中新增功能至关重要。

### 1. 依赖树的构建
依赖树的根节点在 `cmd/server/wire.go`。

```go
// cmd/server/wire.go
func wireApp() (*App, error) {
    panic(wire.Build(
        dep.ProviderSet,
        conf.ProviderSet,
        data.ProviderSet,
        dto.ProviderSet,
        biz.ProviderSet,
        service.ProviderSet,
        server.ProviderSet,
        NewApp,
    ))
}
```

### 2. 各层 ProviderSet
每个层级目录下都有一个 `core.go`，定义了该层对外提供的组件集合。

*   **`internal/data/core.go`**:
    ```go
    var ProviderSet = wire.NewSet(
        NewConsul,
        NewRedis,
        NewMysql,
        NewConfigStore,
        NewData,

        NewDemoRepo,
    )
    ```

*   **`internal/biz/core.go`**:
    ```go
    var ProviderSet = wire.NewSet(
        NewDemoUseCase, // 注册 DemoUseCase
        // 新增的 UseCase 在这里注册
    )
    ```

*   **`internal/service/core.go`**:
    ```go
    var ProviderSet = wire.NewSet(
        NewDemoService,
    )
    ```

### 3. 如何新增依赖
当你创建了一个新的 Repo (`UserRepo`) 和 UseCase (`UserUseCase`) 时：
1.  在 `internal/data/core.go` 中添加 `NewUserRepo`。
2.  在 `internal/biz/core.go` 中添加 `NewUserUseCase`。
3.  在 `internal/service/core.go` 中添加 `NewUserService`。
4.  执行 `wire ./cmd/server`（或 `make init`）重新生成 `wire_gen.go`。

## 配置管理实现

`go-layout` 的配置系统已经统一到 `bootstrap + consul + Store.LoadStoreConfig` 模型：启动时读取引导配置，然后直连 Consul Store 拉取运行期配置。`config` 数据面本身支持 `Store.Get + Watcher.Watch`，但当前模板默认只装配启动期读取链路。

### 1. 引导配置 (Bootstrap)
一切始于 `conf/bootstrap.json`。这是服务启动时读取的第一个文件，定义了服务身份、监听端口、sidecar 与 telemetry 基础信息。

```json
{
  "app": {
    "id": "demo-service",
    "env": "dev",
    "name": "go-layout",
    "secret": "...",
    "version": "v0.0.1"
  },
  "service": {
    "name": "go-layout",
    "type": "svc",
    "namespace": "default",
    "cluster_domain": "cluster.local",
    "port": 9090,
    "weight": 100
  },
  "server_port": 10500,
  "managed_port": 10501,
  "sidecar_agent": { ... },
  "logger": { ... },
  "telemetry": { ... }
}
```

### 2. 配置加载器 (Conf Loader)
位于 `internal/conf/`。
例如 `mysql.go` 定义了如何加载 MySQL 配置：
- **引导配置**：读取 `conf/consul.json`，创建 Consul Client 和 `microConfig.Store`。
- **运行期配置**：根据 `bootstrap.json` 中的 `app.id/app.env/app.secret`，通过 `LoadStoreConfig` 从 Consul Store 拉取 `mysql`、`redis` 等配置。
- **TLS 口径**：TLS 字段直接使用证书文件路径，不再支持把证书正文下发到配置里后由模板落盘。
- **注意**：`config` 数据面支持热更新，但当前模板默认仅在服务启动时加载一次；如需热更新，需要业务服务显式接入 `Watcher` 并实现组件重载策略。

### 3. 在代码中使用配置
配置加载后，通常通过依赖注入传递给 Data 层。
例如 `internal/data/data.go`:

```go
func NewMysql(bootstrapConfig *conf.BootstrapConfig, mysqlConfig *gormx.MysqlConfig) (*gorm.DB, error) {
    mysqlConfig.WithLoggerConsole(bootstrapConfig.Logger.Console)
    mysqlConfig.WithAutoMigrate(false)
    db, err := gormx.NewMysql(mysqlConfig)
    if err != nil {
        return nil, err
    }
    return db.DB, nil
}
```

## 运行托管实现

当前模板以 `go-consul/agent.Agent` 作为裸机 sidecar-agent 桥接单入口：

- `internal/server/register.go` 基于 `agent.ServiceOptions + gateway.manifest.json` 构造 `agent.Agent`，服务能力由 manifest-first 注册链路提供。
- `internal/server/server.go` 通过 `Agent.ConfigureRun(...)` 注入业务 `Serve/Shutdown` 回调。
- `cmd/server/app.go` 最终调用 `Agent.Run(ctx)`，统一驱动 `gRPC + management + sidecar watch/replay`。
- `internal/server/managed.go` 暴露 `/health`、`/ready`、`/info`、`/metrics`，其中 `/ready` 和 `/info` 会返回 `agent.Status` 摘要。

## 服务上下文

gRPC 服务端入口通过 `gm.NewServiceContextUnaryInterceptor(...)` 注入 `go-micro/service.Context`。

- Service 层通过 `service.FromContext(ctx)` 读取用户与服务身份。
- Biz/Data 层不再解析 gRPC metadata。
- `authz_verification` 为空时只解析普通 metadata；显式配置后会加载 authz Ed25519 公钥并校验 `x-firefly-authz-sign`。
- 服务侧本地验签会校验 `target_app_id` 必须等于当前服务的 `app.id`，避免把别的 route 授权结果复用到当前服务。

## 出站调用与服务身份

出站调用统一走 `go-micro/invocation.UnaryInvoker`：

- 透传 `x-firefly-user-authority`，保证用户身份可以贯穿整条链路。
- 透传短 TTL `x-firefly-authz-sign`，供下一跳 authz 复用身份解析结果，但下一跳仍必须按当前 route 重新做权限判定。
- 清理上一跳 authz 注入的普通身份 metadata，避免服务间调用复用上一跳的 `invoke_app_id/target_app_id/api_path`。
- `NewServiceAuthorityProvider` 是模板预留的 service token 获取点；实际业务服务生成 `acme.auth.token.v1` client 后，在这里调用 auth 服务 `GenerateServiceToken`，并由 `UnaryInvoker` 每跳覆盖 `x-firefly-service-authority`。
- `Authorization` 不属于 Firefly current 身份入口，模板不再注入或读取它。

## 总结
- **Wire** 粘合了所有层级，修改组件依赖关系后必须重新生成。
- **Bootstrap** 提供 `app/service` 身份、业务端口、管理端口、telemetry 与 sidecar-agent 基础信息。
- **Conf Loader** 统一走 Consul Store 读取运行期配置，TLS 证书仅保留路径引用；`config` 具备热更新能力，但当前模板默认仅装配启动期加载。
- **Agent** 统一托管业务服务运行和 sidecar-agent 生命周期，不再使用旧 `ServiceLifecycle/ManagedServer` 主线。
- **Authz** 默认不强制本地验签；生产接入时通过 `authz_verification` 显式开启，并通过 auth 服务签发 service token。
- **Makefile** 当前将 `generate/init/run/build` 明确拆分：`run` 只负责启动，生成与依赖整理需要显式执行 `make init`。
