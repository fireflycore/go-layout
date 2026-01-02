---
trigger: manual
---

# 外部文档 URL 规范（go-layout）v1.0

## [规则 1] 外部文档使用固定入口 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 需要引用/查阅外部文档时，优先使用本规则列出的固定入口，避免搜索到旧版或非官方页面
- 当规则或代码中出现同类概念（例如 Proto 校验、Buf 命令、代码生成器），链接以本清单为准

## [规则 2] 常用文档地址清单 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- Protovalidate（buf.validate）字段校验规则：
  - https://buf.build/bufbuild/protovalidate/docs/main:buf.validate
- Buf CLI 官方文档：
  - https://buf.build/docs/
- Goverter 官方文档：
  - https://goverter.jmattheis.de

