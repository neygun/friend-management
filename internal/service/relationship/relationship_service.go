package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/relationship"
	"github.com/neygun/friend-management/internal/repository/user"
)

// Service represents relationship service
type Service interface {
	CreateFriendConnection(ctx context.Context, friendConnReq FriendConnectionInput) (model.Relationship, error)
	GetFriendsList(ctx context.Context, getFriendsInput GetFriendsInput) ([]string, int, error)
	GetCommonFriends(ctx context.Context, getCommonFriendsInput GetCommonFriendsInput) ([]string, int, error)
}

type service struct {
	userRepo         user.Repository
	relationshipRepo relationship.Repository
}

// New instantiates a relationship service
func New(userRepo user.Repository, relationshipRepo relationship.Repository) Service {
	return service{
		userRepo:         userRepo,
		relationshipRepo: relationshipRepo,
	}
}
