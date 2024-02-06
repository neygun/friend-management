package handler

import (
	"net/http"

	"github.com/neygun/friend-management/internal/service/relationship"
	"github.com/neygun/friend-management/internal/service/user"
)

// ConvertError converts service errors to handler errors
func ConvertError(err error) error {
	switch err {
	case relationship.ErrUserNotFound,
		relationship.ErrFriendConnectionExists,
		relationship.ErrSubscriptionExists,
		relationship.ErrBlockExists,
		user.ErrUserNotFound,
		user.ErrUserExists,
		user.ErrWrongPassword:
		return HandlerError{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	default: // Unexpected error
		return err
	}
}
