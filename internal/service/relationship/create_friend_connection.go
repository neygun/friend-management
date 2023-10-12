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
	users, err := s.userRepo.GetUsers(ctx, user.UserFilter{Emails: friendConnInput.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	if len(users) != 2 {
		return model.Relationship{}, ErrInvalidUsersLength
	}

	exists, err := s.relationshipRepo.FriendConnectionExists(ctx, users[0], users[1])
	if err != nil {
		return model.Relationship{}, err
	}
	if exists {
		return model.Relationship{}, ErrFriendConnectionExists
	}

	if s.relationshipRepo.BlockExists() {
		return model.Relationship{}, ErrBlockExists
	}

	if friendConn, err := s.relationshipRepo.CreateFriendConnection(ctx, users[0], users[1]); err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
