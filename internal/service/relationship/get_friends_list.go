package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// GetFriendsList returns the friends list of an email and count
func (s service) GetFriendsList(ctx context.Context, email string) ([]string, int, error) {
	// get user by email
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: []string{email}})
	if err != nil {
		return nil, 0, err
	}

	// check if there is a user with the email
	if len(users) == 0 {
		return nil, 0, ErrUserNotFound
	}

	friendsList, err := s.relationshipRepo.GetFriendsList(ctx, users[0].ID)
	if err != nil {
		return nil, 0, err
	}
	return friendsList, len(friendsList), nil
}
