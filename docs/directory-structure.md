# 项目目录结构详解

本文档详细解析 `go-layout` 模板的目录结构，帮助开发者快速定位代码并理解各个目录的职责。

## 根目录概览

```markdown
├── cmd/                        # 应用程序入口
│   └── server/
│       ├── main.go             # [核心] 程序主入口，负责初始化配置、日志和应用生命周期
│       ├── wire.go             # [核心] Wire 依赖注入定义文件
│       └── wire_gen.go         # [生成] Wire 生成的代码，不要手动修改
├── conf/                       # 配置文件目录
│   ├── bootstrap.json          # [配置] 引导配置，定义服务身份、端口、sidecar 与 telemetry
│   └── consul.json             # [配置] Consul 连接示例，用于创建 Store 读取运行期配置
├── dep/                        # [生成] 外部依赖/生成代码存放区
│   ├── dto/                    # [生成] Goverter 生成的数据转换实现代码
│   └── protobuf/gen/           # [生成] Buf 生成的 gRPC/Proto 结构体代码
├── docs/                       # 项目文档
├── internal/                   # [核心] 业务代码私有目录（Go 语言机制，外部无法 import）
│   ├── biz/                    # [业务] 业务逻辑层 (Business Logic)
│   ├── conf/                   # [配置] 配置加载与解析逻辑
│   ├── data/                   # [数据] 数据访问层 (Data Access)
│   ├── dep/                    # [依赖] 基础设施适配层 (Infrastructure)
│   ├── dto/                    # [转换] DTO 注册与入口
│   ├── server/                 # [服务] gRPC、management、sidecar 生命周期托管
│   └── service/                # [接口] 应用服务层 (Application Service)
├── buf.gen.yaml                # [工具] Buf 生成 Go 代码与 gateway.manifest.json
├── go.mod                      # [依赖] Go 模块定义
└── makefile                    # [工具] 常用命令封装，当前拆分为 generate/init/run/build
```

## Internal 目录深度解析

`internal` 是开发者的主要工作区。

### 1. `internal/biz` (业务逻辑层)
**职责**：定义业务接口、实现核心业务逻辑。该层**不应该**依赖 `data` 或 `service` 层，保持纯净。

- `convert/`: **[开发区]** 定义 DTO 与内部模型（PO/DO）的转换接口。
  - `demo.go`: 示例转换接口。
- `model/`: **[开发区]** 定义领域对象 (DO)。
  - `demo.go`: 示例领域对象。
- `repo/`: **[开发区]** 定义数据访问接口 (Repository Interfaces)。**注意：这里只定义接口，实现在 `data` 层。**
  - `demo.go`: `DemoRepo` 接口定义。
- `core.go`: **[配置]** Wire ProviderSet，注册本层的所有 UseCase。
- `demo.go`: **[开发区]** `DemoUseCase` 实现，编排业务逻辑。

### 2. `internal/data` (数据访问层)
**职责**：实现 `biz/repo` 中定义的接口，负责具体的数据持久化（MySQL, Redis 等）。

- `entity/`: **[开发区]** 定义持久化对象 (PO)，即数据库表结构映射 (GORM Model)。
  - `demo.go`: `Demo` 表结构定义。
- `core.go`: **[配置]** Wire ProviderSet，注册本层的所有 Repo 实现。
- `data.go`: **[基础设施]** 数据库、Redis、Consul Store 客户端的初始化与连接管理。
- `demo.go`: **[开发区]** `DemoRepo` 的具体实现 (DAO)。
- `transaction.go`: 事务支持实现。

### 3. `internal/service` (应用服务层)
**职责**：实现 gRPC 接口，处理请求参数验证，调用 `biz` 层逻辑，返回响应。

- `core.go`: **[配置]** Wire ProviderSet，注册本层的所有 Service。
- `demo.go`: **[开发区]** `DemoService` 实现，直接对应 Proto 定义的 Service。
  - 这里进行 `protovalidate` 参数校验。
  - 这里读取 gRPC middleware 注入的 `go-micro/service.Context`，不再解析 metadata。

### 4. `internal/server` (服务启动层)
**职责**：构建和启动 gRPC 服务器、management 端口，并把 sidecar 生命周期接入统一托管入口。

- `grpc.go`: **[配置]** 业务 gRPC Server 封装，负责监听、OTel、访问日志和优雅停机。
- `managed.go`: **[配置]** 管理端口，暴露 `/health`、`/ready`、`/info`、`/metrics`。
- `register.go`: **[配置]** 基于 `agent.ServiceOptions + gateway.manifest.json` 组装 `go-consul/agent.Agent`。
- `server.go`: **[配置]** 通过 `Agent.ConfigureRun(...)` 注入业务 `Serve/Shutdown`，统一托管 `gRPC + management + sidecar watch/replay`。
- `build_info.go`: **[配置]** 管理端口对外暴露的构建信息。

### 5. `internal/conf` (配置层)
**职责**：加载和解析应用配置。

- `bootstrap.go`: 引导配置加载，统一装配 `app/kernel/service/logger/telemetry`、业务端口、管理端口和 sidecar-agent 配置。
- `consul.go`: 读取 `conf/consul.json`，创建 Consul 连接所需的基础配置。
- `mysql.go`, `redis.go`: 组件配置加载器。
  - 启动阶段基于 `microConfig.Store` 直连 Consul Store 拉取配置。
  - 当前模板默认不提供本地 `mysql.json/redis.json` 热切换模式。

### 6. `internal/dep` (依赖适配层)
**职责**：封装第三方库或基础设施，防止外部依赖污染业务代码。

- `client.go`: 基于 `invocation` 的统一远程调用连接管理器，负责 DNS target、OTel、metadata 透传和服务身份兜底注入。
- `telemetry.go`: 基于 `bootstrap.json` 的 app/service 身份创建 OTel providers。
- 当前目录主要承载 `telemetry/logger/invocation` 相关依赖注入，不再单独保留旧版 remote logger repo。

## 开发者修改指南

| 任务 | 涉及目录/文件 | 说明 |
| :--- | :--- | :--- |
| **新增 API** | `dep/protobuf/gen/` (外部) -> `internal/service/` | 首先在 Proto 仓库定义，更新 `dep`，然后在 `service` 实现接口。 |
| **新增业务逻辑** | `internal/biz/` | 在 `biz` 创建 UseCase，定义 Repo 接口。 |
| **新增数据库表** | `internal/data/entity/` -> `internal/data/` | 定义 PO 结构体，实现 Repo 接口。 |
| **新增配置项** | `internal/conf/` | 修改 `BootstrapConfig` 或新增配置 Loader。 |
| **依赖注入注册** | 各层的 `core.go` -> `cmd/server/wire.go` | 每次新增 struct 需在对应的 `core.go` 中注册，并执行 `wire ./cmd/server`（或 `make init`）。 |
| **日常启动** | `makefile` + `cmd/server/` | 当前 `make run` 只负责 `go run ./cmd/server`；若改动生成链路或依赖注入，需要先显式执行 `make init`。 |
| **运行链路排查** | `cmd/server/` + `internal/server/` + `internal/conf/bootstrap.go` | 统一从 `agent.Agent`、management 端口和 sidecar-agent 状态入手排查。 |

## 示例文件说明
项目中包含的 `demo.go` 文件（分布在各层）是**参考实现**。
- 它们展示了一个完整的 CRUD 流程。
- 在开始新项目时，请**参考**它们的写法，然后**删除**或**替换**为你的实际业务代码。
