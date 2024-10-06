package ecommerce

type OrderItem struct {
	OrderId   int      `gorm:"column:order_id"`
	ProductId int      `gorm:"column:product_id"`
	Quantity  int      `gorm:"column:quantity;default:1"`
	Order     *Order   `gorm:"foreignKey:order_id;References:id"`
	Product   *Product `gorm:"foreignKey:product_id;References:id"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
