# 升级需求：同步更新 fireflycore 依赖包与 import 路径

## 背景
- `go-layout` 仍在使用 `github.com/lhdhtrc/*` 系列基础包。
- `update.pkg.md` 给出了迁移到 `github.com/fireflycore/*` 的包名与版本映射。

## 目标
- 将模板依赖从 `github.com/lhdhtrc/*` 同步升级到 `github.com/fireflycore/*` 对应版本。
- 同步更新相关 import 与少量 API 调整，确保 `make init`、`go test ./...` 可通过。

## 变更摘要

### 依赖升级（go.mod）
- `github.com/lhdhtrc/gorm` → `github.com/fireflycore/gormx@v0.8.1`
- `github.com/lhdhtrc/redis-go` → `github.com/fireflycore/go-redis@v0.1.1`
- `github.com/lhdhtrc/etcd-go` → `github.com/fireflycore/go-etcd@v0.2.6`
- `github.com/lhdhtrc/func-go` → `github.com/fireflycore/go-utils@v0.3.5`
- `github.com/lhdhtrc/logger-go` → `github.com/fireflycore/go-logger@v0.2.1`
- `github.com/lhdhtrc/micro-go` → `github.com/fireflycore/go-micro@v0.6.9`

补充：
- `go` 版本从 `1.25.1` 更新为 `1.25.4`（用于满足 `go-micro@v0.6.9` 的最低版本要求）。

### import 路径调整（代码）
由于新包目录结构与旧包不同，本次迁移不是简单的前缀替换：
- go-micro：
  - 原 `.../micro-go/pkg/core` → 新 `github.com/fireflycore/go-micro/rpc`（UserContextMeta/metadata/WithRemoteInvoke）
  - 原 `.../micro-go/pkg/middleware` → 新 `github.com/fireflycore/go-micro/middleware`
  - 原 `.../micro-go/pkg/etcd` → 新 `github.com/fireflycore/go-micro/registry/etcd`
  - 注册/配置相关类型迁移到 `github.com/fireflycore/go-micro/registry`
- go-utils：
  - `process.Watcher` 来自 `github.com/fireflycore/go-utils/process`
  - 写文件接口从 `file.WriteLocal` 调整为 `file.WriteLocalFile`
- gormx：
  - `gormx` 根包为 `github.com/fireflycore/gormx`
  - scope 为 `github.com/fireflycore/gormx/scope`

### 业务逻辑微调（Demo 更新）
`DemoUseCase.UpdateDemo` 原先依赖的“差异字段过滤”函数在新 `go-utils` 中不存在，改为显式构建更新字段 map：
- 仅当请求字段非零值且与原值不同，才写入 `updates`。

## 回归验证
- `go mod tidy`
- `make init`
- `go test ./...`
- `go vet ./...`

## 适用范围
- 需要从 `github.com/lhdhtrc/*` 系列包迁移到 `github.com/fireflycore/*` 的所有服务仓库。
- 本次模板未使用的映射项（如 `go-mongo`、`go-proxy`）不影响迁移，可按需在具体服务中启用。

