# Lhdht Microservice Golang Version

#### Description
lhdht microservice golang version

#### Function module
- store
    - An in-program global repository that stores global state
- config
    - Local configuration is loaded and used `viper`
    - Remote configuration loading（todo）
- db
    - mongo
        - mode：Supports single-node and cluster configuration
        - auth：None/Account password /TLS
- logger
    - Local adoption `zap`（todo）
    - Remote store to `mongodb`（todo）
    - Extracted as microservice module（To be detached）
- micro
    - register
    - discovery
    - heartbeat
    - load balancing（todo）
    - current limiting fus（todo）

#### Start-up process
- 1.complete `/config.yaml`;