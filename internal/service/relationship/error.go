package relationship

import "errors"

var (
	// ErrUserNotFound occurs when 1 or 2 users not found by emails
	ErrUserNotFound = errors.New("user not found")

	// ErrSubscriptionExists occurs when there is a subscription relationship between 2 users
	ErrSubscriptionExists = errors.New("subscription exists")

	// ErrFriendConnectionExists occurs when there is a blocking relationship between 2 users
	ErrBlockExists = errors.New("blocking relationship exists")
)
