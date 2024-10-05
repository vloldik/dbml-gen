package models

import "time"

type Users1 struct {
	Id        int                                            `gorm:"column:id"`
	Username  string                                         `gorm:"column:username"`
	Role      string                                         `gorm:"column:role"`
	CreatedAt time.Time                                      `gorm:"column:created_at"`
	Follows   `gorm:"gorm:\"foreignKey:id\"" json:"Follows"` // One to many relationship
	Follows   `gorm:"gorm:\"foreignKey:id\"" json:"Follows"` // One to many relationship
}
