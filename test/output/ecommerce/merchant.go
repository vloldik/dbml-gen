package ecommerce

import public "output/public"

// merchants table
type Merchant struct {
	Id           int             `gorm:"column:id;primaryKey;index:ix_merchant_id__country_code"`
	CountryCode  int             `gorm:"column:country_code;index:ix_merchant_id__country_code"`
	MerchantName string          `gorm:"column:merchant_name;size:255"`
	CreatedAt    string          `gorm:"column:created_at"`
	AdminId      int             `gorm:"column:admin_id"`
	Country      *public.Country `gorm:"foreignKey:country_code;References:code"`
	User         *public.User    `gorm:"foreignKey:admin_id;References:id"`
}

func (Merchant) TableName() string {
	return "merchants"
}
