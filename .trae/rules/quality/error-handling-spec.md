---
trigger: manual
---

# 错误处理规范（go-layout）v1.3

## [规则 1] 错误边界明确 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- Biz/Data 返回 Go error
- Service 用响应体 `Code/Message` 表达失败并返回 `nil` error（保持模板一致）
- Code 语义与最小码表：见 `architecture/status-code-spec.md`

## [规则 2] 对外信息可控 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 模板兼容：允许直接透出 `err.Error()` 到 `Message`（含系统/第三方错误）
- 真实错误允许直接写入日志，不做兜底改写
- 对外文案收敛与错误分类映射暂不强制

## [规则 3] 错误包装与判断 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 现阶段不强制错误包装；如需包装用 `%w`，判断用 `errors.Is/As`

## [规则 4] 业务链路禁止 panic [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service/Biz/Data 禁止 `panic`（仅启动失败可 panic）
