---
trigger: manual
---

# 规则索引（go-layout）v1.0

GLOBAL:
  DEFAULT_PROFILE: GoService
  LOADING_POLICY:
    - 规则按 PROFILE 加载；GoService 仅加载 CORE，按需加载 OPTIONAL
  PROTO_REPO:
    SOURCE_OF_TRUTH: proto_repo
    REQUIRED_FOR_PROTO_CHANGES: true
    AUTO_DISCOVERY:
      ENABLED: true
      SEARCH_MARKERS:
        - buf.yaml
        - buf.work.yaml
      PREFERRED_DIR_NAMES:
        - demo-proto
    LOCAL_PATH_FALLBACK: /Users/lhdht/product/lhdht/code/firefly
  RULE_PRECEDENCE:
    - core/requirements-spec.md
    - core/workflow-spec.md
    - core/naming-conventions.md
    - architecture/api-design-spec.md
    - quality/error-handling-spec.md
    - quality/security-spec.md
    - quality/testing-spec.md
  REALITY_SOURCES:
    - makefile
    - go.mod
    - internal/**
    - dep/**
  PROTO_COLLAB_FLOW:
    - 修改 Proto：优先从工作区自动发现 Proto 仓库（marker：`buf.yaml`/`buf.work.yaml`），必要时使用 `LOCAL_PATH_FALLBACK`
    - 校验并发布：在 Proto 仓库执行 `buf lint` 与 `buf push`
    - 同步生成：在本仓库执行 `buf generate`（或 `make generate`/`make init`）
  DELIVERY_CHECKS:
    - go test ./...
    - go vet ./...
  OUTPUT_REQUIREMENTS:
    - 变更文件清单
    - 验证命令与结果
    - 假设与风险提示
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

PROFILE: GoServiceStrict
  CORE:
    - core/requirements-spec.md
    - core/workflow-spec.md
    - core/naming-conventions.md
    - architecture/api-design-spec.md
    - quality/security-spec.md
    - quality/error-handling-spec.md
    - quality/testing-spec.md

USAGE:
  - 优先引用：@.trae/rules/core/spec-index.md
  - 默认按 PROFILE: GoService 执行，仅在需要时加载 OPTIONAL 模块
  - 需要更严格时使用 PROFILE: GoServiceStrict
