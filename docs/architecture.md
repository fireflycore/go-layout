# 架构分层

1. **配置层 (internal/conf)**
   - 职责：统一管理应用配置，支持多配置源 (Local, Etcd, Env)

2. **业务逻辑层 (internal/biz)**
   - 职责：实现核心业务逻辑，保持框架无关性
   - 组件：
     - `UseCase`: 业务用例实现
     - `Repo Interface`: 数据访问接口定义
     - `Convert Interface`: 数据转换接口定义
     - `Model`: 领域对象 (DO, 可选)

3. **数据访问层 (internal/data)**
   - 职责：数据持久化实现，封装数据库细节
   - 组件：
     - `Repo Implementation`: 接口实现 (DAO)
     - `Entity`: 持久化对象 (PO)
     - `Data`: 数据库连接管理

4. **应用服务层 (internal/service)**
   - 职责：对外提供 API 服务，处理协议转换，参数验证 (`protovalidate`)

5. **服务层 (internal/server)**
   - 职责：服务注册、网络层处理 (gRPC Server, HTTP Server)

6. **依赖层 (internal/dep)**
   - 职责：基础设施依赖的封装 (Logger, Remote Clients)

## 集成工具
- **goverter**: 用于生成数据转换代码，避免手动编写重复代码
  - https://goverter.jmattheis.de/
- **buf-cli**: 用于管理 proto 和生成 gRPC 代码
  - https://buf.build/docs/

## 注意事项
- 本项目是一个单独的服务，拥有独立 git 仓库，并非传统的大仓。
- 本项目的 proto 定义是在统一 proto 仓库中，通过 buf-cli 进行管理和生成相关 gRPC 代码。
- **数据转换策略**：
   - 优先使用 buf-cli 生成的 DTO 代码。
   - 在 `internal/biz/convert` 下定义数据转换接口，使用 goverter 生成实现。
   - 避免冗余的数据转换：
     - 如果可以直接使用 DTO，则不强制转换为 DO。
     - Repo 层可以直接返回 DTO 给 Biz 层。
     - 写入时，Biz 层负责将 DTO 转换为 PO 传给 Data 层。
