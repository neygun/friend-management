package handler

import (
	"net/http"

	"github.com/neygun/friend-management/internal/service/relationship"
)

// ConvertErr converts service errors to handler errors
func ConvertErr(err error) error {
	switch err {
	case relationship.ErrUserNotFound, relationship.ErrFriendConnectionExists, relationship.ErrSubscriptionExists, relationship.ErrBlockExists:
		return HandlerErr{
			Code:        http.StatusBadRequest,
			Description: err.Error(),
		}
	default: // Unexpected error
		return err
	}
}
