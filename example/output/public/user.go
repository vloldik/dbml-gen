// Code generated from DBML. DO NOT EDIT
package public

import "time"

type User struct {
	ID          *int       `gorm:"column:id;primaryKey;autoIncrement"`
	FullName    *string    `gorm:"column:full_name"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:TIMESTAMP"`
	CountryCode *int       `gorm:"column:country_code"`
	Country     *Country   `gorm:"foreignKey:Code;References:CountryCode"`
}

func (User) TableName() string {
	return "users"
}
