// Code generated from DBML. DO NOT EDIT
package ecommerce

type OrderItem struct {
	OrderID   int      `gorm:"column:order_id"`
	ProductID int      `gorm:"column:product_id"`
	Quantity  int      `gorm:"column:quantity;default:1"`
	Order     *Order   `gorm:"foreignKey:OrderID;References:ID"`
	Product   *Product `gorm:"foreignKey:ProductID;References:ID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
