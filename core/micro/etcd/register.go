package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"microservice-go/model/base"
	"microservice-go/store"
	"microservice-go/utils"
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
		utils.Use.Logger.Error(err.Error())
		return
	}
	utils.Use.Logger.Info(fmt.Sprintf("register microservice: %s, %s", key, val))
}

// CreateLease etcd create service instance lease
func (Entrance) CreateLease(cli *clientv3.Client) (lease clientv3.LeaseID) {
	logPrefix := "create lease"
	utils.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	if cli == nil {
		utils.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, "etcd client not found"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grant, ge := cli.Grant(ctx, 5)
	if ge != nil {
		utils.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, ge.Error()))
		return
	}

	kac, ke := cli.KeepAlive(store.Use.Micro.Ctx, grant.ID)
	if ke != nil {
		utils.Use.Logger.Error(fmt.Sprintf("%s %s", logPrefix, ke.Error()))
		return
	}

	go func() {
		for v := range kac {
			utils.Use.Logger.Info(fmt.Sprintf("microservice lease keepalive success, lease %d, ttl %d", v.ID, v.TTL))
		}
		utils.Use.Logger.Info("microservice stop lease alive success")
	}()
	utils.Use.Logger.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))

	return grant.ID
}
