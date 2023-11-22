package dao

import (
	"context"
	"dishes/order-dishes/internal/dao/model"
	"github.com/Rascal0814/boot/orm"

	blog "github.com/Rascal0814/boot/log"
	"gorm.io/gorm"
)

type Dishes struct {
	db  *gorm.DB
	log *blog.Logger
}

func NewDishes(db *gorm.DB, log *blog.Logger) *Dishes {
	return &Dishes{db: db, log: log}
}

func (d *Dishes) Insert(ctx context.Context, f func(dish *model.Dish) error) (*model.Dish, error) {
	var m model.Dish
	err := f(&m)
	if err != nil {
		return nil, d.log.Error("insert dish error", err)
	}
	err = d.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		return nil, d.log.Error("insert dish error", err)
	}
	return &m, err
}

func (d *Dishes) Get(ctx context.Context, id int64) (*model.Dish, error) {
	var m = &model.Dish{}
	err := d.db.WithContext(ctx).Where("id = ?", id).Find(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d *Dishes) Del(ctx context.Context, id int64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Dish{}).Error
	if err != nil {
		return d.log.Error("del dish error", err)
	}
	return nil
}

func (d *Dishes) List(ctx context.Context, p *ListDishesParams) ([]*model.Dish, int64, error) {
	var res = make([]*model.Dish, 0)
	var total int64
	err := d.db.WithContext(ctx).Model(&model.Dish{}).Where("name like ?", "%"+p.Name+"%").Count(&total).Limit(int((p.PageIndex - 1) * p.PageSize)).Offset(int(p.PageIndex - 1)).Find(&res).Error
	if err != nil {
		return nil, 0, d.log.Error("get dish list error", err)
	}
	return res, total, nil
}

func (d *Dishes) Update(ctx context.Context, m *model.Dish) (*model.Dish, error) {
	err := d.db.WithContext(ctx).Updates(m).Error
	if err != nil {
		return nil, d.log.Error("update dish error", err)
	}
	return m, nil
}

type ListDishesParams struct {
	orm.Pager
	Name string
}
