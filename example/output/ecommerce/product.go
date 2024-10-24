// Code generated from DBML. DO NOT EDIT
package ecommerce

import "time"

type Product struct {
	ID         *int       `gorm:"column:id;primaryKey;index:ux_product_id,unique"`
	Name       *string    `gorm:"column:name;size:255"`
	MerchantID int        `gorm:"column:merchant_id;not null;index:product_status"`
	Price      *float64   `gorm:"column:price;type:DECIMAL(5,2)"`
	Status     *int       `gorm:"column:status;index:product_status"`
	CreatedAt  *time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	Merchant   *Merchant  `gorm:"foreignKey:ID;References:MerchantID;constraint:OnUpdate:CASCADE"`
}

func (Product) TableName() string {
	return "products"
}
