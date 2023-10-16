package relationship

import (
	"context"
	"errors"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

var (
	// ErrUserNotFound occurs when 1 or 2 users not found by emails
	ErrUserNotFound = errors.New("user not found")

	// ErrFriendConnectionExists occurs when there is a friend connection between 2 users
	ErrFriendConnectionExists = errors.New("friend connection exists")

	// ErrFriendConnectionExists occurs when there is a blocking relationship between 2 users
	ErrBlockExists = errors.New("blocking relationship exists")
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
	user1BlockExists, err := s.relationshipRepo.BlockExists(ctx, users[0], users[1], model.RelationshipTypeBlock)
	if err != nil {
		return model.Relationship{}, err
	}

	user2BlockExists, err := s.relationshipRepo.BlockExists(ctx, users[1], users[0], model.RelationshipTypeBlock)
	if err != nil {
		return model.Relationship{}, err
	}

	if user1BlockExists || user2BlockExists {
		return model.Relationship{}, ErrBlockExists
	}

	// create friend connection
	friendConn, err := s.relationshipRepo.Save(ctx, users[0], users[1], model.RelationshipTypeFriend)
	if err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
