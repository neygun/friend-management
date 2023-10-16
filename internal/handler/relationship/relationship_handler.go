package relationship

import (
	"github.com/neygun/friend-management/internal/service/relationship"
)

// Handler represents relationship handler
type Handler struct {
	relationshipService relationship.Service
}

// New instantiates a RelationshipHandler
func New(relationshipService relationship.Service) Handler {
	return Handler{
		relationshipService: relationshipService,
	}
}
