// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	opt "output/opt"
	public "output/public"
)

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{db: db}
}
func (r *CountryRepository) GetByCode(ctx context.Context, code any, opts ...any) (*public.Country, error) {
	var country public.Country
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).First(&country, "code = ?", code)
	if result.Error != nil {
		return nil, result.Error
	}
	return &country, nil
}
func (r *CountryRepository) DeleteByCode(ctx context.Context, code any, opts ...any) error {
	var country public.Country
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Delete(&country, "code = ?", code)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *CountryRepository) Create(ctx context.Context, model public.Country, opts ...any) (*public.Country, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *CountryRepository) List(ctx context.Context, limit int, offset int, opts ...any) ([]*public.Country, error) {
	var list []*public.Country
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *CountryRepository) Update(ctx context.Context, model public.Country, opts ...any) (*public.Country, error) {
	result := opt.ApplyOptions(r.db.WithContext(ctx), opts...).Updates(&model)
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
