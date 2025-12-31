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
#    "@.trae/rules/core/spec-index.md @.trae/rules/core/requirements-spec.md @.trae/rules/core/workflow-spec.md"
# 3. 设置 GLOBAL 开关与 PROFILE；用于说明在不同项目类型下推荐启用哪些规则。
#
# 最后更新：2025-12-31
# ============================================

GLOBAL:
  DEFAULT_PROFILE: GoService      # Options: GoService | CLI | Library
  ENABLE_MODULES:
    core/requirements-spec.md: ENABLED
    core/workflow-spec.md: ENABLED
    core/naming-conventions.md: ENABLED
    architecture/api-design-spec.md: ENABLED
    quality/security-spec.md: ENABLED
    quality/error-handling-spec.md: ENABLED
    quality/testing-spec.md: ENABLED
  LANGUAGE_PAIRS: ENABLED

MODULE: core/requirements-spec.md
  STATUS: ENABLED
  VERSION: v1.1

MODULE: core/workflow-spec.md
  STATUS: ENABLED
  VERSION: v1.0

MODULE: core/naming-conventions.md
  STATUS: ENABLED
  VERSION: v1.0

MODULE: architecture/api-design-spec.md
  STATUS: ENABLED
  VERSION: v1.0

MODULE: quality/security-spec.md
  STATUS: ENABLED
  VERSION: v1.0

MODULE: quality/error-handling-spec.md
  STATUS: ENABLED
  VERSION: v1.0

MODULE: quality/testing-spec.md
  STATUS: ENABLED
  VERSION: v1.0

DEPENDENCIES:
  core/requirements-spec.md::RULE 1 -> core/requirements-spec.md::RULE 10
    note: 完整性依赖于成功编译
  core/requirements-spec.md::RULE 6 -> architecture/api-design-spec.md::RULE 1
    note: API 设计需要与现有 Proto/生成产物一致

CONFLICTS:
  CASE: core/workflow-spec.md::RULE 6 vs core/requirements-spec.md::RULE 5
  RESOLUTION:
    - 若代码改动影响公共 API 面或用户可见行为，则必须同步文档
    - 其他情况优先最小化修改；文档调整安排到后续文档任务

ACTIVE:
  PROFILE: GoService
  MODULES: core/requirements-spec.md (ENABLED), core/workflow-spec.md (ENABLED), core/naming-conventions.md (ENABLED), architecture/api-design-spec.md (ENABLED), quality/security-spec.md (ENABLED), quality/error-handling-spec.md (ENABLED), quality/testing-spec.md (ENABLED)

