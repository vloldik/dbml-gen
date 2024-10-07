package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&public.Country{}, &public.User{}, &ecommerce.Order{}, &ecommerce.Merchant{}, &ecommerce.MerchantPeriod{}, &ecommerce.Product{}, &ecommerce.ProductTag{}, &ecommerce.OrderItem{})
}
