---
trigger: manual
---

# 规则索引（go-layout）v1.0

GLOBAL:
  DEFAULT_PROFILE: GoService
  ENABLE_MODULES:
    core/requirements-spec.md: ENABLED
    core/workflow-spec.md: ENABLED
    core/naming-conventions.md: ENABLED
    architecture/api-design-spec.md: ENABLED
    quality/security-spec.md: ENABLED
    quality/error-handling-spec.md: ENABLED
    quality/testing-spec.md: ENABLED

MODULE: core/requirements-spec.md
  STATUS: ENABLED
  VERSION: v1.0

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

PROFILE: GoService
  CORE:
    - core/requirements-spec.md
    - core/workflow-spec.md
    - core/naming-conventions.md
  OPTIONAL:
    - architecture/api-design-spec.md
    - quality/security-spec.md
    - quality/error-handling-spec.md
    - quality/testing-spec.md

USAGE:
  - 优先引用：@.trae/rules/core/spec-index.md
  - 需要更严格时再补充引用 OPTIONAL 模块
