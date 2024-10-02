package models

type UserTests struct {
	TestId    int    `gorm:"column:test_id" json:"test_id"`
	UserId    int    `gorm:"column:user_id" json:"user_id"`
	TestField string `gorm:"column:testField" json:"testField"`
}
