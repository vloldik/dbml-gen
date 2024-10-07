// Code generated from DBML. DO NOT EDIT
package ecommerce

import public "output/public"

type Order struct {
	Id     int    `gorm:"column:id;primaryKey"`
	UserId int    `gorm:"column:user_id"`
	Status string `gorm:"column:status"`
	// When order created
	CreatedAt string       `gorm:"column:created_at"`
	User      *public.User `gorm:"foreignKey:user_id;References:id"`
}

func (Order) TableName() string {
	return "orders"
}
