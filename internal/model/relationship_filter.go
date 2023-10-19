package model

// RelationshipFilter defines filtering options for relationship repo methods
type RelationshipFilter struct {
	RequestorID int64
	TargetID    int64
	Type        RelationshipType
}
