---
trigger: manual
---

# 开发需求规范（go-layout）v1.0

## [规则 1] 产物必须可编译可运行 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 交付前必须能通过 `go test ./...`
- 不交付占位符/半实现

## [规则 2] 不发明接口、类型与依赖 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- gRPC/Proto 类型以 `dep/protobuf/gen` 生成产物为准
- 新增依赖前先确认 go.mod 现有依赖；新增后执行 `go mod tidy`

## [规则 3] 严格遵守分层与依赖方向 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- `internal/service` 只编排入口：校验、提取 metadata、调用 Biz
- `internal/biz` 只依赖 `internal/biz/repo` 接口，不依赖 data/service 实现
- `internal/data` 实现 repo 接口，不反向依赖 biz/service

## [规则 4] 参数校验复用 protovalidate [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service 入口统一 `protovalidate.Validate(req)`，避免重复校验逻辑漂移

## [规则 5] 不手改生成代码 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不修改 `wire_gen.go`、`*_gen.go` 与 `dep/` 下生成内容
- 通过 `make dto` / `make init`（buf/goverter/wire）更新生成产物

## [规则 6] 修改范围最小化 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 只改与需求直接相关的文件与逻辑，避免顺手重构

## [规则 7] DTO/PO/DO 字段类型必须一致 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 同一业务字段在 `pb`（DTO）、`internal/data/entity`（PO）、`internal/biz/model`（DO）中的类型必须对齐
- 允许 Data 直接返回 `pb` 以优化读性能，但要保证字段映射可控且可验证
- 涉及指针/零值语义差异时，必须显式处理 `nil` 与默认值，避免隐式丢失信息

## [规则 8] 依赖注入必须通过 Wire ProviderSet [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 各层依赖以 `internal/*/core.go` 的 `ProviderSet` 暴露与装配为准
- 禁止通过包级全局变量持有 DB/Redis/Etcd/Logger/ClientConn 等长生命周期资源

## [规则 9] 上下文必须贯穿全链路 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 所有跨层方法都以 `context.Context` 作为首参，并从入口向下传递
- 禁止在业务链路中使用 `context.Background/TODO` 替代入参 ctx

## [规则 10] 代码格式化与导入必须规范 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 交付代码必须通过 `gofmt`，导入分组保持最小且一致

## [规则 11] 配置与密钥不得硬编码入代码 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 运行配置来自 `conf/bootstrap.json` 或远程配置服务加载结果
- 禁止在代码中写死密钥/令牌/密码；日志与响应体不得输出敏感信息
