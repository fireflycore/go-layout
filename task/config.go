package task

import (
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/func-go/file"
	taskModel "github.com/lhdhtrc/task-go/model"
	"microservice-go/store"
)

func ReadRemoteConfig(source []string, config []interface{}) {
	for i, it := range source {
		store.Use.Task.Add(taskModel.TaskEntity{
			Id: fmt.Sprintf("ReadRemoteConfig_%d", i),
			Handle: func() error {
				bytes, err := file.ReadRemote(it)
				if err != nil {
					return err
				}
				err = json.Unmarshal(bytes, config[i])
				return err
			},
		})
	}
}
