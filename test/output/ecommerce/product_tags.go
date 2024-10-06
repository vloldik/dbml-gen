package ecommerce

type ProductTags struct {
	Id       int    `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Products []*Products
}
