package relationship

import (
	service "github.com/neygun/friend-management/internal/service/user"
)

type RelationshipHandler struct {
	userService service.UserService
}
