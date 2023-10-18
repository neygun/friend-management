package model

import "time"

// Relationship maps 'relationship' table
type Relationship struct {
	ID          int64
	RequestorID int64
	TargetID    int64
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
