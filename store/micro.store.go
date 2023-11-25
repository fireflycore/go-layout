package store

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type _MicroStoreEntrance struct {
	Cli     *clientv3.Client
	Ctx     context.Context // Ctx microservice context
	Cancel  context.CancelFunc
	Lease   clientv3.LeaseID    // Lease microservice instance lease，instance
	Service map[string][]string // Service discover
}
