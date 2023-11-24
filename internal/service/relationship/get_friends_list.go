package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// GetFriendsInput is the input from request to retrieve friends list
type GetFriendsInput struct {
	Email string
}

// GetFriendsList returns the friends list of an email and count
func (s service) GetFriendsList(ctx context.Context, input GetFriendsInput) ([]string, int, error) {
	// get user by email
	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, 0, err
	}

	// check if the user exists
	if user == (model.User{}) {
		return nil, 0, ErrUserNotFound
	}

	// Get friends list
	friendsList, err := s.relationshipRepo.GetFriendsList(ctx, user.ID)
	if err != nil {
		return nil, 0, err
	}
	return friendsList, len(friendsList), nil
}
