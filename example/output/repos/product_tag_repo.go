// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
	opt "output/opt"
)

type ProductTagRepository struct {
	db *gorm.DB
}

func NewProductTagRepository(db *gorm.DB) *ProductTagRepository {
	return &ProductTagRepository{db: db}
}
func (r *ProductTagRepository) GetByID(ctx context.Context, id any, opts ...any) (*ecommerce.ProductTag, error) {
	var productTag ecommerce.ProductTag
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).First(&productTag, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &productTag, nil
}
func (r *ProductTagRepository) DeleteByID(ctx context.Context, id any, opts ...any) error {
	var productTag ecommerce.ProductTag
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Delete(&productTag, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *ProductTagRepository) Create(ctx context.Context, model ecommerce.ProductTag, opts ...any) (*ecommerce.ProductTag, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *ProductTagRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*ecommerce.ProductTag, error) {
	var list []*ecommerce.ProductTag
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *ProductTagRepository) Update(ctx context.Context, model ecommerce.ProductTag, opts ...any) (*ecommerce.ProductTag, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *ProductTagRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.ProductTag{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *ProductTagRepository) GetDB() *gorm.DB {
	return r.db
}
