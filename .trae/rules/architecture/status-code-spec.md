---
trigger: manual
---

# 响应码语义规范（go-layout）v1.3

## [规则 1] Code 是业务码，语义稳定 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- Code 表达业务结果，不等同于 gRPC status，也不要求与 HTTP 状态码一一对应

## [规则 2] 推荐的最小码表 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 200：成功（默认 `Message=success`）
- 400：参数/业务校验失败（Message 允许返回可读原因）
- 404：资源不存在（如按 id 查询未找到）
- 500：系统/第三方依赖失败（模板兼容允许直接透出 `err.Error()`）

## [规则 3] 错误映射落地规则 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- protovalidate 校验失败：Code=400，Message=err.Error()
- 用户上下文解析失败：默认 Code=400（模板兼容）；若服务已区分鉴权语义可用 401/403
- Biz 可预期错误：默认 Code=400，Message 使用安全可读的业务文案
- Data/Remote 系统错误：Code=500；模板兼容允许直接透出 `err.Error()`；并把真实 err 写入日志
