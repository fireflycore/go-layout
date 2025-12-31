---
trigger: manual
---

# 工作流程规范（go-layout）v1.0

## [规则 1] 生成链路必须一致 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 以根目录 `makefile` 为准，优先使用 `make init`
- 等价链路：`buf generate` → `goverter gen ./internal/biz/convert` → `wire ./cmd/server` → `go mod tidy`

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

## [规则 6] 新增功能按“协议→入口→业务→数据→装配”推进 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 先更新 Proto（通常在独立 Proto 仓库）并同步 `dep/protobuf/gen`
- 再实现 `internal/service`（校验/metadata/调用 Biz）与 `internal/biz`（UseCase + repo 接口）
- 最后实现 `internal/data`（repo 实现与持久化）并把构造函数加入对应层 `core.go` 的 `ProviderSet`

## [规则 7] 新增构造函数必须注册并重新生成 Wire [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 新增/变更 Provider 后执行 `wire ./cmd/server`（或 `make init`）确保依赖图最新

## [规则 8] 提交前最小回归集 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 必须通过 `go test ./...` 与 `go vet ./...`
- 改动涉及生成链路时，必须保证生成产物与引用路径一致（`dep/protobuf/gen`、`dep/dto`、`wire_gen.go`）
