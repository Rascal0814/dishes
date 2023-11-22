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

type DishesService struct {
	v1.UnimplementedDishesServiceServer

	dao  *dao.Dishes
	log  *blog.Logger
	dto  *dto.Dto
	snow *snowflake.Snowflake
}

func NewDishesService(dao *dao.Dishes, log *blog.Logger, dto *dto.Dto, snow *snowflake.Snowflake) *DishesService {
	return &DishesService{dao: dao, log: log, dto: dto, snow: snow}
}

func (d *DishesService) GetDishes(ctx context.Context, request *v1.GetDishesRequest) (*v1.GetDishesResponse, error) {
	m, err := d.dao.Get(ctx, d.snow.ParseId(request.Id))
	if err != nil {
		_ = d.log.Error("get dish error", err)
		return nil, api.ErrDishesNotFound
	}
	return &v1.GetDishesResponse{Dishes: d.dto.Dishes(m)}, nil
}

func (d *DishesService) DelDishes(ctx context.Context, request *v1.DelDishesRequest) (*v1.DelDishesResponse, error) {
	err := d.dao.Del(ctx, d.snow.ParseId(request.Id))
	if err != nil {
		_ = d.log.Error("del dish error", err)
		return nil, api.ErrDishesDelete
	}
	return &v1.DelDishesResponse{}, err
}

func (d *DishesService) UpdateDishes(ctx context.Context, request *v1.UpdateDishesRequest) (*v1.UpdateDishesResponse, error) {
	info, err := d.dao.Get(ctx, d.snow.ParseId(request.Id))
	if err != nil {
		_ = d.log.Error("get dish error", err)
		return nil, api.ErrDishesNotFound
	}

	info.Name = request.Name
	info.Remark = request.Remark
	info.Logo = request.Logo
	info, err = d.dao.Update(ctx, info)
	if err != nil {
		_ = d.log.Error("update dish error", err)
		return nil, api.ErrDishesUpdate
	}
	return &v1.UpdateDishesResponse{
		Dishes: d.dto.Dishes(info),
	}, err
}

func (d *DishesService) ListDishes(ctx context.Context, request *v1.ListDishesRequest) (*v1.ListDishesResponse, error) {
	list, total, err := d.dao.List(ctx, &dao.ListDishesParams{
		Pager: orm.Pager{
			PageIndex: request.PageIndex,
			PageSize:  request.PageSize,
		},
		Name: request.Name,
	})
	if err != nil {
		_ = d.log.Error("get dish list error", err)
		return nil, err
	}

	var resp = &v1.ListDishesResponse{
		Dishes: make([]*v1.Dishes, 0),
		Total:  total,
	}
	for _, dish := range list {
		resp.Dishes = append(resp.Dishes, d.dto.Dishes(dish))
	}

	return resp, nil
}

func (d *DishesService) CreateDishes(ctx context.Context, request *v1.CreateDishesRequest) (*v1.CreateDishesResponse, error) {
	m, err := d.dao.Insert(ctx, func(dish *model.Dish) error {
		dish.ID = d.snow.GenId()
		dish.Name = request.Name
		dish.Remark = request.Remark
		dish.Logo = request.Logo
		return nil
	})

	if err != nil {
		_ = d.log.Error("make dishes error", err)
		return nil, api.ErrDishesInsert
	}
	return &v1.CreateDishesResponse{Dishes: d.dto.Dishes(m)}, nil
}
