//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"go-layout/internal/data"
	"go-layout/internal/dep"
	"go-layout/internal/dto"
	"go-layout/internal/server"
	"go-layout/internal/service"
)

func wireApp() (*App, error) {
	panic(wire.Build(
		dep.ProviderSet,
		conf.ProviderSet,
		data.ProviderSet,
		dto.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
		NewApp,
	))
}
