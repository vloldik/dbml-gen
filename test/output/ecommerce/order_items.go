package ecommerce

type OrderItems struct {
	OrderId   int `gorm:"column:order_id"`
	ProductId int `gorm:"column:product_id"`
	Quantity  int `gorm:"column:quantity;default:1"`
	Orders    *Orders
	Products  *Products
}
