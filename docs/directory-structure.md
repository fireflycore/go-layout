# 项目目录

```markdown
├── .github/
│   └── dependabot.yml          # GitHub 依赖管理机器人配置
├── .run/                       # Goland 运行配置文件，用于启动项目和调试
├── .trae/                      # Trae IDE 配置文件，包含代码规范和架构规则
├── cmd/
│   ├── server/
│   │   ├── app.go              # 服务实例
│   │   ├── main.go             # 服务入口
│   │   ├── wire.go             # 依赖注入入口
│   │   └── wire_gen.go         # Wire 自动生成的依赖注入代码
├── conf/
│   └── bootstrap.json          # 服务引导配置
├── dep/                        # (Generated, Ignored) 生成的依赖代码目录
│   ├── dto/                    # 通过 goverter 生成的转换方法实现
│   └── protobuf/               # 通过 buf-cli 生成的 Proto 代码
├── docs/                       # 项目文档
│   ├── project-guide.md        # 项目指引
│   ├── directory-structure.md  # 项目目录结构
│   ├── data-flow.md            # 数据流转说明
│   ├── core-concepts.md        # 核心概念说明
│   ├── architecture.md         # 架构分层说明
│   └── best-practices.md       # 最佳实践
├── internal/
│   ├── biz/                    # 业务逻辑层
│   │   ├── convert/            # 定义 DTO ↔ PO 转换接口
│   │   │   ├── demo.go         # DemoConvert interface
│   │   │   └── utils.go        # convert 通用方法定义
│   │   ├── model/              # 定义业务领域模型 (DO)，非必要无需定义
│   │   │   └── demo.go         # Demo 领域对象
│   │   ├── repo/               # 定义业务领域依赖的数据接口
│   │   │   ├── demo.go         # DemoRepo interface
│   │   │   ├── rs_config.go    # ConfigRepo interface
│   │   │   ├── rs_logger.go    # LoggerRepo interface
│   │   │   └── transaction.go  # TransactionRepo interface
│   │   ├── demo.go             # DemoUseCase 实现
│   │   └── core.go             # 定义biz层入口和依赖注入（wire）
│   ├── conf/                   # 配置管理
│   │   ├── bootstrap.go        # 定义引导配置
│   │   ├── etcd.go             # 定义获取 etcd 配置的方法
│   │   ├── mysql.go            # 定义获取 mysql 配置的方法
│   │   ├── redis.go            # 定义获取 redis 配置的方法
│   │   ├── utils.go            # 配置辅助方法
│   │   └── core.go             # 定义conf层入口和依赖注入（wire）
│   ├── data/                   # 数据访问层
│   │   ├── entity/             # 定义数据模型 (PO)
│   │   │   └── demo.go         # DemoPO 持久化对象
│   │   ├── demo.go             # DemoRepo 实现 (DAO)
│   │   ├── rs_config.go        # 配置服务相关方法实现
│   │   ├── rs_logger.go        # 日志服务相关方法实现
│   │   ├── transaction.go      # 事务实现
│   │   ├── data.go             # 数据层连接
│   │   └── core.go             # 定义data层入口和依赖注入（wire）
│   ├── dep/                    # 基础设施依赖接口及实现
│   │   ├── logger.go           # 定义日志记录器
│   │   ├── remote.go           # 定义远程服务入口
│   │   └── core.go             # 定义dep层入口和依赖注入（wire）
│   ├── dto/                    # 数据转换层入口
│   │   ├── core.go             # 定义dto层入口和依赖注入（wire）
│   │   └── demo.go             # DemoDTO 构造函数
│   ├── server/                 # 服务层
│   │   ├── core.go             # 服务实例入口
│   │   ├── grpc.go             # 注册 GRPC 服务
│   │   ├── register.go         # 初始化注册服务
│   │   └── server.go           # 初始化 TCP 服务器
│   ├── service/                # 应用服务层
│   │   ├── demo.go             # DemoService 实现
│   │   ├── core.go             # 服务层入口
│   │   └── remote.go           # 定义远程服务入口
├── .gitignore
├── buf.gen.yaml                # buf cli 配置文件
├── go.mod
├── go.sum
├── LICENSE
├── makefile                    # 项目构建和任务管理
├── README.md                   # 服务使用文档
├── update.sh                   # 服务更新脚本（可选）
└── run.sh                      # 服务运行脚本
```
