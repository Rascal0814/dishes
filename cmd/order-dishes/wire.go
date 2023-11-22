//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"dishes/order-dishes/config"
	"dishes/order-dishes/internal/dao"
	"dishes/order-dishes/internal/dto"
	"dishes/order-dishes/internal/server"
	"dishes/order-dishes/internal/service"
	bc "github.com/Rascal0814/boot/config"
	"github.com/Rascal0814/boot/kratos/depend"
	blog "github.com/Rascal0814/boot/log"
	"github.com/Rascal0814/boot/snowflake"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(serviceName string) (*kratos.App, func(), error) {
	panic(wire.Build(
		//log.ProvideLogger,
		depend.NewConsulRegistrar,
		bc.DefaultDB,
		blog.NewLogger,
		snowflake.NewSnowflake,
		server.ProviderSet,
		service.ProviderSet,
		config.ProviderSet,
		dao.ProviderSet,
		dto.ProviderSet,
		newApp))
}
