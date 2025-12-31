---
trigger: manual
---

# 命名约定（go-layout）v1.0

## [约定 1] 变量与字段 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 本地变量与非导出字段用 lowerCamelCase
- 布尔变量用 is/has/can/should 前缀
- 常用：`ctx`、`req`/`request`、`res`/`result`、`err`

## [约定 2] 函数/类型/接口 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 导出标识符用 PascalCase；非导出用 lowerCamelCase
- 构造函数以 `New` 开头
- 分层后缀保持一致：`UseCase` / `Repo` / `Service`

## [约定 3] 包与文件 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 包名全小写、短、无下划线
- 文件名全小写；多个词使用下划线（如 `wire_gen.go`）

## [约定 4] gRPC/Proto [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service/RPC/Message 用 PascalCase；Service 方法名与 RPC 一致
- Service 入口用 `protovalidate.Validate(req)` 校验

## [约定 5] GORM/JSON [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- Go 字段 PascalCase；JSON tag 使用 snake_case（与接口返回对齐）

## [约定 6] 生成代码 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不修改 `wire_gen.go`、`*_gen.go` 与 `dep/` 下生成内容
