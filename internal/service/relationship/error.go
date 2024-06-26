package relationship

import "errors"

var (
	// ErrUserNotFound occurs when 1 or 2 users not found by emails
	ErrUserNotFound = errors.New("user not found")

	// ErrFriendConnectionExists occurs when there is a friend connection between 2 users
	ErrFriendConnectionExists = errors.New("friend connection exists")

	// ErrSubscriptionExists occurs when there is a subscription relationship between 2 users
	ErrSubscriptionExists = errors.New("subscription exists")

	// ErrBlockExists occurs when there is a blocking relationship between 2 users
	ErrBlockExists = errors.New("blocking relationship exists")
)
