---
trigger: manual
---

# API 设计规范（Proto/gRPC）v1.0
# ============================================
# 面向 go-layout（Proto First + gRPC）的 API 设计约束
# 通过将 [ENABLED] 更改为 [DISABLED] 来启用/禁用规则
#
# 最后更新：2025-12-31
# ============================================

## [规则 1] Proto First [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
说明：
- 以 Proto 定义驱动开发：先改 Proto（通常在独立 Proto 仓库），再生成，再实现
- 生成产物位置以项目约定为准（本项目为 `dep/protobuf/gen`）
- 不在 Go 代码中“手写”与 Proto 不一致的接口契约

## [规则 2] 请求参数验证（protovalidate）[ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 请求入口（Service 层）必须执行参数校验，优先使用已集成的 `protovalidate`
- 不重复实现与 Proto 规则同义的校验，避免规则漂移

## [规则 3] RPC 命名与消息结构 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- Service/RPC/Message 使用 PascalCase
- RPC 名称用动词短语（如 Create/Get/Update/Delete/List）
- 请求/响应消息显式区分（`CreateXxxRequest`/`CreateXxxResponse`）

## [规则 4] 分页与列表返回 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
说明：
- 列表接口统一使用 `page`/`page_size` 与 `total`、`list` 字段表达分页
- 过滤条件使用明确字段（如 `search_key`），避免“一字段承载多语义”

## [规则 5] 向后兼容 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 避免删除/复用字段号；需要弃用时使用 `deprecated` 语义并保留字段号
- 新增字段应提供合理默认值语义，避免破坏旧客户端

## [规则 6] 错误对外呈现与响应风格 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 与项目示例一致：Service 层通常通过响应体 `Code/Message` 表达业务失败，返回 `nil` error
- Biz/Data 返回 Go error；Service 将 error 映射为响应体字段
- 不在 `Message` 里暴露敏感信息

