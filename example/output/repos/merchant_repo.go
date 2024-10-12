// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
	opt "output/opt"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}
func (r *MerchantRepository) GetByID(ctx context.Context, id any, opts ...any) (*ecommerce.Merchant, error) {
	var merchant ecommerce.Merchant
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).First(&merchant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &merchant, nil
}
func (r *MerchantRepository) DeleteByID(ctx context.Context, id any, opts ...any) error {
	var merchant ecommerce.Merchant
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Delete(&merchant, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *MerchantRepository) Create(ctx context.Context, model ecommerce.Merchant, opts ...any) (*ecommerce.Merchant, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *MerchantRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*ecommerce.Merchant, error) {
	var list []*ecommerce.Merchant
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *MerchantRepository) Update(ctx context.Context, model ecommerce.Merchant, opts ...any) (*ecommerce.Merchant, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *MerchantRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.Merchant{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *MerchantRepository) GetDB() *gorm.DB {
	return r.db
}
