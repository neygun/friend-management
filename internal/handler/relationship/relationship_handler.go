package relationship

import (
	"github.com/neygun/friend-management/internal/service/relationship"
)

type RelationshipHandler struct {
	relationshipService relationship.RelationshipService
}

func New(relationshipService relationship.RelationshipService) RelationshipHandler {
	return RelationshipHandler{
		relationshipService: relationshipService,
	}
}
