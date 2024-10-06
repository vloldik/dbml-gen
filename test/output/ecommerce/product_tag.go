package ecommerce

type ProductTag struct {
	Id       int        `gorm:"column:id;primaryKey"`
	Name     string     `gorm:"column:name"`
	Products []*Product `gorm:"foreignKey:id;References:id;many2many:product_tag_products"`
}

func (ProductTag) TableName() string {
	return "product_tags"
}
