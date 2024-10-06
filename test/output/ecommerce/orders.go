package ecommerce

type Orders struct {
	Id        int    `gorm:"column:id;primaryKey"`
	UserId    int    `gorm:"column:user_id;not null"`
	Status    string `gorm:"column:status"`
	CreatedAt string `gorm:"column:created_at"`
}
