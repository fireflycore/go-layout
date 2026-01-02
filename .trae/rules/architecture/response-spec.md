---
trigger: manual
---

# API Response 规范（go-layout）v1.2

## [规则 1] 统一响应格式 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Proto
说明：
- 所有 Response 必须包含：`uint32 code = 1;`、`string message = 2;`
- 存量 Proto 可能存在 `message=1, code=2` 的历史写法；新写/新改按本规则逐步对齐
- 存量兼容：仅在改动到该 Response 时对齐字段顺序与 data 结构，避免顺手批量重排
- Code 语义：见 `architecture/status-code-spec.md`
- 返回值必须用 `data`（field=3）
  - 单值：`Type data = 3;`
  - 列表：`XxxList data = 3;`（`XxxList` 内部 `int64 total = 1; repeated Xxx list = 2;`）
- 无实际数据：`string data = 3;` 或 `optional string data = 3;`，且仅保留注释 `// 空值无用意`

## [规则 2] 列表分页字段一致 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- Request：`uint64 page = 1;`、`uint64 page_size = 2;`
- Response：`data.total = 1;`、`data.list = 2;`
