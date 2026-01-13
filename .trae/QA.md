# QA 参考（给 AI 的补充口径）

## 方法替换与兼容
Q：一个方法变成另一个方法，旧方法要保留兼容吗？
A：一般不需要。方法替换/重命名时直接删除旧方法，并同步更新所有调用方，不保留兼容层。

## 对外错误透传
Q：对外错误是否允许透传 `err.Error()`？
A：允许。go-layout 是模板仓库，当前不强制做对外文案收敛或敏感信息兜底改写。

## gormx 日志链路
Q：gorm SQL 执行日志/异常怎么上报？业务层要处理吗？
A：不需要。gormx 是独立封装库，SQL 执行日志可通过 OperationLogger 上报到日志服务，业务层无需人工干预。

## 业务 ctx 贯通
Q：Service/Biz/Data 的 ctx 要怎么写才算贯通？
A：Service 收到的入参 ctx 原样透传给 Biz/Repo；Data 层所有 GORM/DB 调用必须 `db.WithContext(ctx)`；业务方法内禁止用 `context.Background/TODO` 替代入参 ctx。

## 服务端日志
Q：服务端应该怎么打日志？
A：使用 ServerLogger（封装后的 zap）。通过 go-logger 包以注入方式接入 ServerLogger 并上报日志服务；避免 `fmt.Println/println`。

## 注释策略
Q：注释策略是什么？
A：未改代码不改现有注释（除非注释错误/误导/与实现不一致）；改代码时对改动涉及代码逐行补充说明。

## ID/IP 命名
Q：ID/IP 命名要不要遵循团队风格？
A：要，统一用 `Id/Ip`（如 `TraceId`、`ClientIp`），不使用 `ID/IP`。
