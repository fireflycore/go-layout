---
trigger: manual
---

# 测试规范（go-layout）v1.0

## [规则 1] 新逻辑必须可测 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 新增逻辑补充单元测试；修复缺陷补充回归用例

## [规则 2] 分层测试 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- Biz：以单元测试为主，mock Repo 接口
- Data：以集成测试为主（必要时引入测试库/容器）
- Service：关注入口校验与 `Code/Message` 映射

## [规则 3] 表驱动与子用例 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
- 使用表驱动覆盖正常/边界/异常路径
- 用 `t.Run` 给子用例起行为描述名

## [规则 4] Mock 只在边界 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
- 只 mock Repo/Remote 等边界依赖，避免绑定实现细节
