package model

import "time"

// User maps 'user' table
type User struct {
	ID        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
