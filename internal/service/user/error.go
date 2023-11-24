package user

import "errors"

var (
	// ErrUserExists occurs when trying to create a user that already exists
	ErrUserExists = errors.New("user exists")

	// ErrUserNotFound occurs when retrieving a user that does not exist
	ErrUserNotFound = errors.New("user not found")

	// ErrWrongPassword occurs when users type wrong password
	ErrWrongPassword = errors.New("wrong password")
)
