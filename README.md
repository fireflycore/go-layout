## Firefly Microservice Framework Golang Version
A minimalist microservices framework based on microcore-go

### config.yaml
- system
  - app_id -------- No special significance, only as a unique identifier.
  - run_port -------- Running port.
- logger
  - addr ------ Remote log storage address.
  - console ------ Local log output.
  - remote: ------ Remote log output.
- micro
  - address: ------ Service Address
  - namespace ------ Service namespace
  - max_retry ------ Maximum number of lease retries
  - ttl ------ Lease heartbeat time

### start-up process
- Deploying etcd.
- Complete config.yaml.
- create task in bootstrap/core.go, read local or remote etcd config file.
- `go mod tidy`
- `go run script/core.go` 
  - Execute the buf cli script where protobuf is stored in the cloud, optional
- `go run main.go`
- 
