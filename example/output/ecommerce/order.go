// Code generated from DBML. DO NOT EDIT
package ecommerce

import public "output/public"

type Order struct {
	ID     int    `gorm:"column:id;primaryKey"`
	UserID int    `gorm:"column:user_id"`
	Status string `gorm:"column:status"`
	// When order created
	CreatedAt string       `gorm:"column:created_at"`
	User      *public.User `gorm:"foreignKey:UserID;References:ID"`
}

func (Order) TableName() string {
	return "orders"
}
