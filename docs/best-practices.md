# 开发最佳实践与扩展指南

本文档提供基于 `go-layout` 模板开发新功能的标准工作流和最佳实践建议。

## 新功能开发标准流程 (Checklist)

假设你要开发一个新的业务模块（例如 `Order` 订单模块），请遵循以下步骤：

### 1. 接口定义 (Proto First)
- [ ] 在 Proto 仓库中创建 `order.proto`，定义 Service 和 Message。
- [ ] 定义 API 接口（如 `CreateOrder`）和字段验证规则 (`protovalidate`)。
- [ ] 运行 `buf generate`，确保生成的代码同步到本项目的 `dep/protobuf` 目录。

### 2. 基础设施准备 (Data Layer)
- [ ] **定义 Entity (PO)**: 在 `internal/data/entity/order.go` 中定义数据库表结构。
- [ ] **定义 Repo 接口**: 在 `internal/biz/repo/order.go` 中定义 `OrderRepo` 接口。
- [ ] **实现 Repo**: 在 `internal/data/order.go` 中实现 `OrderRepo`。
- [ ] **注册依赖**: 在 `internal/data/core.go` 的 `ProviderSet` 中添加 `NewOrderRepo`。

### 3. 业务逻辑实现 (Biz Layer)
- [ ] **定义 Converter**: 在 `internal/biz/convert/order.go` 中定义 DTO <-> PO 转换接口。
- [ ] **生成转换代码**: 运行 `go generate ./...` 生成 `dep/dto/order.go`。
- [ ] **实现 UseCase**: 创建 `internal/biz/order.go`，编写 `OrderUseCase`，注入 Repo 和 Converter。
- [ ] **注册依赖**: 在 `internal/biz/core.go` 的 `ProviderSet` 中添加 `NewOrderUseCase`。

### 4. 服务接口实现 (Service Layer)
- [ ] **实现 Service**: 创建 `internal/service/order.go`，实现 Proto 定义的 Server 接口。
- [ ] **参数验证**: 在入口处调用 `protovalidate.Validate(req)`。
- [ ] **调用 Biz**: 将请求转发给 `OrderUseCase`。
- [ ] **注册依赖**: 在 `internal/service/core.go` 的 `ProviderSet` 中添加 `NewOrderService`。

### 5. 服务注册与启动 (Server Layer)
- [ ] **注册 gRPC**: 修改 `internal/server/register.go`，将 `OrderService` 注册到 gRPC Server。
- [ ] **依赖注入**: 运行 `cd cmd/server && wire`，更新依赖注入代码。

### 6. 验证与测试
- [ ] 运行服务 `go run cmd/server/main.go`。
- [ ] 使用 Postman 或 grpcurl 测试新接口。

---

## 编码规范与建议

### 1. 错误处理
- **Service 层**: 负责将 error 转换为 gRPC Status Code。
  - 参数错误 -> `codes.InvalidArgument`
  - 业务错误 -> 自定义 Code 或 `codes.FailedPrecondition`
  - 系统错误 -> 记录日志并返回 `codes.Internal`
- **Biz/Data 层**: 直接返回 Go error，不要依赖 gRPC 状态码包（保持层级独立）。

### 2. 日志规范
- 使用注入的 `dep.AccessLogger` (请求日志) 和 `dep.OperationLogger` (操作日志)。
- 避免在循环中打日志。
- Error 日志应包含堆栈信息（Logger 库通常已封装）。

### 3. 事务处理
- 使用 `internal/biz/repo/transaction.go` 中定义的事务接口。
- 事务逻辑应控制在 Biz 层，通过 Closure (闭包) 传递给 Data 层执行。

### 4. 避免循环依赖
- `Biz` 层绝对不能 import `Service` 层。
- `Data` 层绝对不能 import `Service` 或 `Biz` 的具体实现（只能 import `Biz/Repo` 接口）。
- 如果出现循环依赖，通常意味着代码分层不清，请检查是否需要提取公共定义到 `dep` 或独立包。

### 5. Git 提交规范
- 遵循 Conventional Commits 规范 (e.g., `feat: add order module`, `fix: order status update bug`).
