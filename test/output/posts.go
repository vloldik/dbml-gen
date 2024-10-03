package models

type Posts struct {
	Id        int
	Title     string
	Body      string
	UserId    int
	Status    string
	CreatedAt time.Time
}
