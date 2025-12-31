---
trigger: manual
---

# API 设计规范（go-layout，Proto/gRPC）v1.0

## [规则 1] Proto First [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
- 先改 Proto（通常在独立 Proto 仓库），再生成，再实现
- 生成产物以项目约定为准（本项目：`dep/protobuf/gen`）

## [规则 2] 入口统一做参数校验 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- Service 入口必须 `protovalidate.Validate(req)`
- 不重复实现与 Proto 同义的校验逻辑

## [规则 3] RPC 与消息命名 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- Service/RPC/Message 用 PascalCase
- RPC 用动词（Create/Get/Update/Delete/List）

## [规则 4] 列表分页字段一致 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
- `page`/`page_size` + `total` + `list`

## [规则 5] 向后兼容 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 不删除/复用字段号；弃用用 `deprecated` 并保留字段号

## [规则 6] 错误对外呈现 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- Biz/Data 返回 Go error
- Service 用响应体 `Code/Message` 表达失败并返回 `nil` error
