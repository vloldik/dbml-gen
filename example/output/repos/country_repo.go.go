// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	public "output/public"
)

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{db: db}
}
func (r *CountryRepository) GetByCode(ctx context.Context, code any) (*public.Country, error) {
	var country public.Country
	result := r.db.WithContext(ctx).First(&country, "code = ?", code)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}
func (r *CountryRepository) DeleteByCode(ctx context.Context, code any) error {
	var country public.Country
	result := r.db.WithContext(ctx).Delete(&country, "code = ?", code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *CountryRepository) Create(ctx context.Context, model public.Country) (*public.Country, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *CountryRepository) List(ctx context.Context, limit int, offset int) ([]*public.Country, error) {
	var list []*public.Country
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *CountryRepository) Update(ctx context.Context, model public.Country) (*public.Country, error) {
	result := r.db.WithContext(ctx).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *CountryRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&public.Country{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *CountryRepository) GetDB() *gorm.DB {
	return r.db
}
