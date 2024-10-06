package ecommerce

import public "guthub.com/vloldik/dbml-gen/test/output/public"

type Merchants struct {
	Id           int    `gorm:"column:id"`
	CountryCode  int    `gorm:"column:country_code"`
	MerchantName string `gorm:"column:merchant_name;size:255"`
	CreatedAt    string `gorm:"column:created_at"`
	AdminId      int    `gorm:"column:admin_id"`
	Countries    *public.Countries
	Users        *public.Users
}
