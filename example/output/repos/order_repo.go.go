// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
func (r *OrderRepository) GetByID(ctx context.Context, id any) (*ecommerce.Order, error) {
	var order ecommerce.Order
	result := r.db.WithContext(ctx).First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}
func (r *OrderRepository) DeleteByID(ctx context.Context, id any) error {
	var order ecommerce.Order
	result := r.db.WithContext(ctx).Delete(&order, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *OrderRepository) Create(ctx context.Context, model ecommerce.Order) (*ecommerce.Order, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *OrderRepository) List(ctx context.Context, limit int, offset int) ([]*ecommerce.Order, error) {
	var list []*ecommerce.Order
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *OrderRepository) Update(ctx context.Context, model ecommerce.Order) (*ecommerce.Order, error) {
	result := r.db.WithContext(ctx).Updates(&model)
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
