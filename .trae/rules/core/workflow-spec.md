---
trigger: manual
---

# 工作流程规范（go-layout）v1.0

## [规则 1] 生成链路必须一致 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 更新链路：`buf generate` → `make dto` → `wire ./cmd/server` → `go mod tidy`
- 优先使用 `make init`（等价于上述链路）

## [规则 2] 改动 Proto 必须同步生成产物 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Proto 变更后同步更新 `dep/protobuf/gen`（本项目不在仓库内维护 `.proto`）

## [规则 3] 修改 Converter 必须重新生成 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 修改 `internal/biz/convert` 接口后运行 `make dto`

## [规则 4] 依赖变更必须整理与验证 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 依赖变更后执行 `go mod tidy`
- 交付前执行 `go test ./...` 与 `go vet ./...`

## [规则 5] 版本号更新位置 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 需要对外发布时更新 `conf/bootstrap.json` 的 `version`
