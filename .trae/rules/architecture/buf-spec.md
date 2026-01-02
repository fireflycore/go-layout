---
trigger: manual
---

# Buf/生成链路规范（go-layout）v1.3

## [规则 1] 生成配置以 buf.gen.yaml 为准 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- Proto 生成只通过 `buf generate`（或 `make generate/make init`）完成
- 存量兼容：不为“对齐结构”修改生成配置，除非需求明确要求

## [规则 2] Proto 仓库变更需发布 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- Proto 仓库新增/修改/删除 Proto：至少 `buf lint`
- 需要让下游仓库同步生成时：执行 `buf push` 发布到 BSR（或团队约定的 Registry）

## [规则 3] go_package 由 Buf 托管 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- managed 模式启用时，优先让 Buf 统一管理 `go_package_prefix`
- 不手动在业务 proto 里写 `option go_package` 来修补导入路径

## [规则 4] 生成物不可手改 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不修改 `dep/protobuf/gen/**` 下文件，变更通过修改 Proto/配置再重新生成

## [规则 5] 新增 Proto 模块的落点 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 新增依赖类型时，在 `buf.gen.yaml` 的 `inputs` 中补充 module/types
- 生成后代码路径固定：`dep/protobuf/gen/<package path>`
