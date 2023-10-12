package user

import "github.com/neygun/friend-management/internal/service/user"

type UserHandler struct {
	userService user.UserService
}

func New(userService user.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
