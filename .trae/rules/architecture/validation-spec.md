---
trigger: manual
---

# 字段校验规范（go-layout）v1.2

## [规则 1] 字段约束按需添加 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Proto
说明：
- 字段约束不是必须的：用户未要求时，AI 根据场景判断是否需要
- 需要约束时推荐使用 `buf.validate`（protovalidate），避免在 Service/Biz 重复校验
- 可选字段优先用 `optional` 表达“缺省语义”，避免用零值承载“未传”
- 参考：`architecture/protovalidate-cheatsheet.md`
