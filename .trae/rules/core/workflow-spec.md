---
trigger: manual
---

# 工作流程规范（go-layout）v1.4

## [规则 1] 生成链路以 makefile 为准 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 优先 `make init`；无 make 时按 makefile 等价命令执行

## [规则 2] Proto 协作流程 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 按仓库识别：存在 `buf.yaml/buf.work.yaml` 为 Proto 仓库；存在 `buf.gen.yaml` 为微服务仓库
- 在 Proto 仓库改 Proto：`buf lint` → `buf push`
- 在微服务仓库同步生成：`buf generate`（或 `make generate/make init`）

## [规则 3] 新增功能推进顺序 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 协议 → Service → Biz/Repo 接口 → Data 实现 → ProviderSet 注册 → wire 生成

## [规则 4] 交付前验证 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 验收按场景选择（见 `quality/testing-spec.md`）
