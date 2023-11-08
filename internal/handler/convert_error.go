package handler

import (
	"net/http"

	"github.com/neygun/friend-management/internal/service/relationship"
)

// ConvertError converts service errors to handler errors
func ConvertError(err error) error {
	switch err {
	case relationship.ErrUserNotFound,
		relationship.ErrFriendConnectionExists,
		relationship.ErrSubscriptionExists,
		relationship.ErrBlockExists:
		return HandlerError{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	default: // Unexpected error
		return err
	}
}
