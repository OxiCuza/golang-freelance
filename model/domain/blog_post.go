package domain

import "time"

type BlogPost struct {
	Id        string
	Title     string
	Content   string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
