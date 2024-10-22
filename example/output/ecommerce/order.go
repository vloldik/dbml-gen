// Code generated from DBML. DO NOT EDIT
package ecommerce

import public "output/public"

type Order struct {
	ID     *int    `gorm:"column:id;primaryKey"`
	UserID int     `gorm:"column:user_id;unique;not null"`
	Status *string `gorm:"column:status"`
	// When order created
	CreatedAt *string      `gorm:"column:created_at"`
	User      *public.User `gorm:"foreignKey:ID;References:UserID"`
}

func (Order) TableName() string {
	return "orders"
}
