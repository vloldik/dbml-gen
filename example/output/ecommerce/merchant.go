// Code generated from DBML. DO NOT EDIT
package ecommerce

import public "output/public"

// merchants table
type Merchant struct {
	// name chosen by configuration
	ID           int             `gorm:"column:id;primaryKey;not null;index:ix_merchant_id__country_code"`
	CountryCode  *int            `gorm:"column:country_code;index:ix_merchant_id__country_code"`
	MerchantName *string         `gorm:"column:merchant_name;size:255"`
	CreatedAt    *string         `gorm:"column:created_at"`
	AdminID      *int            `gorm:"column:admin_id"`
	Country      *public.Country `gorm:"foreignKey:Code;References:CountryCode"`
	User         *public.User    `gorm:"foreignKey:ID;References:AdminID"`
}

func (Merchant) TableName() string {
	return "merchants"
}
