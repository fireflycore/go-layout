---
trigger: manual
---

# 安全规范（go-layout）v1.4

## [规则 1] 输入与权限只在服务端处理 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 外部输入统一在 Service 入口校验（优先复用 Proto 校验）
- 权限判断在服务端完成；用户上下文从 metadata 提取并校验

## [规则 2] 不泄露敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 真实密钥/令牌/密码/个人敏感信息不得进入代码与仓库（含配置文件、示例与日志文件）
- 对外错误信息允许直接透传 `err.Error()`；不做敏感信息判定与兜底改写
- 真实错误允许直接写入日志；不做敏感信息判定与兜底改写
敏感信息最小判定（用于检查是否误入库，示例）：
- 认证凭据：Authorization、Bearer token、Cookie/Set-Cookie、session
- 密钥材料：app_secret、api_key、access_key、secret_key、private_key、证书明文
- 个人敏感信息：身份证号、银行卡号、手机号（按项目定义）

## [规则 3] 防注入与日志卫生 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- DB 查询参数化，禁止拼接用户输入
- 服务端日志统一使用 ServerLogger（封装 zap），通过 go-logger 注入并上报日志服务
- 避免 `fmt.Println/println`
