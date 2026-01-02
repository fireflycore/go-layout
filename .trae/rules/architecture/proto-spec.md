---
trigger: manual
---

# Proto 规范（go-layout × Proto 仓库）v1.2

## [规则 1] Proto First [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Proto
说明：
- 先改 Proto（通常在 Proto 独立仓库，常见为 firefly），再生成（`dep/protobuf/gen`），再实现
- 存量兼容：仅在改动到该 proto 文件时按规则对齐，避免顺手重排字段/响应结构
- 生成链路：见 `architecture/buf-spec.md`

## [规则 2] 命名一致 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- Service/RPC/Message 用 PascalCase；RPC 用动词（Create/Get/Update/Delete/List）

## [规则 3] 注释规范 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- 见 `architecture/comment-spec.md`

## [规则 4] 向后兼容 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- 不删除/复用字段号；弃用用 `deprecated` 并保留字段号
