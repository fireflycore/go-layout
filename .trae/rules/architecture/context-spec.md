---
trigger: manual
---

# Context/Timeout 规范（go-layout）v1.1

## [规则 1] ctx 必须贯穿业务链路 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service/Biz/Data/Remote 调用必须使用入参 `ctx` 传递取消与链路信息
- Data 层所有 GORM/DB 调用必须使用 `db.WithContext(ctx)`
- 业务链路禁止使用 `context.Background/TODO` 替代入参 `ctx`

## [规则 2] 超时按场景派生 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 远程调用可在入参 `ctx` 基础上派生超时，避免无边界阻塞
- 后台任务允许 `context.Background()`，但必须设置超时与取消
