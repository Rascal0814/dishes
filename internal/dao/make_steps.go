package dao

import (
	"context"
	"dishes/order-dishes/internal/dao/model"
)

// Inserts 批量插入
func (dao *MakeStepDao) Inserts(ctx context.Context, m []*model.MakeStep) error {
	return dao.db.WithContext(ctx).Create(m).Error
}

// GetStepsByDishId 获取步骤
func (dao *MakeStepDao) GetStepsByDishId(ctx context.Context, id int64) ([]*model.MakeStep, error) {
	var res = make([]*model.MakeStep, 0)
	err := dao.db.WithContext(ctx).Where("dish_id = ?", id).Order("step").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DelStepsByDishId 删除做法
func (dao *MakeStepDao) DelStepsByDishId(ctx context.Context, dishId int64) error {
	return dao.db.WithContext(ctx).Where("dish_id = ?", dishId).Delete(&model.MakeStep{}).Error
}
