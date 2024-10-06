package public

type Countries struct {
	Code          int    `gorm:"column:code;primaryKey"`
	Name          string `gorm:"column:name"`
	ContinentName string `gorm:"column:continent_name"`
}
