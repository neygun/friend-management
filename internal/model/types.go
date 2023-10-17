package model

// RelationshipType is the type of relationship between two users
type RelationshipType string

// List of relationship types
const (
	RelationshipTypeFriend    RelationshipType = "FRIEND"
	RelationshipTypeBlock     RelationshipType = "BLOCK"
	RelationshipTypeSubscribe RelationshipType = "SUBSCRIBE"
)

// ToString converts Relationship to string
func (r RelationshipType) ToString() string {
	return string(r)
}
