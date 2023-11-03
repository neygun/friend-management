package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// GetCommonFriendsInput is the input from request to retrieve the common friends list between two email addresses
type GetCommonFriendsInput struct {
	Friends []string
}

// GetCommonFriends returns the common friends list between two emails and count
func (s service) GetCommonFriends(ctx context.Context, input GetCommonFriendsInput) ([]string, int, error) {
	// get users by emails
	users, err := s.userRepo.GetByCriteria(ctx, model.UserFilter{Emails: input.Friends})
	if err != nil {
		return nil, 0, err
	}

	// check if there are 2 users with the emails
	if len(users) != 2 {
		return nil, 0, ErrUserNotFound
	}

	// Get the common friends list
	commonFriends, err := s.relationshipRepo.GetCommonFriends(ctx, users[0].ID, users[1].ID)
	if err != nil {
		return nil, 0, err
	}
	return commonFriends, len(commonFriends), nil
}
