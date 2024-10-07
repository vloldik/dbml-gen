// Code generated from DBML. DO NOT EDIT
package ecommerce

import public "output/public"

// merchants table
type Merchant struct {
	ID           int             `gorm:"column:id;primaryKey;index:ix_merchant_id__country_code"`
	CountryCode  int             `gorm:"column:country_code;index:ix_merchant_id__country_code"`
	MerchantName string          `gorm:"column:merchant_name;size:255"`
	CreatedAt    string          `gorm:"column:created_at"`
	AdminID      int             `gorm:"column:admin_id"`
	Country      *public.Country `gorm:"foreignKey:CountryCode;References:Code"`
	User         *public.User    `gorm:"foreignKey:AdminID;References:ID"`
}

func (Merchant) TableName() string {
	return "merchants"
}
