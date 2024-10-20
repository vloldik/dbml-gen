package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&public.User{}, &public.Country{}, &ecommerce.Order{}, &ecommerce.Product{}, &ecommerce.ProductTag{}, &ecommerce.Merchant{}, &ecommerce.OrderItem{}, &ecommerce.MerchantPeriod{})
}
