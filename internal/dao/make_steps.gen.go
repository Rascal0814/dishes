// Code generated by boot. DO NOT EDIT.

package dao

import (
	"context"

	model "dishes/order-dishes/internal/dao/model"
	"github.com/Rascal0814/boot/log"
	"github.com/Rascal0814/boot/orm"
	"gorm.io/gorm"
)

type MakeStepDao struct {
	db  *gorm.DB
	log *log.Logger
}

func NewMakeStepDao(db *gorm.DB, log *log.Logger) *MakeStepDao {
	return &MakeStepDao{db: db, log: log}
}

func (dao *MakeStepDao) Insert(ctx context.Context, f func(m *model.MakeStep) error) (*model.MakeStep, error) {
	var m model.MakeStep
	err := f(&m)
	if err != nil {
		return nil, dao.log.Error("insert MakeStep error", err)
	}
	err = dao.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		return nil, dao.log.Error("insert MakeStep error", err)
	}
	return &m, err
}

func (dao *MakeStepDao) Get(ctx context.Context, id int64) (*model.MakeStep, error) {
	var m = &model.MakeStep{}
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (dao *MakeStepDao) Del(ctx context.Context, id int64) error {
	err := dao.db.WithContext(ctx).Where("id = ?", id).Delete(&model.MakeStep{}).Error
	if err != nil {
		return dao.log.Error("del MakeStep error", err)
	}
	return nil
}

func (dao *MakeStepDao) List(ctx context.Context,p *orm.Pager) ([]*model.MakeStep, int64, error) {
	var res = make([]*model.MakeStep, 0)
	var total int64
	err := dao.db.WithContext(ctx).Model(&model.MakeStep{}).Count(&total).Limit(int((p.PageIndex) * p.PageSize)).Offset(int(p.PageIndex - 1)).Find(&res).Error
	if err != nil {
		return nil, 0, dao.log.Error("get MakeStep list error", err)
	}
	return res, total, nil
}

func (dao *MakeStepDao) Update(ctx context.Context, m *model.MakeStep) (*model.MakeStep, error) {
	err := dao.db.WithContext(ctx).Updates(m).Error
	if err != nil {
		return nil, dao.log.Error("update MakeStep error", err)
	}
	return m, nil
}
