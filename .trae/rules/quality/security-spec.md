---
trigger: manual
---

# 安全规范（go-layout）v1.0

## [规则 1] 输入必须验证 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 外部输入（RPC/配置/环境变量/第三方返回）均要验证
- 校验优先复用 Proto 规则：`protovalidate`

## [规则 2] 认证授权只在服务端 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 权限判断必须在服务端完成
- 用户上下文从入站 metadata 提取并校验（对齐 `micro.ParseUserContextMeta`）

## [规则 3] 不泄露敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 日志/错误/响应体不得包含密钥、令牌、密码、个人敏感信息
- 当前模板远程 gRPC 客户端示例使用 `insecure`；生产环境启用 TLS 时需同时对齐 Client/Server 与配置加载方式，避免“只改一侧”导致误连或降级

## [规则 4] 防注入 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- DB 查询参数化，避免拼接用户输入

## [规则 5] 依赖安全 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 最小化新增依赖；变更后执行 `go mod tidy` 并验证可构建

## [规则 6] 日志卫生 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 用注入的 Logger 输出；必要字段可记录，但需脱敏/裁剪

## [规则 7] 仓库内配置只允许示例值 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- `conf/bootstrap.json` 与文档中出现的密钥/令牌/地址必须是示例值
- 真实密钥只能通过本地私有配置或远程配置服务注入

## [规则 8] 远程调用与超时 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 业务链路的远程调用优先沿用入参 `ctx`，必要时基于它派生超时
- 非业务链路的后台任务允许使用 `context.Background()`，但必须设置合理超时与取消
