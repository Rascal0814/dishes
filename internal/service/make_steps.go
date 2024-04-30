package service

import (
	"context"
	"dishes/order-dishes/api"
	v1 "dishes/order-dishes/api/v1"
	"dishes/order-dishes/internal/dao"
	"dishes/order-dishes/internal/dao/model"
	"dishes/order-dishes/internal/dto"
	blog "github.com/Rascal0814/boot/log"
	"github.com/Rascal0814/boot/snowflake"
)

type MakeStepsService struct {
	v1.UnimplementedMakeStepsServiceServer

	dao  *dao.MakeStepDao
	log  *blog.Logger
	dto  *dto.Dto
	snow *snowflake.Snowflake
}

func NewMakeStepsService(dao *dao.MakeStepDao, log *blog.Logger, dto *dto.Dto, snow *snowflake.Snowflake) *MakeStepsService {
	return &MakeStepsService{dao: dao, log: log, dto: dto, snow: snow}
}

func (m MakeStepsService) CreateMakeSteps(ctx context.Context, request *v1.CreateMakeStepsRequest) (*v1.CreateMakeStepsResponse, error) {
	var models = make([]*model.MakeStep, 0, len(request.Steps))
	for _, step := range request.Steps {
		models = append(models, &model.MakeStep{
			ID:      m.snow.GenId(),
			DishID:  m.snow.ParseId(request.DishId),
			Step:    step.Step,
			Content: step.Content,
			Creator: 1,
		})
	}
	err := m.dao.Inserts(ctx, models)
	if err != nil {
		_ = m.log.Error("inserts make steps error", err)
		return nil, api.ErrInsert
	}
	var res = &v1.CreateMakeStepsResponse{
		Steps: make([]*v1.MakeStep, 0),
	}
	for _, v := range models {
		res.Steps = append(res.Steps, m.dto.MakeSteps(v))
	}

	return res, nil
}

func (m MakeStepsService) GetMakeSteps(ctx context.Context, request *v1.GetMakeStepsRequest) (*v1.GetMakeStepsResponse, error) {
	s, err := m.dao.Get(ctx, m.snow.ParseId(request.Id))
	if err != nil {
		_ = m.log.Error("get make steps error", err)
		return nil, api.ErrNotFound
	}
	return &v1.GetMakeStepsResponse{MakeStep: m.dto.MakeSteps(s)}, nil
}

func (m MakeStepsService) DelMakeSteps(ctx context.Context, request *v1.DelMakeStepsRequest) (*v1.DelMakeStepsResponse, error) {
	err := m.dao.Del(ctx, m.snow.ParseId(request.Id))
	if err != nil {
		_ = m.log.Error("del make steps error", err)
		return nil, api.ErrDelete
	}
	return &v1.DelMakeStepsResponse{}, err
}

func (m MakeStepsService) DelMakeStepsByDishId(ctx context.Context, request *v1.DelMakeStepsByDishIdRequest) (*v1.DelMakeStepsByDishIdResponse, error) {
	err := m.dao.DelStepsByDishId(ctx, m.snow.ParseId(request.DishId))
	if err != nil {
		_ = m.log.Error("del make steps by dish_id error", err)
		return nil, api.ErrDelete
	}

	return &v1.DelMakeStepsByDishIdResponse{}, nil
}

func (m MakeStepsService) UpdateMakeSteps(ctx context.Context, request *v1.UpdateMakeStepsRequest) (*v1.UpdateMakeStepsResponse, error) {
	info, err := m.dao.Get(ctx, m.snow.ParseId(request.Id))
	if err != nil {
		_ = m.log.Error("get make step error", err)
		return nil, api.ErrNotFound
	}
	info.Content = request.Content
	info, err = m.dao.Update(ctx, info)
	if err != nil {
		_ = m.log.Error("update make step error", err)
		return nil, api.ErrUpdate
	}
	return &v1.UpdateMakeStepsResponse{MakeStep: m.dto.MakeSteps(info)}, err
}

func (m MakeStepsService) GetDishesMakeSteps(ctx context.Context, request *v1.DishesMakeStepsRequest) (*v1.DishesMakeStepsResponse, error) {
	list, err := m.dao.GetStepsByDishId(ctx, m.snow.ParseId(request.DishId))
	if err != nil {
		_ = m.log.Error("get dish make step error", err)
		return nil, api.ErrNotFound
	}
	var res = &v1.DishesMakeStepsResponse{
		Steps: make([]*v1.MakeStep, 0, len(list)),
	}
	for _, step := range list {
		res.Steps = append(res.Steps, m.dto.MakeSteps(step))
	}
	return res, nil
}
