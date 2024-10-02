package models

type TestUser struct {
	Id      int    `gorm:"column:id;primaryKey" json:"id"`
	Test    string `gorm:"column:test" json:"test"`
	UserId  int    `gorm:"column:user_id" json:"user_id"`
	UserId1 int    `gorm:"column:user_id1" json:"user_id1"`
}
