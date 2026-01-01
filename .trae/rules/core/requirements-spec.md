---
trigger: manual
---

# 开发需求规范（go-layout）v1.1

## [规则 1] 交付必须可用 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 验收按场景选择（见 `quality/testing-spec.md`）
- 涉及 Go 代码变更时：必须通过 `go test ./...`
- 不交付占位符/半实现

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

## [规则 6] 禁止硬编码敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 密钥/令牌/密码不得进入代码、日志、响应体
