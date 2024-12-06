package plugin

import (
	"github.com/google/wire"
	task "github.com/lhdhtrc/task-go/pkg"
	"go-layout/internal/conf"
)

var ProviderSet = wire.NewSet(NewTask)

func NewTask(bc *conf.BootstrapConf) *task.Instance {
	return task.New(bc.Task)
}
