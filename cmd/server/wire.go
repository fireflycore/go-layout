//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"go-layout/internal/data"
	"go-layout/internal/plugin"
	"go-layout/internal/server"
	"go-layout/internal/service"
)

func wireApp(bc *conf.BootstrapConf) (*App, func(), error) {
	panic(wire.Build(
		plugin.ProviderSet,
		conf.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
