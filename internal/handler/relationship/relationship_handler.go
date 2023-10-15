package relationship

import (
	"github.com/neygun/friend-management/internal/service/relationship"
)

// RelationshipHandler represents relationship handler
type RelationshipHandler struct {
	relationshipService relationship.RelationshipService
}

// New instantiates a RelationshipHandler
func New(relationshipService relationship.RelationshipService) RelationshipHandler {
	return RelationshipHandler{
		relationshipService: relationshipService,
	}
}
