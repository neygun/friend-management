package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateUser creates a user
func (s userService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userRepo.CreateUser(ctx, user)
}
