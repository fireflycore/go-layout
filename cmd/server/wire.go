//go:build wireinject
// +build wireinject

package main

import (
	"go-layout/internal/biz"
	"go-layout/internal/conf"
	"go-layout/internal/data"
	"go-layout/internal/dep"
	"go-layout/internal/dto"
	"go-layout/internal/server"
	"go-layout/internal/service"

	"github.com/google/wire"
)

// wireApp 使用 Wire 把各层 ProviderSet 组装成应用根对象。
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
