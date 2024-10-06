package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.OrderItem{}, &ecommerce.ProductTag{}, &ecommerce.MerchantPeriod{}, &ecommerce.Product{}, &ecommerce.Merchant{}, &public.User{}, &public.Country{}, &ecommerce.Order{})
}
