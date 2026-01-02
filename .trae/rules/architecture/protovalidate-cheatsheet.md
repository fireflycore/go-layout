---
trigger: manual
---

# Protovalidate 速查（buf.validate）v1.0

适用范围：
- Proto 字段级约束，优先用 `buf.validate` 表达
- 校验落地：Service 入口调用 `protovalidate.Validate(req)`

常用写法：
- 必填：`[(buf.validate.field).required = true]`
- UUID：`[(buf.validate.field).string.uuid = true]`
- 字符串长度：
  - `[(buf.validate.field).string = { min_len: 1, max_len: 64 }]`
- 枚举/数值取值集合：
  - `[(buf.validate.field).uint32 = { in: [0, 1, 2] }]`
- 数值范围（示意，按实际类型选择 int32/int64/uint32/uint64 等）：
  - `[(buf.validate.field).int64 = { gte: 0, lte: 1000 }]`

关于 optional：
- 需要区分“未传”与“传了零值”时，用 `optional` 表达缺省语义
- optional 常与校验搭配使用；不要用零值去承载“未传”的业务语义
