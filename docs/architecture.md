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
  "env": "dev",
  "port": 10500,
  "app_id": "demo-service",
  "service_name": "go-layout",
  "service_namespace": "default",
  "sidecar_agent": { ... },
  "logger": { ... },
  "telemetry": { ... }
}
```

### 2. 配置加载器 (Conf Loader)
位于 `internal/conf/`。
例如 `mysql.go` 定义了如何加载 MySQL 配置：
- **引导配置**：读取 `conf/consul.json`，创建 Consul Client 和 `microConfig.Store`。
- **运行期配置**：根据 `bootstrap.json` 中的 `app_id/env/app_secret`，通过 `LoadStoreConfig` 从 Consul Store 拉取 `mysql`、`redis` 等配置。
- **TLS 口径**：TLS 字段直接使用证书文件路径，不再支持把证书正文下发到配置里后由模板落盘。
- **注意**：`config` 数据面支持热更新，但当前模板默认仅在服务启动时加载一次；如需热更新，需要业务服务显式接入 `Watcher` 并实现组件重载策略。

### 3. 在代码中使用配置
配置加载后，通常通过依赖注入传递给 Data 层。
例如 `internal/data/data.go`:

```go
func NewMysql(bootstrapConf *conf.BootstrapConf, mysqlConf *gormx.MysqlConf) (*gormx.MysqlDB, error) {
    mysqlConf.WithLoggerConsole(bootstrapConf.Logger.Console)
    mysqlConf.WithAutoMigrate(false)
    return gormx.NewMysql(mysqlConf, ...)
}
```

## 总结
- **Wire** 粘合了所有层级，修改组件依赖关系后必须重新生成。
- **Bootstrap** 提供服务身份、端口、telemetry 与 sidecar 基础信息。
- **Conf Loader** 统一走 Consul Store 读取运行期配置，TLS 证书仅保留路径引用；`config` 具备热更新能力，但当前模板默认仅装配启动期加载。
- **Makefile** 当前将 `generate/init/run/build` 明确拆分：`run` 只负责启动，生成与依赖整理需要显式执行 `make init`。
