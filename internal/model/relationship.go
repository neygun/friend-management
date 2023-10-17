package model

import "time"

// Relationship maps 'relationship' table
type Relationship struct {
	ID          int64     `json:"id"`
	RequestorID int64     `json:"requestorId"`
	TargetID    int64     `json:"targetId"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
