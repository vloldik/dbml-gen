package ecommerce

import "time"

type MerchantPeriod struct {
	Id          int       `gorm:"column:id;primaryKey"`
	MerchantId  int       `gorm:"column:merchant_id"`
	CountryCode int       `gorm:"column:country_code"`
	StartDate   time.Time `gorm:"column:start_date;type:DATETIME"`
	EndDate     time.Time `gorm:"column:end_date;type:DATETIME"`
	Merchant    *Merchant `gorm:"foreignKey:merchant_id;References:id"`
}

func (MerchantPeriod) TableName() string {
	return "merchant_periods"
}
