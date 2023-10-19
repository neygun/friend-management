package user

import "github.com/neygun/friend-management/internal/service/user"

// Handler represents user handler
type Handler struct {
	userService user.Service
}

// New instantiates a user handler
func New(userService user.Service) Handler {
	return Handler{
		userService: userService,
	}
}
