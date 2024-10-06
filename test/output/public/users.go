package public

import "time"

type Users struct {
	Id          int       `gorm:"column:id;primaryKey"`
	FullName    string    `gorm:"column:full_name"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CountryCode int       `gorm:"column:country_code"`
	Countries   *Countries
}
