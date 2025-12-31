---
trigger: manual
---

# 工作流程规范 v1.0
# ============================================
# AI 辅助编码的开发工作流程和流程规则
# 通过将 [ENABLED] 更改为 [DISABLED] 来启用/禁用规则
#
# 使用方法：
# 1. 在 AI 对话中引用此文件
# 2. 根据项目需求启用/禁用规则
# 3. AI 将只遵循 ENABLED 的规则
#
# 最后更新：2025-12-31
# ============================================

## [规则 1] 变更日志管理 [ENABLED]
# 在变更日志文件中维护更新记录

STATUS: ENABLED
说明：
- 在 CHANGELOG.md 或类似文件中记录所有重要变更
- 包含版本号、日期和变更描述
- 在提交代码变更前更新变更日志
- 遵循 "Keep a Changelog" 格式（Added/Changed/Deprecated/Removed/Fixed/Security）

## [规则 2] 版本号管理 [ENABLED]
# 在项目中一致地更新版本号

STATUS: ENABLED
说明：
- 遵循语义化版本控制（MAJOR.MINOR.PATCH）
- 发布时更新承载版本信息的位置（本项目默认在 `conf/bootstrap.json` 的 `version` 字段）
- MAJOR：破坏性变更
- MINOR：新功能（向后兼容）
- PATCH：错误修复（向后兼容）

## [规则 3] Git 提交信息格式 [DISABLED]
# 标准化提交信息结构（Conventional Commits）

STATUS: DISABLED

## [规则 4] 分支命名约定 [DISABLED]
# 一致的分支命名策略

STATUS: DISABLED

## [规则 5] 代码审查要求 [DISABLED]
# Pull Request 和审查指南

STATUS: DISABLED

## [规则 6] 文档同步 [ENABLED]
# 保持文档与代码变更同步

STATUS: ENABLED
说明：
- 修改分层边界、依赖注入、配置加载方式、生成链路时更新相关文档
- 更改 gRPC/Proto 接口时更新 Proto 定义（通常在独立 Proto 仓库），并同步生成产物（`buf generate` -> `dep/protobuf`）
- 增加/更改 Make 目标或生成脚本时同步更新 `docs/project-guide.md`
- 破坏性变更需更新 `docs/` 中对应章节，并在变更日志中显式标注

## [规则 7] 测试覆盖率要求 [DISABLED]
# 新代码的测试标准

STATUS: DISABLED

## [规则 8] 部署前检查清单 [DISABLED]
# 部署前的验证步骤

STATUS: DISABLED

## [规则 9] 破坏性变更协议 [ENABLED]
# 如何处理破坏性变更

STATUS: ENABLED
说明：
- 清楚地记录所有破坏性变更
- 为使用方提供迁移指导（必要时包含代码示例）
- 提升 MAJOR 版本号
- 在移除前标记弃用窗口（如适用）

## [规则 10] 依赖更新策略 [ENABLED]
# 管理第三方依赖

STATUS: ENABLED
说明：
- Go 模块变更以 `go.mod/go.sum` 为准；更新后必须执行 `go mod tidy`
- 依赖更新优先跟随上游发布说明与安全公告
- 避免在自动化脚本中滥用 `@latest`；明确版本边界与回归验证
- 本项目存在集中更新脚本时，保持脚本与 go.mod 的一致性

## [规则 11] 文件组织标准 [DISABLED]
# 项目结构和文件放置

STATUS: DISABLED
说明：
- 遵循 `docs/directory-structure.md` 描述的 go-layout 结构
- 将 Biz/Data/Service/Server 的职责与目录保持一致
- 不直接在 `internal` 外暴露业务实现（避免外部 import）

## [规则 12] 错误处理标准 [ENABLED]
# 一致的错误处理方法

STATUS: ENABLED
说明：
- 始终处理错误，绝不静默失败
- 在外部边界（RPC/DB/Redis/IO）明确检查并返回 error
- 记录错误时包含足够上下文，且避免泄露敏感信息
- Biz/Data 层返回 Go error；Service 层按项目约定通过响应体 `Code/Message` 表达业务失败并返回 `nil` error
- 需要分支处理时使用 `errors.Is/As`，必要时定义轻量错误类型

# ============================================
# 摘要 - 仅启用的规则
# ============================================

✅ [规则 1]  变更日志管理
✅ [规则 2]  版本号管理
✅ [规则 6]  文档同步
✅ [规则 9]  破坏性变更协议
✅ [规则 10] 依赖更新策略
✅ [规则 12] 错误处理标准

