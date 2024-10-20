package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.ProductTag{}, &ecommerce.Merchant{}, &public.User{}, &ecommerce.Order{}, &ecommerce.MerchantPeriod{}, &public.Country{}, &ecommerce.OrderItem{}, &ecommerce.Product{})
}
