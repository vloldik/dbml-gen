// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
	opt "output/opt"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}
func (r *OrderItemRepository) Create(ctx context.Context, model ecommerce.OrderItem, opts ...any) (*ecommerce.OrderItem, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *OrderItemRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*ecommerce.OrderItem, error) {
	var list []*ecommerce.OrderItem
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *OrderItemRepository) Update(ctx context.Context, model ecommerce.OrderItem, opts ...any) (*ecommerce.OrderItem, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *OrderItemRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.OrderItem{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *OrderItemRepository) GetDB() *gorm.DB {
	return r.db
}
