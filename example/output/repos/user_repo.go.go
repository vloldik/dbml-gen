// Code generated from DBML. DO NOT EDIT
package repos

import (
	"context"
	"gorm.io/gorm"
	public "output/public"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) GetByID(ctx context.Context, id any) (*public.User, error) {
	var user public.User
	result := r.db.WithContext(ctx).First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *UserRepository) DeleteByID(ctx context.Context, id any) error {
	var user public.User
	result := r.db.WithContext(ctx).Delete(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *UserRepository) Create(ctx context.Context, model public.User) (*public.User, error) {
	result := r.db.WithContext(ctx).Create(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *UserRepository) List(ctx context.Context, limit int, offset int) ([]*public.User, error) {
	var list []*public.User
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
func (r *UserRepository) Update(ctx context.Context, model public.User) (*public.User, error) {
	result := r.db.WithContext(ctx).Updates(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model, nil
}
func (r *UserRepository) TotalCount(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&public.User{}).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}
	return count, nil
}
func (r *UserRepository) GetDB() *gorm.DB {
	return r.db
}
