---
trigger: manual
---

# 规范索引 v1.0
# ============================================
# 规范套件的中心控制文件。
# 管理模块、全局开关、规则依赖、冲突与项目类型配置。
#
# 使用方法：
# 1. 在 AI 对话中引用本索引，并按需引用具体模块。
# 2. 在 AI 对话中与模块一起引用，例如：
#    "@.trae/rules/core/spec-index.zh-CN.md @.trae/rules/core/requirements-spec.zh-CN.md @.trae/rules/core/workflow-spec.zh-CN.md"
# 3. 设置 GLOBAL 开关与 PROFILE；用于说明在不同项目类型下推荐启用哪些规则。
#
# 最后更新：2025-12-31
# ============================================

# ============================================
# GLOBAL CONFIG
# ============================================
GLOBAL:
  DEFAULT_PROFILE: GoService      # Options: GoService | CLI | Library
  ENABLE_MODULES:
    core/requirements-spec.zh-CN.md: ENABLED
    core/workflow-spec.zh-CN.md: ENABLED
    core/naming-conventions.zh-CN.md: ENABLED
    architecture/api-design-spec.zh-CN.md: ENABLED
    quality/security-spec.zh-CN.md: ENABLED
    quality/error-handling-spec.zh-CN.md: ENABLED
    quality/testing-spec.zh-CN.md: ENABLED
  LANGUAGE_PAIRS: ENABLED         # 需要时同时参考英文版本

# ============================================
# MODULES
# ============================================
MODULE: core/requirements-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.1
  SUMMARY:
    - 13 条通用编码规则（CRITICAL/HIGH/MEDIUM）
    - 关注点：完整性、复用、最小依赖、正确性、可构建、与项目现有结构一致
  TOP_PRIORITY:
    - RULE 1: Generate complete, runnable code (CRITICAL)
    - RULE 6: Verify all APIs exist (CRITICAL)
    - RULE 10: Ensure code compiles successfully (CRITICAL)
    - RULE 13: Use only real, existing libraries (CRITICAL)

MODULE: core/workflow-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - 12 条工作流规则，支持 ENABLED/DISABLED 切换
    - 关注点：最小化改动、生成链路（buf/goverter/wire）、文档同步、依赖更新、错误处理
  TOP_PRIORITY:
    - RULE 1: Change Log Management (ENABLED)
    - RULE 2: Version Number Management (ENABLED)
    - RULE 6: Documentation Sync (ENABLED)
    - RULE 9: Breaking Changes Protocol (ENABLED)
    - RULE 10: Dependency Update Policy (ENABLED)
    - RULE 12: Error Handling Standards (ENABLED)

MODULE: core/naming-conventions.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - 12 条命名约定，支持 ENABLED/DISABLED 切换
    - 默认启用：变量、函数、类型、常量、文件、环境变量命名（含 Go 约定）

MODULE: architecture/api-design-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - Proto First 与 gRPC API 设计约束
    - 关注点：protovalidate、分页与查询、向后兼容、响应风格一致性

MODULE: quality/security-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - 安全基线与敏感信息处理约束

MODULE: quality/error-handling-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - Go 错误建模、包装、日志与对外呈现策略

MODULE: quality/testing-spec.zh-CN.md
  STATUS: ENABLED
  VERSION: v1.0
  SUMMARY:
    - Go 测试约束：单元/集成分层、表驱动、mock 边界、回归测试

# ============================================
# RULE DEPENDENCIES (auto-enable dependents)
# ============================================
DEPENDENCIES:
  # requirements-spec
  core/requirements-spec.zh-CN.md::RULE 1 -> core/requirements-spec.zh-CN.md::RULE 10
    note: 完整性依赖于成功编译
  core/requirements-spec.zh-CN.md::RULE 6 -> architecture/api-design-spec.zh-CN.md::RULE 1
    note: API 设计需要与现有 Proto/生成产物一致

  # workflow-spec
  core/workflow-spec.zh-CN.md::RULE 9 -> core/workflow-spec.zh-CN.md::RULE 2
    note: 破坏性变更需正确提升 MAJOR 版本
  core/workflow-spec.zh-CN.md::RULE 9 -> core/workflow-spec.zh-CN.md::RULE 1
    note: 破坏性变更必须记录到变更日志

  # quality
  quality/error-handling-spec.zh-CN.md::RULE 3 -> quality/security-spec.zh-CN.md::RULE 6
    note: 日志必须避免敏感信息泄露

