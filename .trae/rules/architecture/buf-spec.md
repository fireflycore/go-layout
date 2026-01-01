---
trigger: manual
---

# Buf/生成链路规范（Proto）v1.0

## [规则 1] 生成配置以 buf.gen.yaml 为准 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- Proto 生成只通过 `buf generate`（或 `make generate/make init`）完成
- 存量兼容：不为“对齐结构”修改生成配置，除非需求明确要求

## [规则 2] go_package 由 Buf 托管 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- managed 模式启用时，优先让 Buf 统一管理 `go_package_prefix`
- 不手动在业务 proto 里写 `option go_package` 来修补导入路径

## [规则 3] 生成物不可手改 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不修改 `dep/protobuf/gen/**` 下文件，变更通过修改 Proto/配置再重新生成

## [规则 4] 新增 Proto 模块的落点 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 新增依赖类型时，在 `buf.gen.yaml` 的 `inputs` 中补充 module/types
- 生成后代码路径固定：`dep/protobuf/gen/<package path>`

