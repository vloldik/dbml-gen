// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	ecommerce "output/ecommerce"
)

type MerchantPeriodRepository struct {
	db *gorm.DB
}

func NewMerchantPeriodRepository(db *gorm.DB) *MerchantPeriodRepository {
	return &MerchantPeriodRepository{db: db}
}
func (r *MerchantPeriodRepository) GetByID(ctx context.Context, id any) (*ecommerce.MerchantPeriod, error) {
	var merchantPeriod ecommerce.MerchantPeriod
	result := r.db.WithContext(ctx).First(&merchantPeriod, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &merchantPeriod, nil
}
func (r *MerchantPeriodRepository) DeleteByID(ctx context.Context, id any) error {
	var merchantPeriod ecommerce.MerchantPeriod
	result := r.db.WithContext(ctx).Delete(&merchantPeriod, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *MerchantPeriodRepository) Create(ctx context.Context, model ecommerce.MerchantPeriod) (*ecommerce.MerchantPeriod, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *MerchantPeriodRepository) List(ctx context.Context, limit int, offset int) ([]*ecommerce.MerchantPeriod, error) {
	var list []*ecommerce.MerchantPeriod
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *MerchantPeriodRepository) Update(ctx context.Context, model ecommerce.MerchantPeriod) (*ecommerce.MerchantPeriod, error) {
	result := r.db.WithContext(ctx).Save(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *MerchantPeriodRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&ecommerce.MerchantPeriod{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
