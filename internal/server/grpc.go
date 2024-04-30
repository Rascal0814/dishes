package server

import (
	v1 "dishes/order-dishes/api/v1"
	"dishes/order-dishes/internal/service"

	"github.com/Rascal0814/boot/config"
	kratosmiddleware "github.com/Rascal0814/boot/kratos/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *config.Server, dishes *service.DishesService, order *service.OrderService, makeStep *service.MakeStepsService) (*grpc.Server, error) {
	var opts = kratosmiddleware.DefaultGrpcMiddleWare
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDishesServiceServer(srv, dishes)
	v1.RegisterOrdersServiceServer(srv, order)
	v1.RegisterMakeStepsServiceServer(srv, makeStep)
	return srv, nil
}
