---
trigger: manual
---

# 测试规范（go-layout）v1.2

## [规则 1] 新逻辑必须可测 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 新增逻辑补充单元测试；修复缺陷补充回归用例
- 存量兼容：旧模块若无测试不强制补齐，但新增/改动逻辑需覆盖关键分支

## [规则 2] 分层测试 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Biz：以单元测试为主，mock Repo 接口
- Data：以集成测试为主（必要时引入测试库/容器）
- Service：关注入口校验与 `Code/Message` 映射

## [规则 3] 表驱动与子用例 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 使用表驱动覆盖正常/边界/异常路径
- 用 `t.Run` 给子用例起行为描述名

## [规则 4] Mock 只在边界 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 只 mock Repo/Remote 等边界依赖，避免绑定实现细节

## [规则 5] 新增公共工具必须有测试覆盖 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 新增 utils/中间件/通用转换等公共逻辑时，必须补充单元测试覆盖边界条件

## [规则 6] 交付前至少跑通最小回归集 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 验收按场景选择：仅规则/文档改动时，不强制跑 Go 回归；检查字数限制与引用一致性即可
- 涉及 Go 代码变更时：必须通过 `go test ./...`
- 建议同时通过 `go vet ./...`

按变更类型选择最小回归集：
- 仅改规则/文档：检查字数限制与引用一致性即可
- 新增/修改 Proto：Proto 仓库 `buf lint` → `buf push`；微服务仓库 `buf generate`（或 `make generate/make init`）并确保编译通过
- 新增/修改 Service RPC：`go test ./...`（覆盖入口校验与 Code/Message 映射的回归用例）
- 新增/修改 Biz 分支逻辑：`go test ./...`（至少覆盖关键分支的单测，mock Repo）
- 新增/修改 Data 查询/事务：`go test ./...`（必要时补集成测试或等价验证）
- 新增公共工具/中间件：`go test ./...`（覆盖边界条件）
