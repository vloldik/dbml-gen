// Code generated from DBML. DO NOT EDIT
package ecommerce

type OrderItem struct {
	OrderID   *int     `gorm:"column:order_id"`
	ProductID *int     `gorm:"column:product_id"`
	Quantity  *int     `gorm:"column:quantity;default:1"`
	Order     *Order   `gorm:"foreignKey:ID;References:OrderID"`
	Product   *Product `gorm:"foreignKey:ID;References:ProductID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
