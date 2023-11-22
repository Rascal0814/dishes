package main

import (
	"github.com/go-kratos/kratos/v2/registry"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name  = "order-dishes"
	id, _ = os.Hostname()
)

func newApp(r registry.Registrar, gs *grpc.Server, hs *http.Server) (*kratos.App, error) {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	), nil
}

func main() {
	app, cleanup, err := wireApp(Name)
	if err != nil {
		panic(err)
	}
	defer func() { cleanup() }()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
