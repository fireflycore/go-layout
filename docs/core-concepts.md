# 核心概念

1. DO (Domain Object) - 领域对象
    - **定义**：业务领域中的核心实体，承载业务逻辑和业务规则
    - **职责**：封装业务行为和状态，代表业务概念本身
    - **特征**：
        - 包含业务逻辑方法，使用贫血模型，只是定义业务模型
        - 与数据库表结构不一定完全对应
        - 关注业务规则而非数据存储或传输
        - **注意**：在简单业务场景中，DO 可能被省略，直接使用 DTO 或 PO 进行流转
    - **位置**：`internal/biz/model/` (例如 `demo.go`)

2. PO (Persistent Object) - 持久化对象
    - **定义**：与数据库表结构一一对应的对象
    - **职责**：描述数据如何存储在数据库中
    - **特征**：
        - 与数据库表完全映射
        - 通常是贫血的（只有数据字段）
        - 使用 ORM Tag 定义映射关系
    - **位置**：`internal/data/entity/` (例如 `demo.go`)

3. DTO (Data Transfer Object) - 数据传输对象
    - **定义**：用于进程间或网络间数据传输的对象
    - **职责**：定义 API 的请求和响应格式
    - **特征**：
        - 定义 API 契约和版本兼容性
        - 扁平化结构，便于序列化
        - 无业务逻辑
        - 包含数据验证规则注解
        - 通常是 grpc 生成的依赖代码
    - **位置**：
        - `dep/protobuf/`：Proto 生成的代码
        - `dep/dto/`：goverter 生成的转换代码

4. DAO (Data Access Object) - 数据访问对象
    - **定义**：封装对数据源访问的对象
    - **职责**：提供对数据库的增删改查，不暴露数据库内部细节
    - **特征**：
        - 实现 biz 层的 repo 接口
        - 封装所有数据访问细节
        - 处理 PO 对象的 CRUD 操作，或直接返回 DTO 以优化性能
    - **位置**：`internal/data/` 根目录下

## 数据转换

> 数据转换依托于 goverter，通过定义数据转换接口，自动生成转换方法

- **DTO ↔ PO**：在 `internal/biz/convert/` 定义接口
    - 写入场景：Biz 层调用 Convert 接口将 Request DTO 转换为 PO (`entity`)，再传给 Repo。
    - 读取场景：Repo 层可直接返回 DTO (`pb`)，避免不必要的中间转换。

- **Biz 交互原则**
    - Biz 提供给 Service 的通常是 DTO。
    - Biz 入参可以是 DTO 或 Biz Model (DO)。
    - Data 入参通常是 PO，出参可以是 PO 或 DTO。

## 数据验证

> 数据验证依托于 buf-cli 的验证插件, 字段规则定义与 grpc proto 中

1. 在 proto 文件中定义验证规则
2. 使用 `buf generate` 生成包含验证代码的 gRPC 代码
3. 在 Service 层使用 `protovalidate` 进行验证
