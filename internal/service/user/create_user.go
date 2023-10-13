package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

func (s userService) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userRepo.CreateUser(ctx, user)
}
