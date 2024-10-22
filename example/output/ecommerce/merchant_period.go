// Code generated from DBML. DO NOT EDIT
package ecommerce

import "time"

type MerchantPeriod struct {
	ID          *int       `gorm:"column:id;primaryKey"`
	MerchantID  *int       `gorm:"column:merchant_id"`
	CountryCode *int       `gorm:"column:country_code"`
	StartDate   *time.Time `gorm:"column:start_date;type:DATETIME"`
	EndDate     *time.Time `gorm:"column:end_date;type:DATETIME"`
	Merchant    *Merchant  `gorm:"foreignKey:ID;References:MerchantID"`
}

func (MerchantPeriod) TableName() string {
	return "merchant_periods"
}
