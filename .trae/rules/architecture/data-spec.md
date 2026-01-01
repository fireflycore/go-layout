---
trigger: manual
---

# Data/GORM 规范（go-layout）v1.0

## [规则 1] GORM 查询错误处理 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 底层已封装 GORM，SQL 异常会统一上报 Logger
- 普通查询（Find/Scan/Count）无要求则不抛出 error，直接返回结果
- 需要校验存在性（First/Take 且业务依赖 NotFound）才抛出 error
- 存量兼容：仅在改动到该查询时套用本规则，避免顺手重写无关查询

## [规则 2] 查询代码可读性 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 保持“变量定义 / 主逻辑 / 返回值”分块（空行分隔）
- 能直接赋值就不要用中间变量周转
