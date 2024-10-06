package ecommerce

import "time"

type Products struct {
	Id         int       `gorm:"column:id;primaryKey"`
	Name       string    `gorm:"column:name"`
	MerchantId int       `gorm:"column:merchant_id;not null"`
	Price      int       `gorm:"column:price"`
	Status     int       `gorm:"column:status"`
	CreatedAt  time.Time `gorm:"column:created_at;default:now()"`
	Merchants  *Merchants
}
