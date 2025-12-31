# 数据流转

```markdown
[ 外部请求 (gRPC/HTTP) ]
     ↓
[ DTO ] (dep/protobuf/ 生成)
     ↓ (service/convert/ 转换 + protovalidate 验证)
[ Service 层 ] (参数验证、权限检查、事务控制)
     ↓
[ UseCase 层 (Biz) ] (业务逻辑)
     ↓
     ├─ (读取) → [ Repo 接口 ] → [ DAO ] → [ GORM ] → [ DTO (pb.Demo) ]
     ↓
     └─ (写入) → [ Biz/Convert ] (DTO → PO) → [ PO ] → [ Repo 接口 ] → [ DAO ] → [ GORM ] → [ 数据库 ]

响应流程：
[ 数据库 ] → [ PO ] → [ GORM Scan ] → [ DTO ] → [ 客户端响应 ]
```

**说明：**
1. **读取操作**：为提高性能，Data 层 (DAO) 可直接将数据库查询结果映射为 DTO (`pb.Demo`) 返回，Biz 层透传给 Service 层。
2. **写入操作**：Service 层传入 DTO，Biz 层通过 Convert 接口将 DTO 转换为 PO (`entity.Demo`)，然后传递给 Data 层进行持久化。
3. **DO (Domain Object)**：在简单 CRUD 场景中，Biz 层可能不需要显式的 DO 模型，而是直接操作 DTO 和 PO。复杂业务逻辑中可引入 DO。
