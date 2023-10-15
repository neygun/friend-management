package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

// UserService represents user service
type UserService interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type userService struct {
	userRepo user.UserRepository
}

// New instantiates a UserService
func New(userRepo user.UserRepository) UserService {
	return userService{
		userRepo: userRepo,
	}
}
