package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/pkg/util"
)

// FriendConnectionInput represents the input from request to create friend connection
type FriendConnectionInput struct {
	Friends []string
}

// CreateFriendConnection creates a friend connection between 2 users
func (s service) CreateFriendConnection(ctx context.Context, input FriendConnectionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: input.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there are 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	// check if block exists
	userIDs := []int64{users[0].ID, users[1].ID}
	blockExists, err := s.relationshipRepo.BlockExists(ctx, userIDs)
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
		Type:        model.RelationshipTypeFriend,
	}

	friendConn, err = s.relationshipRepo.Create(ctx, friendConn)
	if err != nil {
		if util.UniqueViolation.Is(err) {
			return model.Relationship{}, ErrFriendConnectionExists
		}
		return model.Relationship{}, err
	}

	return friendConn, nil
}
