package service

import (
	"context"
	"dishes/order-dishes/api"
	v1 "dishes/order-dishes/api/v1"
	"dishes/order-dishes/internal/dao"
	"dishes/order-dishes/internal/dao/model"
	"dishes/order-dishes/internal/dto"
	blog "github.com/Rascal0814/boot/log"
	"github.com/Rascal0814/boot/orm"
	"github.com/Rascal0814/boot/snowflake"
)

type OrderService struct {
	v1.UnimplementedOrdersServiceServer

	dao  *dao.OrderDao
	log  *blog.Logger
	dto  *dto.Dto
	snow *snowflake.Snowflake
}

func (o OrderService) CreateOrders(ctx context.Context, request *v1.CreateOrdersRequest) (*v1.CreateOrdersResponse, error) {
	m, err := o.dao.Insert(ctx, func(order *model.Order) error {
		order.ID = o.snow.GenId()
		order.DishID = o.snow.ParseId(request.DishId)
		order.Remark = request.Remake
		order.Creator = 1 // todo access user system
		return nil
	})

	if err != nil {
		_ = o.log.Error("make order error", err)
		return nil, api.ErrInsert
	}
	return &v1.CreateOrdersResponse{Order: o.dto.Orders(m)}, nil
}

func (o OrderService) GetOrders(ctx context.Context, request *v1.GetOrdersRequest) (*v1.GetOrdersResponse, error) {
	m, err := o.dao.Get(ctx, o.snow.ParseId(request.Id))
	if err != nil {
		_ = o.log.Error("get order error", err)
		return nil, api.ErrNotFound
	}
	return &v1.GetOrdersResponse{Order: o.dto.Orders(m)}, nil
}

func (o OrderService) DelOrders(ctx context.Context, request *v1.DelOrdersRequest) (*v1.DelOrdersResponse, error) {
	err := o.dao.Del(ctx, o.snow.ParseId(request.Id))
	if err != nil {
		_ = o.log.Error("del order error", err)
		return nil, api.ErrDelete
	}
	return &v1.DelOrdersResponse{}, err
}

func (o OrderService) UpdateOrders(ctx context.Context, request *v1.UpdateOrdersRequest) (*v1.UpdateOrdersResponse, error) {
	info, err := o.dao.Get(ctx, o.snow.ParseId(request.Id))
	if err != nil {
		_ = o.log.Error("get order error", err)
		return nil, api.ErrNotFound
	}

	info.DishID = o.snow.ParseId(request.DishId)
	info, err = o.dao.Update(ctx, info)
	if err != nil {
		_ = o.log.Error("update order error", err)
		return nil, api.ErrUpdate
	}
	return &v1.UpdateOrdersResponse{Order: o.dto.Orders(info)}, err
}

func (o OrderService) ListOrders(ctx context.Context, request *v1.ListOrdersRequest) (*v1.ListOrdersResponse, error) {
	list, total, err := o.dao.List(ctx, &orm.Pager{
		PageIndex: request.PageIndex,
		PageSize:  request.PageSize,
	})
	if err != nil {
		_ = o.log.Error("get order list error", err)
		return nil, api.ErrList
	}

	var resp = &v1.ListOrdersResponse{
		Orders: make([]*v1.Order, 0),
		Total:  total,
	}
	for _, d := range list {
		resp.Orders = append(resp.Orders, o.dto.Orders(d))
	}

	return resp, nil
}

func NewOrderService(dao *dao.OrderDao, log *blog.Logger, dto *dto.Dto, snow *snowflake.Snowflake) *OrderService {
	return &OrderService{dao: dao, log: log, dto: dto, snow: snow}
}
