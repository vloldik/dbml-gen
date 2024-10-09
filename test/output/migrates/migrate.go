package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.Product{}, &ecommerce.ProductTag{}, &ecommerce.Merchant{}, &public.Country{}, &ecommerce.MerchantPeriod{}, &public.User{}, &ecommerce.OrderItem{}, &ecommerce.Order{})
}
