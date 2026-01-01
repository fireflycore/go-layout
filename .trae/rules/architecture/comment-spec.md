---
trigger: manual
---

# 注释规范（Go/Proto）v1.0

## [规则 1] 注释必须加且按场景写 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: All
说明：
- 注释要解释“为什么/约束/边界/副作用”，避免逐行翻译代码
- 变更涉及协议、权限、金额、状态机、幂等等关键语义时必须补充注释
- 存量兼容：仅在改动到相关模块时补齐注释，不做顺手批量补注释

## [规则 2] Go Doc 注释要求 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 导出类型/函数/方法必须有 Doc 注释，放在声明上方
- 非导出标识符：只在非直观逻辑、易踩坑、隐含约束时加注释

## [规则 3] Service/Biz/Data 注释要点 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service：标注入口步骤（校验、上下文解析、调用 UseCase、错误映射策略）
- Biz：标注业务规则来源、状态/权限校验、关键分支原因
- Data：标注查询条件含义、索引假设、事务/一致性边界与 NotFound 语义

## [规则 4] 最小注释集（建议） [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 每个 Service RPC：至少 3 条步骤注释（校验/上下文/调用/错误映射）
- 每个 UseCase：关键分支至少 1 条“原因”注释
- 每个 Data 查询：where 条件含义 + NotFound 语义至少 1 条注释

## [规则 5] Proto 注释要求 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Proto
说明：
- 注释放在字段/Message/RPC 上方，禁止行尾注释，使用中文
- 字段需要说明：含义、单位、枚举取值范围、是否可选与缺省语义
