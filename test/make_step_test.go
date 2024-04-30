package test

import (
	"context"
	v1 "dishes/order-dishes/api/v1"
	"fmt"
	"testing"
)

func TestMakeStepService(t *testing.T) {
	ctx := context.Background()
	dishes, err := app.Dishes.CreateDishes(ctx, &v1.CreateDishesRequest{Name: "test dish name"})
	if err != nil {
		t.Fatal(err)
	}
	step, err := app.MakeStep.CreateMakeSteps(ctx, &v1.CreateMakeStepsRequest{
		DishId: dishes.Dishes.Id,
		Steps: []*v1.CreateMakeStep{
			{
				Step:    1,
				Content: "boiled water",
			},
			{
				Step:    2,
				Content: "fried and dried",
			},
			{
				Step:    3,
				Content: "take out of the pot",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.MakeStep.UpdateMakeSteps(ctx, &v1.UpdateMakeStepsRequest{
		Id:      step.Steps[0].Id,
		DishId:  step.Steps[0].DishId,
		Content: "boiled water update",
	})
	if err != nil {
		t.Fatal(err)
	}
	listRes, err := app.MakeStep.GetDishesMakeSteps(ctx, &v1.DishesMakeStepsRequest{DishId: dishes.Dishes.GetId()})
	if err != nil {
		t.Fatal(err)
	}
	app.log.Info(fmt.Sprintf("get dishes[%s] make steps:", dishes.Dishes.Id))

	for _, v := range listRes.Steps {
		app.log.Info("", "step", v.Step, "content", v.Content)
	}

	getStep, err := app.MakeStep.GetMakeSteps(ctx, &v1.GetMakeStepsRequest{Id: step.Steps[0].Id})
	if err != nil {
		t.Fatal(err)
	}

	if getStep.MakeStep.Content != "boiled water update" {
		app.log.Debug(fmt.Sprintf("get step[%s] update status not equal，please check update func", getStep.MakeStep.Id))
		t.Fatal("check update func")
	}

	_, err = app.MakeStep.DelMakeSteps(ctx, &v1.DelMakeStepsRequest{Id: step.Steps[0].Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.MakeStep.DelMakeStepsByDishId(ctx, &v1.DelMakeStepsByDishIdRequest{DishId: dishes.Dishes.Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.Dishes.DelDishes(ctx, &v1.DelDishesRequest{Id: dishes.Dishes.Id})
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.MakeStep.DelMakeStepsByDishId(ctx, &v1.DelMakeStepsByDishIdRequest{DishId: dishes.Dishes.GetId()})
	if err != nil {
		t.Fatal(err)
	}

	ss, err := app.MakeStep.GetDishesMakeSteps(ctx, &v1.DishesMakeStepsRequest{DishId: dishes.Dishes.GetId()})
	if err != nil {
		t.Fatal(err)
	}

	if len(ss.Steps) > 0 {
		app.log.Debug(fmt.Sprintf("get dish[%s] steps，please check update func", dishes.Dishes.Id))
		t.Fatal("check DelMakeStepsByDishId func")
	}

}
