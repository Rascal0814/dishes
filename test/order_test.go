package test

import (
	"context"
	v1 "dishes/order-dishes/api/v1"
	"fmt"
	"testing"
)

func TestOrderService(t *testing.T) {
	ctx := context.Background()
	dishes, err := app.Dishes.CreateDishes(ctx, &v1.CreateDishesRequest{Name: "test dish name"})
	if err != nil {
		t.Fatal(err)
	}
	order, err := app.Orders.CreateOrders(ctx, &v1.CreateOrdersRequest{DishId: dishes.Dishes.Id})
	if err != nil {
		t.Fatal(err)
	}
	app.log.Info("create order success", "id", order.Order.Id)

	_, err = app.Orders.UpdateOrders(ctx, &v1.UpdateOrdersRequest{
		Id:     order.Order.Id,
		DishId: dishes.Dishes.Id,
		Status: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	listRes, err := app.Orders.ListOrders(ctx, &v1.ListOrdersRequest{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	app.log.Info("get list order", "total", listRes.Total)
	getDishes, err := app.Orders.GetOrders(ctx, &v1.GetOrdersRequest{Id: order.Order.GetId()})
	if err != nil {
		t.Fatal(err)
	}

	if getDishes.Order.Status != v1.OrderStatus_name[1] {
		app.log.Debug(fmt.Sprintf("get order[%s] update status not equal，please check update func", order.Order.Id))
		t.Fatal("check update func")
	}

	_, err = app.Orders.DelOrders(ctx, &v1.DelOrdersRequest{Id: order.Order.Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.Dishes.DelDishes(ctx, &v1.DelDishesRequest{Id: dishes.Dishes.Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.Orders.GetOrders(ctx, &v1.GetOrdersRequest{Id: order.Order.GetId()})
	if err == nil {
		app.log.Debug(fmt.Sprintf("get order[%s]，please check del func", order.Order.Id))
		t.Fatal("check del func")
	}

}
