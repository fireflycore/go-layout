---
trigger: manual
---

# API 设计规范（go-layout，Proto/gRPC）v1.0

## [规则 1] Proto First [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 先改 Proto（通常在独立 Proto 仓库），再生成，再实现
- 生成产物以项目约定为准（本项目：`dep/protobuf/gen`）

## [规则 2] 入口统一做参数校验 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service 入口必须 `protovalidate.Validate(req)`
- 不重复实现与 Proto 同义的校验逻辑

## [规则 3] RPC 与消息命名 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service/RPC/Message 用 PascalCase
- RPC 用动词（Create/Get/Update/Delete/List）

## [规则 4] 列表分页字段一致 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Request 必须包含 `uint64 page = 1;` 和 `uint64 page_size = 2;`
- Response 的 `data` 字段必须包含 `int64 total = 1;` 和 `repeated Item list = 2;`

## [规则 5] 向后兼容 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 不删除/复用字段号；弃用用 `deprecated` 并保留字段号

## [规则 6] 错误对外呈现 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Biz/Data 返回 Go error
- Service 默认用响应体 `Code/Message` 表达失败并返回 `nil` error
- 若项目已采用 gRPC status/拦截器统一错误语义，则保持一致，避免同一服务内混用两套错误体系

## [规则 7] 请求/响应 DTO 只以 Proto 为准 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service/Biz/Data 交互的 DTO 以 `dep/protobuf/gen` 生成类型为准
- 禁止复制/手写“同名同义”的请求响应结构体，避免漂移与兼容性问题

## [规则 8] 用户上下文只从 metadata 解析一次 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Service 层从入站 metadata 提取并解析用户上下文（对齐 `micro.ParseUserContextMeta`）
- Biz/Data 不依赖 gRPC metadata 类型，只接收解析后的结构或必要字段

## [规则 9] Service 入口遵循统一处理骨架 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 先初始化响应体为成功（示例：`Code=200/Message=success`）
- 依次执行：`protovalidate.Validate(req)` → 解析 metadata（如需用户上下文）→ 调用 Biz
- 任一步失败：仅修改 `Code/Message` 并返回响应体，错误返回值保持为 `nil`

## [规则 10] 统一响应格式 [ENABLED]
STATUS: ENABLED
PRIORITY: CRITICAL
LANGUAGE: Go
说明：
- 所有 RPC Response 必须包含：
  - `uint32 code = 1;`
  - `string message = 2;`
  - `data` 字段 (field 3)，必须始终存在于定义中
- 如果无实际数据返回：
  - 定义为：
    ```protobuf
    // 空值无用意
    optional string data = 3;
    ```
- 如果有实际数据返回：
  - 定义为 `Type data = 3;` 或 `DataList data = 3;`（无上述注释）

## [规则 11] 注释规范 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- Proto 文件中，注释必须位于字段、Message 或 RPC 方法的**上方**，禁止行尾注释
- AI 生成的代码（包括 Proto 和 Go 实现）必须包含清晰的**中文注释**

## [规则 12] 字段校验规范 [ENABLED]
STATUS: ENABLED
PRIORITY: MEDIUM
LANGUAGE: Go
说明：
- 字段约束不是强制的，仅在用户明确要求或 AI 判断场景必要时添加
- 若需要约束，推荐使用 `buf validate`
- AI 应根据业务场景自行判断是否需要约束（例如：状态枚举值范围、ID 格式等）
- 常见约束参考：
  - 必填：`[(buf.validate.field).required = true]`
  - 字符串：`min_len`, `max_len`, `uuid`, `email`
  - 数字：`gt`, `lt`, `gte`, `lte`
  - 枚举/值：`in`, `const`
- 官方文档：`https://buf.build/bufbuild/protovalidate/docs/main:buf.validate`

## [规则 13] GORM 查询错误处理 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 底层已封装 GORM 并统一上报 SQL 错误到 Logger 服务
- 普通查询（Find, Scan, Count）无需在 Data 层抛出错误，直接返回结果即可
- 只有明确需要校验存在性（如 First/Take 且业务依赖 RecordNotFound）的操作才抛出错误

## [规则 14] 代码风格规范 [ENABLED]
STATUS: ENABLED
PRIORITY: HIGH
LANGUAGE: Go
说明：
- 代码应尽可能简洁、统一、工整、易懂、高效
- 能直接赋值的，避免使用不必要的中间变量进行周转
- 保持函数简短，逻辑清晰
- **代码分块**：变量定义、核心逻辑、返回值之间应有空行分隔，保持清晰的视觉结构（参考 `GetDemoList`）
- **命名见名知意**：函数命名应明确操作对象，例如 `GetDemoCount` 优于 `GetCount`