# ============================================
# RULE CONFLICTS & RESOLUTION
# ============================================
CONFLICTS:
  CASE: core/workflow-spec.zh-CN.md::RULE 6 (Documentation Sync) vs core/requirements-spec.zh-CN.md::RULE 5 (Only Requested Changes)
  RISK: 同步文档可能扩大修改范围，与最小化修改策略冲突
  RESOLUTION:
    - 若代码改动影响公共 API 面或用户可见行为，则 Documentation Sync 为必需。
    - 其他情况优先最小化修改；文档更新安排到后续文档任务。
    - 优先级顺序：CRITICAL > HIGH > MEDIUM

  CASE: dependency updates (core/workflow-spec.zh-CN.md::RULE 10) vs minimal deps (core/requirements-spec.zh-CN.md::RULE 3)
  RESOLUTION:
    - 安全补丁与关键修复优先于最小依赖策略。
    - 非关键更新应尽量减少依赖影响。

MODULE_PRECEDENCE:
  - 代码正确性与可运行性（requirements-spec）在生成输出时优先。
  - 流程合规（workflow-spec）在发布治理中优先。
  - 命名（naming-conventions）在不影响代码正确性时适用。

# ============================================
# PROJECT PROFILES (recommended enables)
# ============================================
PROFILE: GoService
  REQUIREMENTS:
    ENABLE: [RULE 1, 2, 3, 5, 6, 7, 10, 12, 13]
    OPTIONAL: [RULE 8, 9, 11]
  WORKFLOW:
    ENABLE: [RULE 2, 6, 9, 10, 12]
    OPTIONAL: [RULE 1, 3, 4, 5, 7, 8, 11]
  NAMING:
    ENABLE: [CONVENTION 1, 2, 3, 4, 5, 9]
    OPTIONAL: [CONVENTION 6, 7, 8, 10, 11, 12]

PROFILE: Web
  REQUIREMENTS:
    ENABLE: [RULE 1, 2, 3, 5, 6, 7, 10, 11, 12, 13]
    OPTIONAL: [RULE 8, 9]
  WORKFLOW:
    ENABLE: [RULE 1, 2, 6, 9, 10, 12]
    OPTIONAL: [RULE 3, 4, 5, 7, 8, 11]
  NAMING:
    ENABLE: [CONVENTION 1, 2, 3, 4, 5, 6, 9]
    OPTIONAL: [CONVENTION 7, 8, 10, 11, 12]

PROFILE: CLI
  REQUIREMENTS:
    ENABLE: [RULE 1, 2, 3, 5, 6, 7, 10, 12, 13]
    OPTIONAL: [RULE 8, 9, 11]
  WORKFLOW:
    ENABLE: [RULE 1, 2, 6, 10, 12]
    OPTIONAL: [RULE 3, 4, 5, 7, 8, 9, 11]
  NAMING:
    ENABLE: [CONVENTION 1, 2, 3, 4, 5, 9]
    OPTIONAL: [CONVENTION 6, 7, 8, 10, 11, 12]

PROFILE: Library
  REQUIREMENTS:
    ENABLE: [RULE 1, 2, 3, 6, 7, 10, 12, 13]
    OPTIONAL: [RULE 5, 8, 9, 11]
  WORKFLOW:
    ENABLE: [RULE 1, 2, 9, 10, 12]
    OPTIONAL: [RULE 3, 4, 5, 6, 7, 8, 11]
  NAMING:
    ENABLE: [CONVENTION 1, 2, 3, 4, 5, 10, 12]
    OPTIONAL: [CONVENTION 6, 7, 8, 9, 11]

# ============================================
# OVERRIDES (optional per-project switches)
# ============================================
# 用于在不编辑模块文件的情况下覆盖模块内状态。
# 示例语法：
OVERRIDES:
  core/requirements-spec.zh-CN.md:
    DISABLE: [RULE 8]        # 速度优先时临时关闭注释一致性
    ENABLE:  [RULE 9]        # 明确启用“功能优先”
  core/workflow-spec.zh-CN.md:
    ENABLE:  [RULE 3]        # 采用 Conventional Commits
    DISABLE: [RULE 8]        # 非生产环境跳过部署前检查清单
  core/naming-conventions.zh-CN.md:
    ENABLE:  [CONVENTION 7]  # 后端服务开启数据库命名约定

# ============================================
# SUMMARY
# ============================================
ACTIVE:
  PROFILE: GoService
  MODULES: core/requirements-spec.zh-CN.md (ENABLED), core/workflow-spec.zh-CN.md (ENABLED), core/naming-conventions.zh-CN.md (ENABLED), architecture/api-design-spec.zh-CN.md (ENABLED), quality/security-spec.zh-CN.md (ENABLED), quality/error-handling-spec.zh-CN.md (ENABLED), quality/testing-spec.zh-CN.md (ENABLED)

# ============================================
# Version History
# ============================================
# v1.0 (2025-11-09) - 首个中心索引：含全局开关、依赖、冲突与项目配置
# ============================================
