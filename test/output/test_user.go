package models

type TestUser struct {
	Id      int    `gorm:"column:id;primaryKey"`
	Test    string `gorm:"column:test"`
	UserId  int    `gorm:"column:user_id"`
	UserId1 int    `gorm:"column:user_id1"`
	User    *User  `gorm:"gorm:\"foreignKey:user_id\"" json:"User"`  // One to one relationship
	User    *User  `gorm:"gorm:\"foreignKey:user_id1\"" json:"User"` // Many to one relationship
}
