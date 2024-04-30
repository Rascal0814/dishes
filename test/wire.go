//go:build wireinject
// +build wireinject

package test

import (
	"dishes/order-dishes/config"
	"dishes/order-dishes/internal/dao"
	"dishes/order-dishes/internal/dto"
	"dishes/order-dishes/internal/service"
	"flag"
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
	app      *TestApp
	err      error
	flagConf string
)

func init() {
	app, err = wireApp(blog.ServiceName("test"), bc.CPath(flagConf))
	if err != nil {
		panic(err)
	}
	flag.StringVar(&flagConf, "conf", "../config", "config path, eg: -conf config.yaml")
}

// wireApp init kratos application.
func wireApp(serviceName blog.ServiceName, confPath bc.CPath) (*TestApp, error) {
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
