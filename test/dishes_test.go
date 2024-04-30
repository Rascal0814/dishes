package test

import (
	"context"
	v1 "dishes/order-dishes/api/v1"
	"fmt"
	"testing"
)

func TestDishesService(t *testing.T) {
	ctx := context.Background()
	dishes, err := app.Dishes.CreateDishes(ctx, &v1.CreateDishesRequest{Name: "test dish name"})
	if err != nil {
		t.Fatal(err)
	}
	app.log.Info("create dishes success", "id", dishes.Dishes.Id)

	_, err = app.Dishes.UpdateDishes(ctx, &v1.UpdateDishesRequest{
		Id:     dishes.Dishes.Id,
		Name:   dishes.Dishes.Name,
		Logo:   "",
		Remark: "test update dishes",
	})
	if err != nil {
		t.Fatal(err)
	}
	listRes, err := app.Dishes.ListDishes(ctx, &v1.ListDishesRequest{
		Name:      "",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatal(err)
	}
	app.log.Info("get list dishes", "total", listRes.Total)
	getDishes, err := app.Dishes.GetDishes(ctx, &v1.GetDishesRequest{Id: dishes.Dishes.GetId()})
	if err != nil {
		t.Fatal(err)
	}

	if getDishes.Dishes.Remark != "test update dishes" {
		app.log.Debug(fmt.Sprintf("get dishes[%s] update remark not equal，please check update func", getDishes.Dishes.Id))
		t.Fatal("check del func")
	}

	_, err = app.Dishes.DelDishes(ctx, &v1.DelDishesRequest{Id: dishes.Dishes.Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.Dishes.GetDishes(ctx, &v1.GetDishesRequest{Id: dishes.Dishes.GetId()})
	if err == nil {
		app.log.Debug(fmt.Sprintf("get dishes[%s]，please check del func", getDishes.Dishes.Id))
		t.Fatal("check del func")
	}
}
