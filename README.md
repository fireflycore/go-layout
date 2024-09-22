## Firefly Microservice Framework Golang Version
A minimalist microservices framework based on microcore-go

## Documentation
[文档](https://firefly.lhdht.cn/guide/)

### start-up process
- Deploying etcd.
- Complete config.yaml.
- create task in bootstrap/core.go, read local or remote etcd config file.
- `go mod tidy`
- `go run script/core.go` 
  - Execute the buf cli script where protobuf is stored in the cloud, optional
- `go run main.go`