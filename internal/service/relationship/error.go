package relationship

import "errors"

var (
	// ErrUserNotFound occurs when 1 or 2 users not found by emails
	ErrUserNotFound = errors.New("user not found")

	// ErrFriendConnectionExists occurs when there is a friend connection between 2 users
	ErrFriendConnectionExists = errors.New("friend connection exists")

	// ErrFriendConnectionExists occurs when there is a blocking relationship between 2 users
	ErrBlockExists = errors.New("blocking relationship exists")
)
