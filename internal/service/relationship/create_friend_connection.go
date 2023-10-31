package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// FriendConnectionInput represents the input from request to create friend connection
type FriendConnectionInput struct {
	Friends []string
}

// CreateFriendConnection gets 2 users from input and create a friend connection between them
func (s service) CreateFriendConnection(ctx context.Context, friendConnInput FriendConnectionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: friendConnInput.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there is 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	// check if block exists
	userIds := []int64{users[0].ID, users[1].ID}
	blockExists, err := s.relationshipRepo.BlockExists(ctx, userIds)
	if err != nil {
		return model.Relationship{}, err
	}
	if blockExists {
		return model.Relationship{}, ErrBlockExists
	}

	// create friend connection
	friendConn := model.Relationship{
		RequestorID: users[0].ID,
		TargetID:    users[1].ID,
		Type:        model.RelationshipTypeFriend.ToString(),
	}

	friendConn, err = s.relationshipRepo.Create(ctx, friendConn)
	if err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
