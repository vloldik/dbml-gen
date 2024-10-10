package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.Order{}, &ecommerce.Product{}, &ecommerce.ProductTag{}, &ecommerce.MerchantPeriod{}, &ecommerce.Merchant{}, &public.User{}, &public.Country{}, &ecommerce.OrderItem{})
}
