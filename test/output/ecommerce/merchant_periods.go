package ecommerce

import "time"

type MerchantPeriods struct {
	Id          int       `gorm:"column:id;primaryKey"`
	MerchantId  int       `gorm:"column:merchant_id"`
	CountryCode int       `gorm:"column:country_code"`
	StartDate   time.Time `gorm:"column:start_date"`
	EndDate     time.Time `gorm:"column:end_date"`
	Merchants   *Merchants
}
