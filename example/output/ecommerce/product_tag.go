// Code generated from DBML. DO NOT EDIT
package ecommerce

type ProductTag struct {
	ID       *int       `gorm:"column:id;primaryKey"`
	Name     *string    `gorm:"column:name"`
	Products []*Product `gorm:"foreignKey:ID;References:ID;constraint:OnDelete:CASCADE;many2many:product_tag_products"`
}

func (ProductTag) TableName() string {
	return "product_tags"
}
