//go:build wireinject
// +build wireinject

package test

import (
	"dishes/order-dishes/config"
	"dishes/order-dishes/internal/dao"
	"dishes/order-dishes/internal/dto"
	"dishes/order-dishes/internal/service"
	bc "github.com/Rascal0814/boot/config"
	blog "github.com/Rascal0814/boot/log"
	"github.com/Rascal0814/boot/snowflake"

	"github.com/google/wire"
)

type TestApp struct {
	Dishes   *service.DishesService
	Orders   *service.OrderService
	MakeStep *service.MakeStepsService
	Config   *bc.Config
	log      *blog.Logger
}

func newTestApp(Dishes *service.DishesService, Orders *service.OrderService, MakeStep *service.MakeStepsService, Config *bc.Config, log *blog.Logger) (*TestApp, error) {
	return &TestApp{
		Dishes:   Dishes,
		Orders:   Orders,
		MakeStep: MakeStep,
		Config:   Config,
		log:      log,
	}, nil
}

var (
	app *TestApp
	err error
)

func init() {
	app, err = wireApp("test")
	if err != nil {
		panic(err)
	}
}

// wireApp init kratos application.
func wireApp(serviceName string) (*TestApp, error) {
	panic(wire.Build(
		blog.NewLogger,
		snowflake.NewSnowflake,
		service.ProviderSet,
		config.ProviderSet,
		dao.ProviderSet,
		dto.ProviderSet,
		bc.DefaultDB,
		newTestApp))
}
