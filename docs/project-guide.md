# Firefly Go Layout (Service Template)

## 项目概述
`go-layout` 是 **Firefly 微服务框架** 的官方 Go 语言项目模板。它不仅仅是一个代码骨架，更是一套经过生产环境验证的最佳实践集合。

本模板旨在为开发者提供一个**开箱即用**、**结构清晰**、**易于扩展**的微服务起点。它预置了微服务开发所需的通用基础设施，让开发者能够专注于业务逻辑的实现。

### 核心特性
- **标准分层架构**：基于 DDD（领域驱动设计）思想，清晰划分 `Service` (接口)、`Biz` (业务)、`Data` (数据) 层。
- **依赖注入**：完全集成 Google `Wire`，实现编译期依赖注入，保证代码的模块化和可测试性。
- **协议优先**：集成 `Buf` 和 `gRPC`，通过 Proto 定义驱动开发，自动生成接口代码和验证逻辑 (`protovalidate`)。
- **配置管理**：启动时读取 `bootstrap.json` 与 `consul.json`，再通过 Consul Store 统一拉取多环境运行配置。`config` 数据面支持热更新，但当前模板默认只做启动期加载，未内置运行时 `Watcher` 重载链路。
- **统一运行托管**：默认接入 `go-consul/agent.Agent`，统一托管业务 gRPC、management 端口和 sidecar-agent watch/replay 生命周期。
- **目标身份链路**：入站由 authz 写入 `x-firefly-authz-sign`，服务侧可按需本地验签；出站由 `go-micro/invocation` 清理旧上下文，并在接入 provider 后覆盖 `x-firefly-service-authority`。
- **数据转换**：集成 `Goverter`，自动生成高效的 DTO <-> PO/DO 转换代码，拒绝反射。
- **统一基础设施**：预置了 `GORM` (MySQL), `Redis`, `Logger` 等常用组件的封装和最佳配置。
- **示例模块**：内置完整的 `Demo` 模块，展示了从 API 定义到数据库存储的完整链路，作为开发的参考范本。

## 快速开始

### 1. 环境准备
确保本地已安装以下工具：
- **Go** (>= 1.25.1)
- **Buf** (用于 Proto 管理): `npm install -g @bufbuild/buf` 或参考官方文档
- **Wire** (用于依赖注入): `go install github.com/google/wire/cmd/wire@latest`
- **Goverter** (用于数据转换): `go install github.com/jmattheis/goverter/cmd/goverter@latest`
- **Protoc-Gen-Go** 相关插件与 **protoc-gen-gateway-manifest** (参考 `buf.gen.yaml`)

### 2. 初始化项目
假设你要创建一个名为 `account-service` 的新服务：

1. **克隆模板**
   ```bash
   git clone https://github.com/your-org/go-layout.git account-service
   cd account-service
   rm -rf .git  # 移除模板的 git 历史
   git init     # 初始化新仓库
   ```

2. **重命名模块**
   在 IDE 中全局替换 `go-layout` 为你的模块名（例如 `account-service`）。
   - 修改 `go.mod`
   - 修改 `cmd/server/main.go` 等文件中的导入路径

3. **清理示例代码**
   `Demo` 模块仅供参考。在熟悉架构后，你可以：
   - 删除 `internal/biz/demo.go`, `internal/biz/model/demo.go`, `internal/biz/repo/demo.go`
   - 删除 `internal/data/demo.go`, `internal/data/entity/demo.go`
   - 删除 `internal/service/demo.go`
   - 删除 `internal/biz/convert/demo.go`
   - **注意**：删除后需要重新运行 `wire ./cmd/server`（或执行 `make init`）生成依赖注入代码。

### 3. 运行服务
```bash
# 直接按当前代码启动服务
make run
```
当前 `make run` 仅执行 `go run ./cmd/server`，通过 `agent.Agent.Run(ctx)` 启动业务 gRPC、management 端口与 sidecar-agent 生命周期。
如果你改动了 Proto、DTO 或依赖注入注册，请先执行 `make init`；如只需更新生成代码，也可按需执行 `make generate` 或 `make dto`。

