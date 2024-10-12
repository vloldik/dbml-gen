package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.MerchantPeriod{}, &ecommerce.Merchant{}, &ecommerce.ProductTag{}, &public.User{}, &public.Country{}, &ecommerce.OrderItem{}, &ecommerce.Order{}, &ecommerce.Product{})
}
