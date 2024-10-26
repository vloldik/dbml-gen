package migrates

import (
	gorm "gorm.io/gorm"
	ecommerce "output/ecommerce"
	public "output/public"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(&ecommerce.ProductTag{}, &ecommerce.MerchantPeriod{}, &ecommerce.Merchant{}, &public.Country{}, &ecommerce.Order{}, &public.User{}, &ecommerce.OrderItem{}, &ecommerce.Product{})
}
