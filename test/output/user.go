package models

import "time"

type User struct {
	Firstname string    `gorm:"column:firstname"`
	Lastname  string    `gorm:"column:lastname"`
	NameA     time.Time `gorm:"column:name_a"`
	NameB     int       `gorm:"column:name_b"`
	Id        int       `gorm:"column:id;primaryKey"`
}
