package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/relationship"
	"github.com/neygun/friend-management/internal/repository/user"
)

// FriendConnectionInput represents the input from request to create friend connection
type FriendConnectionInput struct {
	Friends []string
}

// CreateFriendConnection gets 2 users from input and create a friend connection between them
func (s service) CreateFriendConnection(ctx context.Context, friendConnInput FriendConnectionInput) (model.Relationship, error) {
	// get users by emails
	users, err := s.userRepo.GetByFilter(ctx, user.Filter{Emails: friendConnInput.Friends})
	if err != nil {
		return model.Relationship{}, err
	}

	// check if there is 2 users with the emails
	if len(users) != 2 {
		return model.Relationship{}, ErrUserNotFound
	}

	// check if block exists
	user1Block, err := s.relationshipRepo.GetByFilter(ctx, relationship.Filter{
		RequestorID: users[0].ID,
		TargetID:    users[1].ID,
		Type:        model.RelationshipTypeBlock.ToString(),
	})
	if err != nil {
		return model.Relationship{}, err
	}

	user2Block, err := s.relationshipRepo.GetByFilter(ctx, relationship.Filter{
		RequestorID: users[1].ID,
		TargetID:    users[0].ID,
		Type:        model.RelationshipTypeBlock.ToString(),
	})
	if err != nil {
		return model.Relationship{}, err
	}

	if len(user1Block) != 0 || len(user2Block) != 0 {
		return model.Relationship{}, ErrBlockExists
	}

	// create friend connection
	friendConn := model.Relationship{
		RequestorID: users[0].ID,
		TargetID:    users[1].ID,
		Type:        model.RelationshipTypeFriend.ToString(),
	}

	friendConn, err = s.relationshipRepo.Save(ctx, friendConn)
	if err != nil {
		return model.Relationship{}, err
	}

	return friendConn, nil
}
