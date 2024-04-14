package core

import (
	"encoding/json"
	"fmt"
	"github.com/lhdhtrc/func-go/file"
	"github.com/lhdhtrc/task-go/model"
	taskModel "github.com/lhdhtrc/task-go/model"
)

func ReadRemoteConfig(source []string, config []interface{}) []taskModel.TaskEntity {
	var tasks []taskModel.TaskEntity

	for i, it := range source {
		tasks = append(tasks, createReadRemoteTask(i, it, config[i]))
	}

	return tasks
}

// 创建一个新的函数，接收所需参数，避免闭包捕获问题
func createReadRemoteTask(index int, sourceItem string, configItem interface{}) taskModel.TaskEntity {
	return model.TaskEntity{
		Id: fmt.Sprintf("ReadRemoteConfig_%d", index),
		Handle: func() error {
			bytes, err := file.ReadRemote(sourceItem)
			if err != nil {
				return err
			}
			err = json.Unmarshal(bytes, configItem)
			return err
		},
	}
}
