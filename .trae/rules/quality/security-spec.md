---
trigger: manual
---

# 安全规范（go-layout）v1.2

## [规则 1] 输入与权限只在服务端处理 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 外部输入统一在 Service 入口校验（优先复用 Proto 校验）
- 权限判断在服务端完成；用户上下文从 metadata 提取并校验
- 存量兼容：不为“对齐风格”改动现有鉴权链路与字段语义

## [规则 2] 不泄露敏感信息 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: All
说明：
- 日志/错误/响应体禁止出现密钥、令牌、密码、个人敏感信息
- 仓库内配置可包含示例/开发值；严禁提交真实生产密钥，生产通过私有配置或远程配置注入

## [规则 3] 防注入与日志卫生 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- DB 查询参数化，禁止拼接用户输入
- 统一使用注入的 Logger，避免 `fmt.Println/println`
