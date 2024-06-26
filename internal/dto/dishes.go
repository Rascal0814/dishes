package dto

import (
	v1 "dishes/order-dishes/api/v1"
	"dishes/order-dishes/internal/dao/model"
	"github.com/bwmarrin/snowflake"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Dto struct {
}

func NewDto() *Dto {
	return &Dto{}
}

func (d *Dto) Dishes(m *model.Dish) *v1.Dishes {
	return &v1.Dishes{
		Id:         snowflake.ID(m.ID).String(),
		Name:       m.Name,
		Logo:       m.Logo,
		Remark:     m.Remark,
		CreateTime: timestamppb.New(m.CreatedAt),
		UpdateTime: timestamppb.New(m.UpdatedAt),
	}
}

func (d *Dto) Orders(m *model.Order) *v1.Order {
	return &v1.Order{
		Id:         snowflake.ID(m.ID).String(),
		DishId:     snowflake.ID(m.DishID).String(),
		Status:     v1.OrderStatus_name[m.Status],
		Creator:    "",
		Remark:     m.Remark,
		CreateTime: timestamppb.New(m.CreatedAt),
		UpdateTime: timestamppb.New(m.UpdatedAt),
	}
}

func (d *Dto) MakeSteps(m *model.MakeStep) *v1.MakeStep {
	return &v1.MakeStep{
		Id:         snowflake.ID(m.ID).String(),
		DishId:     snowflake.ID(m.DishID).String(),
		Step:       m.Step,
		Content:    m.Content,
		Creator:    m.Creator,
		CreateTime: timestamppb.New(m.CreatedAt),
		UpdateTime: timestamppb.New(m.UpdatedAt),
	}
}
