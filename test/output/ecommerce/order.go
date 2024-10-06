// Code generated from DBML. DO NOT EDIT
package ecommerce

type Order struct {
	Id     int    `gorm:"column:id;primaryKey"`
	UserId int    `gorm:"column:user_id;not null"`
	Status string `gorm:"column:status"`
	// When order created
	CreatedAt string `gorm:"column:created_at"`
}

func (Order) TableName() string {
	return "orders"
}
