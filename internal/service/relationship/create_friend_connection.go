package relationship

import (
	"context"
	"errors"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

var (
	ErrInvalidUsersLength     = errors.New("the number of users must be 2")
	ErrFriendConnectionExists = errors.New("friend connection exists")
	ErrBlockExists            = errors.New("blocking relationship exists")
)

type FriendConnectionInput struct {
	Friends []string
}

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
