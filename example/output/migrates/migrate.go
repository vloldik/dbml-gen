package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.Order{}, &ecommerce.Product{}, &ecommerce.Merchant{}, &ecommerce.OrderItem{}, &ecommerce.ProductTag{}, &ecommerce.MerchantPeriod{}, &public.User{}, &public.Country{})
}
