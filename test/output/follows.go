package models

import "time"

type Follows struct {
	FollowingUserId int       `gorm:"column:following_user_id;default:1"`
	FollowedUserId  int       `gorm:"column:followed_user_id"`
	CreatedAt       time.Time `gorm:"column:created_at;default:now() - interval '5 days'"`
}