### 4. 管理端口
- 默认管理端口来自 `bootstrap.json.managed_port`，未配置时回落到 `bootstrap.json.server_port + 1`。
- 常用探针：
  - `GET /health`: 存活检查
  - `GET /ready`: 就绪状态与 sidecar 接管状态
  - `GET /info`: 构建信息、监听地址、管理端口和 sidecar 快照
  - `GET /metrics`: Prometheus 指标

## 工具链指南

### Buf (Protobuf 管理)
本项目不直接包含 `.proto` 文件，而是假设 Proto 定义在独立的仓库中管理（推荐做法）。
- `buf.gen.yaml`: 定义了如何从 Proto 生成 Go 代码和 `gateway.manifest.json`。
- **生成代码**：通常通过 CI/CD 管道或脚本执行 `buf generate`，生成的 Go 代码和 `dep/protobuf/gen/gateway.manifest.json` 位于 `dep/protobuf/gen`。
- 模板不提供业务服务级 api-gateway descriptor 配置；namespace descriptor 的生成与发布由 proto 仓库和 Firefly CLI 负责。
- **服务 token 客户端**：公开 demo 模板不默认生成 `acme.auth.token.v1`。实际业务服务接入 service authority 时，需要在自身 `buf.gen.yaml` 输入中加入 `acme.auth.token.v1`，但 manifest 的 `include_package_prefix` 仍只覆盖当前业务服务包，避免把 auth 的接口注册成当前服务能力。

### Authz 与 Service Authority
- `bootstrap.json` 默认不写 `authz_verification`。未配置时 gRPC middleware 只构造 `service.Context`；显式配置后启动阶段会加载 Ed25519 公钥，并对 `x-firefly-authz-sign` 做本地验签。
- Firefly current 身份入口只使用 `x-firefly-user-authority` 和 `x-firefly-service-authority`。`Authorization` 不作为模板身份入口。
- `internal/dep/client.go` 的 `NewServiceAuthorityProvider` 是业务服务获取 service token 的统一接入点。业务服务生成 auth token client 后，在这里使用 `ConnectionManager.Dial(...)` 直连 auth 服务调用 `GenerateServiceToken(app_id, app_secret)`，再用 `authz.NewServiceAuthorityToken` 包装返回值。
- 获取 service token 不能走 `UnaryInvoker` 或 `RemoteServiceManaged`，因为它们会反过来依赖 provider 注入 `x-firefly-service-authority`，容易形成递归依赖。
- 出站调用由 `go-micro/invocation.UnaryInvoker` 统一处理 metadata：透传用户 authority 和短 TTL authz sign，清理上一跳普通身份 metadata，并在 provider 存在时覆盖当前服务的 service authority。

### Wire (依赖注入)
- **入口**：`cmd/server/wire.go`
- **各层 Provider**：每个层级 (`internal/biz`, `internal/data`, `internal/service` 等) 都有一个 `core.go`，定义了该层的 `ProviderSet`。
- **新增组件**：当你新增一个 Repo 或 Service 时，记得将其构造函数加入到对应层的 `core.go` 中，然后执行 `wire ./cmd/server`（或 `make init`）。

### Goverter (数据转换)
- **定义**：在 `internal/biz/convert` 中定义接口。
- **生成**：在项目根目录运行 `make dto`（等价于 `goverter gen ./internal/biz/convert`）。生成代码位于 `dep/dto/`，通过 `internal/dto` 适配为 Biz 层依赖的接口。

## 目录导航
- [目录结构详解](directory-structure.md)
- [核心概念说明](core-concepts.md)
- [数据流转图解](data-flow.md)
- [架构分层深度解析](architecture.md)
- [开发最佳实践](best-practices.md)
