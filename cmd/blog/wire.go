//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"blog/configs"
	"blog/internal/biz"
	"blog/internal/infra"
	"blog/internal/server"
	"blog/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*configs.ApplicationConfig, *configs.DBConfig, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			infra.ProviderSet,
			biz.ProviderSet,
			service.ProviderSet,
			newApp,
			wire.Bind(new(biz.GreeterRepo), new(*infra.GreeterRepo)),
		),
	)
}
