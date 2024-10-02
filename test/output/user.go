package models

type User struct {
	Firstname string    `gorm:"column:firstname" json:"firstname"`
	Lastname  string    `gorm:"column:lastname" json:"lastname"`
	NameA     time.Time `gorm:"column:name_a" json:"name_a"`
	NameB     int       `gorm:"column:name_b" json:"name_b"`
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
}
