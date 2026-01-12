# 升级需求：注入 zap.Logger 并替换 fmt 输出

## 背景
- 统一使用注入的 Logger，避免 `fmt.Println/println`（见安全规范“防注入与日志卫生”）。
- 为后续对每个服务逐个升级提供统一的改造清单与改动点参考。

## 目标
- 将 `*zap.Logger` 通过依赖注入贯穿需要记录日志的位置（启动、注册、访问日志 Console 输出等）。
- 替换项目中的 `fmt.Print/Println/Printf` 这类“日志输出”用法为 `zap.Logger` 结构化日志。
- `fmt.Sprintf` 仅用于必要的字符串拼接；涉及日志输出时优先改为结构化字段（`zap.String/zap.Int/...`）。

## 特别注意：AccessLogger Console 输出避免双写
- `AccessLogger` 的远程写入已经通过 `loggerRepo.CreateAccessLog(...)` 进入 Logger 服务的 AccessLog 管道。
- 如果在 `AccessLogger` 的 Console 分支里再用注入的 `*zap.Logger` 输出，会走 `ServerLogger` 管道，导致同一条访问日志被写两遍（AccessLog + ServerLog），造成重复与浪费。
- 因此：当 `bootstrapConf.Logger.Console == true` 时，AccessLogger 的 Console 输出应直接写到标准输出，不走 zap、不走 ServerLogger。
  - 推荐：`os.Stdout.WriteString(msg)`

## 本次模板改动点（可作为所有服务升级的参考）
- 访问日志 Console 输出从 `fmt.Print(msg)` 改为使用注入的 `*zap.Logger` 输出：
  - `internal/dep/logger.go`：`NewAccessLogger(bootstrapConf, loggerRepo, zapLogger)`。
- 服务注册失败从 `fmt.Println(errs)` 改为 `logger.Error(...)`：
  - `internal/server/register.go`：`NewRegisterCenterRepo(register, logger)`。
- 注册重试钩子从 `fmt.Println(...)` 改为 `logger.Info(...)`：
  - `cmd/server/app.go`：`NewApp(..., logger)` 内的 `WithRetryBefore/WithRetryAfter`。
- 启动日志去除 `fmt.Sprintf`，改为结构化日志字段：
  - `cmd/server/main.go`：`logger.Info("...", zap.String(...))`。
- 网络地址拼接去除 `fmt.Sprintf`，改为 `net.JoinHostPort`：
  - `internal/server/server.go`、`internal/conf/bootstrap.go`。
- logger repo 内部不再直接 `fmt.Println(err)`：
  - `internal/data/rs_logger.go`：反序列化失败直接返回（避免标准输出污染）。

## 服务升级检查清单（每个服务都按此执行）
- 确保所有“日志输出”不使用 `fmt.Print/Println/Printf`，统一用注入的 `*zap.Logger`。
- 启动/注册/拦截器/中间件等基础设施层，如需输出信息，必须走 Logger。
- Wire/DI 更新：新增构造函数参数后，重新生成注入代码（例如 `wire ./cmd/server`）。
- 回归命令（至少）：`go test ./...`、`go vet ./...`。
