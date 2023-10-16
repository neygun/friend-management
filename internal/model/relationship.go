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

// RelationshipType is the type of relationship between two users
type RelationshipType string

// List of relationship types
const (
	RelationshipTypeFriend RelationshipType = "FRIEND"
	RelationshipTypeBlock  RelationshipType = "BLOCK"
)

// ToString converts Relationship to string
func (r RelationshipType) ToString() string {
	return string(r)
}
