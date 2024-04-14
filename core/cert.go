package core

import (
	"fmt"
	"github.com/lhdhtrc/func-go/file"
	taskModel "github.com/lhdhtrc/task-go/model"
	"path/filepath"
	"reflect"
)

func GetRemoteCert(dir string, config interface{}) []taskModel.TaskEntity {
	var tasks []taskModel.TaskEntity
	dirPath := filepath.Join("dep", "cert", dir)

	// 遍历config的字段
	valueOfConfig := reflect.ValueOf(config).Elem()
	typeOfConfig := valueOfConfig.Type()
	for i := 0; i < valueOfConfig.NumField(); i++ {
		fieldValue := valueOfConfig.Field(i)
		fieldType := typeOfConfig.Field(i)
		if fieldValue.IsValid() && !fieldValue.IsZero() && fieldType.Type.Kind() == reflect.String {
			remote := fieldValue.String()

			// 分割路径，得到文件名部分
			f := filepath.Base(remote)
			local := filepath.Join(dirPath, f)

			// 创建一个新闭包，并立即执行，以避免闭包捕获问题
			task := createDownloadTask(dir, remote, local)
			tasks = append(tasks, task)
		}
	}

	return tasks
}

// 创建一个新的函数，接收所需参数，避免闭包捕获问题
func createDownloadTask(dir, remote, local string) taskModel.TaskEntity {
	return taskModel.TaskEntity{
		Id: fmt.Sprintf("DownloadRemoteCert_%s_%s", dir, filepath.Base(remote)),
		Handle: func() error {
			read, err := file.ReadRemote(remote)
			if err != nil {
				return err
			}

			err = file.WriteLocal(local, read)
			return err
		},
	}
}
