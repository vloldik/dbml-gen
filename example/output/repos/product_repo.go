// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
	opt "output/opt"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
func (r *ProductRepository) GetByID(ctx context.Context, id any, opts ...any) (*ecommerce.Product, error) {
	var product ecommerce.Product
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
func (r *ProductRepository) DeleteByID(ctx context.Context, id any, opts ...any) error {
	var product ecommerce.Product
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Delete(&product, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *ProductRepository) Create(ctx context.Context, model ecommerce.Product, opts ...any) (*ecommerce.Product, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *ProductRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*ecommerce.Product, error) {
	var list []*ecommerce.Product
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *ProductRepository) Update(ctx context.Context, model ecommerce.Product, opts ...any) (*ecommerce.Product, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *ProductRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.Product{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *ProductRepository) GetDB() *gorm.DB {
	return r.db
}
