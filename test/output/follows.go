package models

type Follows struct {
	FollowingUserId int
	FollowedUserId  int
	CreatedAt       time.Time
}
