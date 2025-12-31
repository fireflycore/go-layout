---
trigger: manual
---

# 错误处理规范（Go 微服务）v1.0
# ============================================
# 面向 go-layout 的错误处理标准与要求
# 通过将 [ENABLED] 更改为 [DISABLED] 来启用/禁用规则
#
# 最后更新：2025-12-31
# ============================================

## [规则 1] 错误分类体系 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
说明：
- 参数/校验错误：请求非法或缺少必需字段（Service 层优先由 protovalidate 发现）
- 业务错误：业务规则不满足（可恢复/可提示）
- 系统错误：DB/Redis/网络/配置等基础设施异常
- 第三方错误：调用外部服务失败（考虑降级/重试）

## [规则 2] 错误传播边界 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- Biz/Data 层返回 Go error，不依赖 gRPC status
- Service 层将 error 映射为响应体 `Code/Message`，并返回 `nil` error（与项目示例一致）
- 仅在“框架无法继续处理”的情况下返回非 nil error

## [规则 3] 错误包装与判断 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 包装错误使用 `%w`
- 分支判断使用 `errors.Is/As`
- 需要稳定语义时定义哨兵错误或轻量错误类型

## [规则 4] 日志与错误信息 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 记录错误时包含上下文，但避免泄露敏感信息
- 外部可见的 `Message` 使用可理解且不暴露内部细节的描述

