# 升级需求：make run 按目录启动以包含 wire_gen.go

## 问题现象
- 执行 `make run` 报错：`cmd/server/main.go:15:14: undefined: wireApp`
- GoLand 使用运行配置（目录：`$PROJECT_DIR$/cmd/server`）可以正常启动。

## 原因
- `wireApp` 定义在 `cmd/server/wire_gen.go`（`!wireinject` build tag 生效时参与编译）。
- `go run ./cmd/server/main.go` 只编译单个文件，不会把同目录下的 `wire_gen.go` 一起编译进来，因此 `wireApp` 未定义。
- GoLand 的“目录运行”会以包为单位编译 `cmd/server` 目录下所有文件，所以可正常找到 `wireApp`。

## 目标
- `make run` 的启动方式与 GoLand 目录运行保持一致，确保 `wire_gen.go` 被编译。

## 改动
- `makefile`：
  - `go run ./cmd/server/main.go` 改为 `go run ./cmd/server`

## 升级适用范围
- 所有使用 Wire 生成注入代码、且 `make run` 使用 `go run` 启动的服务。

## 回归验证
- 至少确保 `go test ./cmd/server` 编译通过。

