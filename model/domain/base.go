package domain

import "time"

type BaseModel struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
