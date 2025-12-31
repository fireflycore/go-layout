---
trigger: manual
---

# 安全规范（go-layout）v1.0

## [规则 1] 输入必须验证 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
- 外部输入（RPC/配置/环境变量/第三方返回）均要验证
- 校验优先复用 Proto 规则：`protovalidate`

## [规则 2] 认证授权只在服务端 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
- 权限判断必须在服务端完成
- 用户上下文从入站 metadata 提取并校验（对齐 `micro.ParseUserContextMeta`）

## [规则 3] 不泄露敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
- 日志/错误/响应体不得包含密钥、令牌、密码、个人敏感信息
- 生产环境使用 TLS

## [规则 4] 防注入 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- DB 查询参数化，避免拼接用户输入

## [规则 5] 依赖安全 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 最小化新增依赖；变更后执行 `go mod tidy` 并验证可构建

## [规则 6] 日志卫生 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
- 用注入的 Logger 输出；必要字段可记录，但需脱敏/裁剪
