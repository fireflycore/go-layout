package data

import (
	"github.com/google/wire"
	"go-layout/dep/goverter/data"
)

var ProviderSet = wire.NewSet(NewMysql, NewData, NewDemoRepo)

var (
	demoDTO = new(data.DemoConverterImpl)
)
