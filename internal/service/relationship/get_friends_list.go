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
func (s service) GetFriendsList(ctx context.Context, getFriendsInput GetFriendsInput) ([]string, int, error) {
	// get user by email
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{getFriendsInput.Email}})
	if err != nil {
		return nil, 0, err
	}

	// check if there is a user with the email
	if len(users) == 0 {
		return nil, 0, ErrUserNotFound
	}

	// Get friends list
	friendsList, err := s.relationshipRepo.GetFriendsList(ctx, users[0].ID)
	if err != nil {
		return nil, 0, err
	}
	return friendsList, len(friendsList), nil
}
