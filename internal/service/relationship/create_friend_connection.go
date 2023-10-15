package relationship

import (
	"context"
	"errors"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

var (
	// ErrInvalidUsersLength occurs when the number of users returned from GetUsers is not 2
	ErrInvalidUsersLength = errors.New("the number of users must be 2")

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
func (s relationshipService) CreateFriendConnection(ctx context.Context, friendConnInput FriendConnectionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetUsers(ctx, user.UserFilter{Emails: friendConnInput.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if number of users = 2
	if len(users) != 2 {
		return model.Relationship{}, ErrInvalidUsersLength
	}

	// check if friend connection exists
	friendConnExists, err := s.relationshipRepo.FriendConnectionExists(ctx, users[0], users[1])
	if err != nil {
		return model.Relationship{}, err
	}
	if friendConnExists {
		return model.Relationship{}, ErrFriendConnectionExists
	}

	// check if block exists
	user1BlockExists, err := s.relationshipRepo.BlockExists(ctx, users[0], users[1])
	if err != nil {
		return model.Relationship{}, err
	}

	user2BlockExists, err := s.relationshipRepo.BlockExists(ctx, users[1], users[0])
	if err != nil {
		return model.Relationship{}, err
	}

	if user1BlockExists || user2BlockExists {
		return model.Relationship{}, ErrBlockExists
	}

	// create friend connection
	friendConn, err := s.relationshipRepo.CreateFriendConnection(ctx, users[0], users[1])
	if err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
