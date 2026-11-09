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
    - 准备独立 AuthzSign Ed25519 公钥并配置 `authz_verification.public_key_path/kid/issuer`；示例只提供路径，不包含密钥。缺失配置或公钥加载失败时，业务 gRPC 在监听前启动失败。
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
- 服务上下文：gRPC unary/stream 入口注入 `go-micro/service.Context`；本地 AppId/InstanceId 来自配置，实际方法来自适配器。必须有用户的用例通过 `service.RequireUserFromContext` 读取用户；允许无用户的用例继续处理可选 `UserContext`。
- 本地验签：业务入口必须配置 `authz_verification` 并核对签名 `service_app_id/service_method/service_path`，不能用签名反填本地事实。公共业务接口同样需要本跳证明；健康 Check/List/Watch 精确跳过，management HTTP 端口保持独立。
- 远程调用：`internal/dep/client.go` 保留 `ConnectionManager / UnaryInvoker / RemoteServiceManaged` 模板，新增下游服务时集中登记 `invocation.DNS`。
- 服务身份：`NewServiceAuthorityProvider` 通过精确方法的 `BootstrapUnaryInvoker` 调用 `GenerateServiceToken`，仅使用本地 AppId/app_secret，清空业务 metadata，保留取消和期限，不递归依赖自身 provider。
- 出站传播：每次独立构造允许列表，私有继承用户凭证，服务凭证来自当前 provider，不传播上一跳 AuthzSign；缺失或获取失败直接拒绝，不降级为用户直调。Invoke 方法来自当前实际入口，OTel 使用当前活动 span。
- ORM 日志：通过 `WithLogIdentityReader` 注入框架五项可信身份，不解析原始 metadata 或增加身份集合。

## Context 迁移边界

当前使用 Go `1.27.3`，依赖为 `go-micro v1.7.6`、`go-consul v0.4.7`、`gormx v1.0.5`。签名协议沿用 `v1.7.3`，与更早的旧协议不兼容，须协调签发端及全部在线验签端一起切换或回退，不支持双读和缺字段回退。

Demo 创建需要用户上下文；无用户列表返回空结果，不能用服务身份替代用户或删除原有用户过滤。模板不检查用户与服务租户相等，也不代替业务所有权与数据范围规则。

当前保留的本地测试覆盖签名绑定、伪造输入、访问日志脱敏，以及真实 TCP 健康 Check/List/Watch。未完成完整模板生成、Consul、数据库、外部授权提供方或 Envoy 环境联调，也未在本模板复验 bootstrap、取消/超时及并发出站，不代表模板生成后的业务链已经端到端验收。
