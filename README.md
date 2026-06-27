# Firefly Go Layout

`go-layout` 是 Firefly 微服务框架的 Go 版本标准项目模板。它提供了一套标准化的目录结构和基础设施配置，旨在帮助开发者快速构建规范的微服务应用。

本模板基于 **[go-micro](https://github.com/fireflycore/go-micro)**（Firefly 微服务框架的 Go 版本核心库）构建。
默认启动链路为：读取 `conf/bootstrap.json` 和 `conf/consul.json`，通过 Consul Store 加载运行期配置，再以 `go-consul/agent.Agent` 托管业务 gRPC、management 端口和 sidecar-agent watch/replay 生命周期。`config` 数据面本身支持 `watch`/热更新，但 `go-layout` 默认模板当前只接入启动期加载，如需运行时热更新需要业务服务显式补充 `Watcher` 装配与组件重载策略。

> [在线文档](https://firefly.lhdht.cn/guide/)

## 快速开始

### 初始化项目

1.  **Clone 项目**

    ```bash
    git clone https://github.com/fireflycore/go-layout.git my-project
    cd my-project
    ```

2.  **重命名模块**

    使用提供的脚本将模块名（默认 `go-layout`）替换为你自己的模块名（例如 `github.com/myuser/my-project`）。

    **Windows:**
    ```cmd
    .\rename_project.bat github.com/myuser/my-project
    ```

    **Linux / macOS:**
    ```bash
    chmod +x rename_project.sh
    ./rename_project.sh github.com/myuser/my-project
    ```

3.  **整理依赖**

    ```bash
    go mod tidy
    ```

4.  **准备配置**

    - 修改 `conf/bootstrap.json`，填写 `app`、`service`、端口、sidecar-agent、logger 与 telemetry 基础信息。
    - 修改 `conf/consul.json`，填写 Consul 地址、协议、数据中心与令牌。

### 常用命令

- `make generate`: 执行 `buf generate` 生成 Go 代码、`dep/protobuf/gen/gateway.manifest.json`，并生成 DTO
- `make init`: 执行生成链路、`wire ./cmd/server` 和 `go mod tidy`
- `make run`: 直接执行 `go run ./cmd/server`
- `make build`: 先执行 `make init`，再注入构建信息并编译服务

模板侧的 `buf generate` 只负责当前服务的 Go 代码和 `gateway.manifest.json`；namespace 级 descriptor 的生成与发布由 proto 仓库和 Firefly CLI 单独负责。

## 当前框架主线

- 启动托管：`App.Run(ctx)` 进入 `agent.Agent.Run(ctx)`，由 Agent 统一驱动 `gRPC + management + sidecar watch/replay`。
- 服务注册：`internal/server/register.go` 基于 `agent.ServiceOptions + gateway.manifest.json` 组装 `agent.Agent`，服务能力统一由 manifest-first 注册链路提供。
- 管理端口：`internal/server/managed.go` 暴露 `/health`、`/ready`、`/info`、`/metrics`，并在 `/ready`、`/info` 输出 sidecar 状态摘要。
- 服务上下文：gRPC 入口通过 `gm.NewServiceContextUnaryInterceptor` 注入 `go-micro/service.Context`，业务代码不再解析旧 `invocation.UserContextMeta`。
- 本地验签：`authz_verification` 不再默认写入 `bootstrap.json`；配置对象为空时只解析普通 metadata，显式配置后才加载 authz 公钥并校验 `x-firefly-authz-sign`。
- 远程调用：`internal/dep/client.go` 保留 `ConnectionManager / UnaryInvoker / RemoteServiceManaged` 模板，新增下游服务时集中登记 `invocation.DNS`。
- 服务身份：出站调用只使用 `x-firefly-user-authority`、`x-firefly-service-authority` 和 `x-firefly-authz-sign` 主线；模板提供 `NewServiceAuthorityProvider` 装配点，实际业务服务生成 `acme.auth.token.v1` client 后在此接入 `GenerateServiceToken`。
