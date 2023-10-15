package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/relationship"
	"github.com/neygun/friend-management/internal/repository/user"
)

// RelationshipService represents relationship service
type RelationshipService interface {
	CreateFriendConnection(ctx context.Context, friendConnReq FriendConnectionInput) (model.Relationship, error)
}

type relationshipService struct {
	userRepo         user.UserRepository
	relationshipRepo relationship.RelationshipRepository
}

// New instantiates a RelationshipService
func New(userRepo user.UserRepository, relationshipRepo relationship.RelationshipRepository) RelationshipService {
	return relationshipService{
		userRepo:         userRepo,
		relationshipRepo: relationshipRepo,
	}
}
