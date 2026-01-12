---
trigger: manual
---

# 开发需求规范（go-layout）v1.3

## [规则 1] 交付必须可用 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 验收按场景选择（见 `quality/testing-spec.md`）
- 涉及 Go 代码变更时：必须通过 `go test ./...`
- 不交付占位符/半实现

交付触发清单（硬性）：
- 新增/修改对外 RPC：必须补入口校验与错误映射；必要时补 protovalidate 规则；至少补 1 个回归用例
- 新增/修改 Biz 分支逻辑：必须补单元测试覆盖关键分支（mock Repo）
- 新增/修改 Data 查询条件：必须明确 NotFound 语义；涉及事务/一致性时补集成测试或等价验证
- 新增公共工具/中间件：必须补单元测试覆盖边界条件

最小交付切片（用于缩小改动面）：
- 新增/改动 RPC：只改本 RPC 的 Service + 对应 UseCase/Repo + 对应 Data
- 只改 Biz：不动 Proto/生成物；只调整用例与其依赖边界
- 只改 Data：不动 Proto/Service；只调整查询/事务与其返回语义

## [规则 2] 严守分层依赖 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- Service：入口编排（校验/metadata/调用 Biz）
- Biz：只依赖 Repo 接口
- Data：实现 Repo，不反向依赖 Biz/Service

## [规则 3] 以仓库现状为准 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 路径/命令/生成物位置先以 `makefile`、`go.mod`、现有代码为准

## [规则 4] 生成代码不可手改 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不修改 `wire_gen.go`、`*_gen.go`、`dep/` 下内容

## [规则 5] 变更范围最小 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 只改与需求直接相关的代码
- 方法替换/重命名：直接删除旧方法并同步更新所有调用方，不保留兼容层

## [规则 6] 禁止硬编码敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 真实密钥/令牌/密码/个人敏感信息不得进入代码与仓库
