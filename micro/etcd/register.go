package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"microservice-go/model/base"
	"microservice-go/store"
	"strconv"
	"strings"
	"time"
)

// Register etcd service register
func (Entrance) Register(service string) {
	key := fmt.Sprintf("/microservice/%s/%s/%s/%d", store.Use.Config.Local.Micro.GatewayName, store.Use.Config.Local.Micro.Namespace, service, store.Use.Micro.Lease)
	val, _ := json.Marshal(base.MicroserviceEntity{
		Name:      service,
		Endpoints: store.Use.Config.Local.Micro.Address.Outside,
	})
	_, err := store.Use.Micro.Cli.Put(store.Use.Micro.Ctx, key, string(val), clientv3.WithLease(store.Use.Micro.Lease))
	if err != nil {
		store.Use.Logger.Func.Error(err.Error())
		return
	}
	store.Use.Logger.Func.Info(fmt.Sprintf("register microservice: %s, %s", key, val))
}

// Deregister etcd service deregister
func (Entrance) Deregister() {
	if _, err := store.Use.Micro.Cli.Revoke(store.Use.Micro.Ctx, store.Use.Micro.Lease); err != nil {
		store.Use.Logger.Func.Error(err.Error())
		return
	}
	store.Use.Logger.Func.Info("revoke service lease success")

	key := fmt.Sprintf("/microservice/%s/%s", store.Use.Config.Local.Micro.GatewayName, store.Use.Config.Local.Micro.Namespace)
	res, rErr := store.Use.Micro.Cli.KV.Get(store.Use.Micro.Ctx, key, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if rErr != nil {
		store.Use.Logger.Func.Error(rErr.Error())
		return
	}
	for _, item := range res.Kvs {
		if strings.Contains(string(item.Key), strconv.FormatInt(int64(store.Use.Micro.Lease), 10)) {
			if _, err := store.Use.Micro.Cli.Delete(store.Use.Micro.Ctx, key); err != nil {
				store.Use.Logger.Func.Error(err.Error())
				continue
			}
		}
	}
	store.Use.Logger.Func.Info("deregister service success")
}

// CreateLease etcd create service instance lease
func (Entrance) CreateLease(cli *clientv3.Client) (lease clientv3.LeaseID) {
	logPrefix := "create lease"
	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	if cli == nil {
		store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, "etcd client not found"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grant, ge := cli.Grant(ctx, 5)
	if ge != nil {
		store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, ge.Error()))
		return
	}

	kac, ke := cli.KeepAlive(store.Use.Micro.Ctx, grant.ID)
	if ke != nil {
		store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, ke.Error()))
		return
	}

	go func() {
		//for v := range kac {
		//	store.Use.Logger.Func.Info(fmt.Sprintf("microservice lease keepalive success, lease %d, ttl %d", v.ID, v.TTL))
		//}
		for range kac {
		}
		store.Use.Logger.Func.Info("microservice stop lease alive success")
	}()
	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	return grant.ID
}
