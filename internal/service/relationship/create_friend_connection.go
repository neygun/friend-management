package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

// FriendConnectionInput represents the input from request to create friend connection
type FriendConnectionInput struct {
	Friends []string
}

// CreateFriendConnection gets 2 users from input and create a friend connection between them
func (s service) CreateFriendConnection(ctx context.Context, friendConnInput FriendConnectionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetUsers(ctx, user.UserFilter{Emails: friendConnInput.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there is 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	// check if block exists
	user1Block, err := s.relationshipRepo.GetRelationship(ctx, users[0].ID, users[1].ID, model.RelationshipTypeBlock)
	if err != nil {
		return model.Relationship{}, err
	}

	user2Block, err := s.relationshipRepo.GetRelationship(ctx, users[1].ID, users[0].ID, model.RelationshipTypeBlock)
	if err != nil {
		return model.Relationship{}, err
	}

	if user1Block != (model.Relationship{}) || user2Block != (model.Relationship{}) {
		return model.Relationship{}, ErrBlockExists
	}

	// create friend connection
	friendConn, err := s.relationshipRepo.Save(ctx, users[0].ID, users[1].ID, model.RelationshipTypeFriend)
	if err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
