---
trigger: manual
---

# Proto 规范（go-layout × Proto 仓库）v1.4

## [规则 1] Proto First [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Proto
说明：
- 先改 Proto（通常在独立 Proto 仓库），再生成（`dep/protobuf/gen`），再实现
- Proto 仓库提交前：至少 `buf lint`；需要对外发布供下游生成时：`buf push`
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
