---
trigger: manual
---

# 错误处理规范（go-layout）v1.0

## [规则 1] 错误分类清晰 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 参数/校验错误（入口优先 protovalidate）
- 业务错误（可提示、可恢复）
- 系统错误（DB/Redis/网络/配置）
- 第三方错误（必要时重试/降级）

## [规则 2] 错误边界明确 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
- Biz/Data 返回 Go error
- Service 将 error 映射为响应体 `Code/Message` 并返回 `nil` error（模板风格）

## [规则 3] 错误包装与判断 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 包装用 `%w`；判断用 `errors.Is/As`

## [规则 4] 日志与对外信息 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 日志带上下文但不泄露敏感信息
- 对外 `Message` 不暴露内部细节

## [规则 5] panic 只允许在启动失败场景 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 业务链路（Service/Biz/Data）禁止 `panic`
- 仅允许在启动/装配失败且无法继续运行时 `panic`（例如配置加载失败）
