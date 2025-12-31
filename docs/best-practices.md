# Firefly Go 最佳实践

1. **领域对象设计**
   - 灵活使用充血模型或贫血模型。
   - 简单 CRUD 场景可直接使用 DTO/PO，减少样板代码。
   - 复杂业务逻辑应封装在 UseCase 或 DO 中。

2. **数据转换管理**
   - 所有转换通过明确的转换器 (`internal/biz/convert`) 进行。
   - 转换逻辑集中管理，便于测试维护。
   - 使用 `goverter` 自动生成转换代码，减少手动错误。

3. **验证策略**
   - DTO 字段验证使用 `protovalidate` (在 Service 层入口处)。
   - 业务规则验证在 `UseCase` 中进行。
   - 数据库约束在 `PO` (GORM tags) 中定义。

4. **错误处理**
   - 验证错误：返回 `codes.InvalidArgument`。
   - 业务错误：返回对应 gRPC 状态码。
   - 系统错误：记录日志，返回 `codes.Internal`。

5. **配置管理**
   - 关键配置在 `bootstrap.json` 定义。
   - 支持多配置源（本地文件、etcd 等）。
   - 配置变更支持热更新 (通过 Etcd Watch)。

6. **依赖管理**
   - 使用 `Wire` 进行依赖注入。
   - 明确各层依赖关系：Service -> Biz -> Data。
   - `internal/dep` 封装外部基础设施依赖，解耦业务逻辑。
