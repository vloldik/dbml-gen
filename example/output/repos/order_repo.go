// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
	opt "output/opt"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
func (r *OrderRepository) GetByID(ctx context.Context, id any, opts ...any) (*ecommerce.Order, error) {
	var order ecommerce.Order
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
func (r *OrderRepository) DeleteByID(ctx context.Context, id any, opts ...any) error {
	var order ecommerce.Order
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Delete(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *OrderRepository) Create(ctx context.Context, model ecommerce.Order, opts ...any) (*ecommerce.Order, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *OrderRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*ecommerce.Order, error) {
	var list []*ecommerce.Order
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *OrderRepository) Update(ctx context.Context, model ecommerce.Order, opts ...any) (*ecommerce.Order, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *OrderRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.Order{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *OrderRepository) GetDB() *gorm.DB {
	return r.db
}
