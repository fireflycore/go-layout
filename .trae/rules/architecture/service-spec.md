---
trigger: manual
---

# Service 入口规范（go-layout）v1.1

## [规则 1] 入口统一参数校验 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service 入口统一 `protovalidate.Validate(req)`
- 不重复实现与 Proto 同义的校验逻辑

## [规则 2] 用户上下文只解析一次 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service 从入站 metadata 解析（对齐 `micro.ParseUserContextMeta`）
- Biz/Data 不依赖 metadata 类型

## [规则 3] 入口处理骨架一致 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 先初始化成功响应（`Code=200/Message=success`）
- 失败时只改 `Code/Message` 并返回响应体，error 返回值保持 `nil`
- 存量兼容：仅在该 RPC 需要改动时按骨架对齐，避免顺手批量重排
