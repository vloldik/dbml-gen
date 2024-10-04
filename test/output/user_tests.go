package models

type UserTests struct {
	TestId    int        `gorm:"column:test_id"`
	UserId    int        `gorm:"column:user_id"`
	TestField string     `gorm:"column:testField"`
	TestUsers []TestUser `gorm:"gorm:\"foreignKey:test_id\"" json:"TestUsers"` // Many to many relationship
	Users     []User     `gorm:"gorm:\"foreignKey:user_id\"" json:"Users"`     // Many to many relationship
}
