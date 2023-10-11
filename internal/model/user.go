package model

import "time"

type User struct {
	ID        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
