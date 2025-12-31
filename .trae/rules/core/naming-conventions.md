---
trigger: manual
---

# 命名约定规范 v1.0
# ============================================
# 面向 go-layout（Go 微服务模板）的命名标准
# 通过将 [ENABLED] 更改为 [DISABLED] 来启用/禁用约定
#
# 最后更新：2025-12-31
# ============================================

## [约定 1] 变量命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 本地变量与非导出字段使用 lowerCamelCase（例如 `repo`, `demoRepo`）
- 布尔变量使用问题式命名（如 `isReady`, `hasMore`），并与项目现有风格保持一致
- 常用约定：`ctx`（context）、`req`/`request`、`res`/`result`、`err`
- 集合/切片使用复数（如 `items`, `ids`），单个对象用单数
- 对历史风格差异（例如 `UserId` vs `UserID`）以项目现状为准，避免混用

## [约定 2] 函数/方法命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 导出函数/方法使用 PascalCase（如 `CreateDemo`, `NewDemoService`）
- 非导出函数/方法使用 lowerCamelCase
- 构造函数以 `New` 开头，返回具体类型或接口（与现有代码保持一致）
- 错误变量统一使用 `err`，并在同一作用域内避免复用不同语义

## [约定 3] 类型/接口命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 结构体/接口/类型别名使用 PascalCase
- 接口不使用 `I` 前缀；单方法接口可使用 `-er` 形式（如 `Reader`），但以项目现有风格为准
- 分层后缀保持一致：Biz 层 UseCase（如 `DemoUseCase`）、Data 层 Repo（如 `DemoRepo`）、Service 层 Service（如 `DemoService`）

## [约定 4] 常量与枚举命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- Go 常量按导出性选择 PascalCase/lowerCamelCase
- 使用 `const (...)` 分组相关常量；需要枚举时使用 `iota` 并保持语义清晰
- 环境变量名使用 UPPER_SNAKE_CASE（见 [约定 9]）

## [约定 5] 文件与包命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 包名：全小写、短、无下划线（如 `biz`, `data`, `service`）
- 文件名：全小写；多个单词使用下划线分隔（如 `rs_logger.go`, `wire_gen.go`）
- 与目录职责一致：Biz/Data/Service/Server 的文件放在对应目录

## [约定 6] 接收者命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 方法接收者使用 1–3 个字符的缩写，并在同一类型上保持一致
- 建议与现有代码对齐：UseCase 用 `uc`，Service 用 `srv`，Repo 用 `uc` 或 `repo`

## [约定 7] 数据库实体命名（GORM）[DISABLED]
STATUS: DISABLED
LANGUAGE: Go
说明：
- Go 结构体字段使用 PascalCase
- 数据库表名/列名按项目约定实现（如通过 `Table()` 返回表名）
- JSON tag 与接口返回保持一致（本项目示例使用 `json:"user_id"` 这类 snake_case）

## [约定 8] gRPC/Proto 命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- Service/RPC/Message 使用 PascalCase（如 `DemoService`, `CreateDemoRequest`）
- Service 层方法名与 RPC 保持一致（例如 `CreateDemo`）
- 参数校验使用 `protovalidate.Validate(req)`，避免漂移

## [约定 9] 环境变量命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: All
说明：
- 使用 UPPER_SNAKE_CASE
- 使用应用/服务名称前缀以提高清晰度
- 使用公共前缀分组相关变量

## [约定 10] 错误命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 可复用的哨兵错误使用 `var ErrXxx = errors.New("...")`
- 错误包装使用 `%w`，分支判断使用 `errors.Is/As`
- 对外返回时避免暴露敏感信息（与安全规范一致）

## [约定 11] 测试命名 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- 测试文件使用 `_test.go`
- 测试函数使用 `TestXxx`
- 子测试使用 `t.Run("case-name", ...)`，名称描述行为

## [约定 12] 生成代码命名与维护 [ENABLED]
STATUS: ENABLED
LANGUAGE: Go
说明：
- `wire_gen.go`、`*_gen.go`、`dep/` 下生成内容不手动修改
- 需要更新生成产物时，使用既有生成链路（`buf generate`、`goverter`、`wire`）

# ============================================
# 摘要 - 启用的约定
# ============================================

✅ [约定 1] 变量命名
✅ [约定 2] 函数/方法命名
✅ [约定 3] 类型/接口命名
✅ [约定 4] 常量与枚举命名
✅ [约定 5] 文件与包命名
✅ [约定 6] 接收者命名
✅ [约定 8] gRPC/Proto 命名
✅ [约定 9] 环境变量命名
✅ [约定 10] 错误命名
✅ [约定 11] 测试命名
✅ [约定 12] 生成代码命名与维护

