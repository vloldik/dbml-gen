package models

import "time"

type Posts struct {
	Id        int       `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title;default:text"`
	Body      string    `gorm:"column:body"` // Content of the post
	UserId    int       `gorm:"column:user_id"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Users1    *Users1   `gorm:"gorm:\"foreignKey:user_id\"" json:"Users1"` // Many to one relationship
}
