# Lhdht Microservice Golang Version

#### 介绍
梨花炖海棠微服务，golang版本

#### 功能模块
- store
  - 程序内全局仓库，存放全局状态
- config
  - 本地配置加载使用`viper`
  - 远程配置加载（待开发）
- db 
  - mongo
    - 模式：支持单节点，集群配置
    - 认证：无/账号密码/TLS
- logger 
  - 本地采用`zap`日志（待开发）
  - 远程存储到mongodb中（待开发）
  - 抽取为微服务模块（待抽离）
- micro
  - register
  - discovery
  - heartbeat detection
  - load balancing（负载均衡-待开发）
  - current limiting fus（限流熔断-待开发）

#### 启动流程
- 1.补全`/config.yaml`下的所有内容，`/bootstrap/file/config.yaml`中的配置可以直接打包到应用中，如果部署完想要修改配置，则在程序同级目录创建`/config.yaml`;