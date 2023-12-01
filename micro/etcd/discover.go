package etcd

import (
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"microservice-go/model/base"
	"microservice-go/store"
	"microservice-go/utils/array"
	"reflect"
	"strings"
)

// Watcher etcd service watcher
func (Entrance) Watcher(config *[]base.MicroserviceDiscoverEntity) {
	logPrefix := "[service_endpoint_change] service"
	for _, row := range *config {
		prefix := []string{"/microservice"}
		value := reflect.ValueOf(row)
		for i := 0; i < value.NumField(); i++ {
			v := value.Field(i).String()
			if v != "" {
				prefix = append(prefix, v)
			}
		}

		// todo 第一次先将需要的服务先拉下来
		initService(strings.Join(prefix, "/"))

		wc := store.Use.Micro.Cli.Watch(store.Use.Micro.Ctx, strings.Join(prefix, "/"), clientv3.WithPrefix(), clientv3.WithPrevKV())
		go func() {
			for v := range wc {
				for _, e := range v.Events {
					var (
						bytes []byte
						key   string
						val   base.MicroserviceEntity
					)

					if e.PrevKv != nil {
						key = string(e.PrevKv.Key)
						bytes = e.PrevKv.Value
					} else {
						key = string(e.Kv.Key)
						bytes = e.Kv.Value
					}

					if err := json.Unmarshal(bytes, &val); err != nil {
						store.Use.Logger.Func.Warning(err.Error())
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
						store.Use.Logger.Func.Success(fmt.Sprintf("%s %s put endpoint, key: %s, endpoint: %s", logPrefix, val.Name, key, val.Endpoints))
					// DELETE
					case 1:
						store.Use.Micro.Service[key] = array.Filter(store.Use.Micro.Service[val.Name], func(index int, item string) bool {
							return item != val.Endpoints
						})
						store.Use.Logger.Func.Warning(fmt.Sprintf("%s %s delete endpoint, key: %s, endpoint: %s", logPrefix, val.Name, key, val.Endpoints))
					}
				}
			}
		}()
	}
}

// initService etcd service init
func initService(prefix string) {
	logPrefix := "service discover init service"
	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s", logPrefix, "start ->"))

	res, rErr := store.Use.Micro.Cli.KV.Get(store.Use.Micro.Ctx, prefix, clientv3.WithPrefix())
	if rErr != nil {
		store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, rErr.Error()))
		return
	}

	for _, item := range res.Kvs {
		key := string(item.Key)

		var val base.MicroserviceEntity
		if err := json.Unmarshal(item.Value, &val); err != nil {
			store.Use.Logger.Func.Error(fmt.Sprintf("%s %s", logPrefix, err.Error()))
			return
		}

		st := strings.Split(key, "/")
		st = st[:len(st)-1]
		key = strings.Join(st, "/")

		temp := append(store.Use.Micro.Service[key], val.Endpoints)
		store.Use.Micro.Service[key] = array.Unique[string](temp, func(index int, item string) string {
			return item
		})
	}

	store.Use.Logger.Func.Info(fmt.Sprintf("%s %s", logPrefix, "success ->"))
}
