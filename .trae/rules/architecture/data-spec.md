---
trigger: manual
---

# Data/GORM 规范（go-layout）v1.1

## [规则 1] GORM 查询错误处理 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- gormx 封装 GORM，SQL 执行日志与异常可通过 OperationLogger 上报日志服务，无需业务层手工干预
- 普通查询（Find/Scan/Count）无要求则不抛出 error，直接返回结果
- 需要校验存在性（First/Take 且业务依赖 NotFound）才抛出 error

## [规则 2] 查询代码可读性 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 保持“变量定义 / 主逻辑 / 返回值”分块（空行分隔）
- 能直接赋值就不要用中间变量周转
