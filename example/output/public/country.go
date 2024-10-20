// Code generated from DBML. DO NOT EDIT
package public

type Country struct {
	Code          *int    `gorm:"column:code;primaryKey"`
	Name          *string `gorm:"column:name"`
	ContinentName *string `gorm:"column:continent_name"`
}

func (Country) TableName() string {
	return "countries"
}
