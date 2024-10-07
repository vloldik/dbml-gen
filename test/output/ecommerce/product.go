// Code generated from DBML. DO NOT EDIT
package ecommerce

import "time"

type Product struct {
	ID         int       `gorm:"column:id;primaryKey;index:ux_product_id,unique"`
	Name       string    `gorm:"column:name;size:255"`
	MerchantID int       `gorm:"column:merchant_id;not null;index:product_status"`
	Price      int       `gorm:"column:price"`
	Status     int       `gorm:"column:status;index:product_status"`
	CreatedAt  time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Merchant   *Merchant `gorm:"foreignKey:MerchantID;References:ID"`
}

func (Product) TableName() string {
	return "products"
}
