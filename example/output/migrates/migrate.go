package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.Order{}, &ecommerce.Product{}, &ecommerce.MerchantPeriod{}, &public.User{}, &public.Country{}, &ecommerce.ProductTag{}, &ecommerce.Merchant{}, &ecommerce.OrderItem{})
}
