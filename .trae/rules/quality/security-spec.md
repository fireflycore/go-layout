---
trigger: manual
---

# 安全规范（Go 微服务）v1.0
# ============================================
# 面向 go-layout 的安全标准与要求
# 通过将 [ENABLED] 更改为 [DISABLED] 来启用/禁用规则
#
# 最后更新：2025-12-31
# ============================================

## [规则 1] 输入验证与清理 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
说明：
- 对所有外部输入（RPC 请求、配置、环境变量、第三方返回）进行验证
- 以白名单为主：类型、长度、范围、枚举值
- 参数校验优先复用 Proto 规则（`protovalidate`），避免双份逻辑

## [规则 2] 认证与授权 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
说明：
- 所有授权判断必须在服务端完成
- 用户上下文从入站 metadata 提取并校验（与项目 `micro.ParseUserContextMeta` 用法一致）
- 按最小权限原则设计接口

## [规则 3] 敏感数据保护 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
说明：
- 不在日志、错误消息、响应体中暴露密钥、令牌、密码、个人敏感信息
- 传输层使用 TLS（生产环境必须）
- 密码哈希使用成熟算法（如 bcrypt），不要自研

## [规则 4] 数据库/存储注入防护 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 所有查询必须参数化，避免拼接 SQL
- 使用 ORM/Query Builder 时也避免拼接未转义的用户输入

## [规则 5] 依赖安全管理 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 最小化新增依赖，优先复用 go.mod 现有依赖
- 更新依赖后执行 `go mod tidy` 并确保可构建
- 对安全敏感依赖升级保持谨慎：优先安全补丁与关键修复

## [规则 6] 日志安全 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
说明：
- 记录错误时包含上下文（请求标识、业务关键字段），但避免泄露敏感信息
- 统一通过注入的 Logger 打印日志，避免随意 `fmt.Println` 输出敏感内容

