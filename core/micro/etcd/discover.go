package etcd

import (
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"microservice-go/model/base"
	"microservice-go/store"
	"microservice-go/utils"
	"microservice-go/utils/array"
	"reflect"
	"strings"
)

// Discover etcd service discover
func (Entrance) Discover(config []*base.MicroserviceDiscoverEntity) {
	for _, row := range config {
		var prefix []string
		value := reflect.ValueOf(row)
		for i := 0; i < value.NumField(); i++ {
			v := value.Field(i).String()
			if v != "" {
				prefix = append(prefix, v)
			}
		}

		wc := store.Use.Micro.Cli.Watch(store.Use.Micro.Ctx, strings.Join(prefix, "/"), clientv3.WithPrefix(), clientv3.WithKeysOnly())
		go func() {
			for v := range wc {
				for _, e := range v.Events {
					key := string(e.Kv.Key)
					var val base.MicroserviceEntity
					if err := json.Unmarshal(e.Kv.Value, &val); err != nil {
						utils.Use.Logger.Warning(err.Error())
						continue
					}

					st := strings.Split(key, "/")
					st = st[:len(st)-1]
					key = strings.Join(st, "/")

					switch e.Type {
					// PUT，新增或替换
					case 0:
						temp := append(store.Use.Micro.Service[key], val.Endpoints)
						store.Use.Micro.Service[key] = array.Unique[string](temp, func(index int, item string) string {
							return item
						})
						fmt.Printf("[service_endpoint_change] service %s put endpoint, key: %s, endpoint: %s\n", val.Name, key, val.Endpoints)
					// DELETE
					case 1:
						store.Use.Micro.Service[key] = array.Filter(store.Use.Micro.Service[val.Name], func(index int, item string) bool {
							return item != val.Endpoints
						})
						fmt.Printf("[service_endpoint_change] service %s delete endpoint, key: %s, endpoint: %s\n", val.Name, key, val.Endpoints)
					}
				}
			}
		}()
	}
}

func (Entrance) GetService(name string) []*base.MicroserviceEntity {
	key := fmt.Sprintf("/microservice/%s/%s/%s", store.Use.Config.Local.Micro.GatewayName, store.Use.Config.Local.Micro.Namespace, name)
	kvs, ke := store.Use.Micro.Cli.KV.Get(store.Use.Micro.Ctx, key, clientv3.WithPrefix())
	if ke != nil {
		utils.Use.Logger.Error(ke.Error())
		return nil
	}
	items := make([]*base.MicroserviceEntity, 0, len(kvs.Kvs))
	for _, kv := range kvs.Kvs {
		var val base.MicroserviceEntity
		if err := json.Unmarshal(kv.Value, &val); err != nil {
			utils.Use.Logger.Error(err.Error())
			return nil
		}
		if val.Name != name {
			continue
		}
		items = append(items, &val)
	}
	return items
}
