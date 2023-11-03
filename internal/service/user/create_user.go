package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateUser creates a user
func (s service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userRepo.Create(ctx, user)
}
