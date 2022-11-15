package domain

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
