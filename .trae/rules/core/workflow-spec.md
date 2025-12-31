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
- 若本机环境无 `make`，按 `makefile` 的目标内容逐条执行等价命令
- 常用目标语义：
  - `make dto`：仅生成转换代码（goverter）
  - `make generate`：生成 Proto 代码（buf）并生成 DTO（goverter）
  - `make init`：generate + wire + go mod tidy
  - `make run`：init + go run

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

## [规则 9] 交付输出必须包含固定要素 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 变更文件清单（路径 + 简述）
- 运行的验证命令与结果（至少 `go test ./...`、`go vet ./...`）
- 若未运行命令，必须说明原因与建议运行方式
- 必要的假设与风险提示（如环境差异、工具缺失、外部依赖不可用）

## [规则 10] Buf-CLI 协同管理 Proto [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- Proto 的修改/新增/删除在独立 Proto 仓库完成（优先从工作区自动发现），本仓库只通过 `buf generate` 同步生成代码到 `dep/protobuf/gen`
- 在 Proto 仓库变更后依次执行：`buf lint` → `buf push`
- Proto 仓库推送成功后，在本仓库执行 `buf generate`（或 `make generate`/`make init`）同步生成代码，并运行 `go test ./...`、`go vet ./...`
