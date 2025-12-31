---
trigger: manual
---

# 错误处理规范（go-layout）v1.0

## [规则 1] 错误分类清晰 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 参数/校验错误（入口优先 protovalidate）
- 业务错误（可提示、可恢复）
- 系统错误（DB/Redis/网络/配置）
- 第三方错误（必要时重试/降级）

## [规则 2] 错误边界明确 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- Biz/Data 返回 Go error
- Service 默认将 error 映射为响应体 `Code/Message` 并返回 `nil` error（模板风格）
- 若采用 gRPC status/拦截器统一错误语义，则保持一致，避免同一服务内混用

## [规则 3] 错误包装与判断 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 包装用 `%w`；判断用 `errors.Is/As`

## [规则 4] 日志与对外信息 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 日志带上下文但不泄露敏感信息
- 对外 `Message` 不暴露内部细节

## [规则 5] panic 只允许在启动失败场景 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 业务链路（Service/Biz/Data）禁止 `panic`
- 仅允许在启动/装配失败且无法继续运行时 `panic`（例如配置加载失败）

## [规则 6] 对外错误信息可控 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 可预期的业务错误可对外透出简明信息
- 系统/第三方错误优先记录日志，对外返回通用信息，避免把内部细节透出到 `Message`
- 不确定错误内容是否安全时，优先使用通用 `Message`
