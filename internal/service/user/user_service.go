package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
)

// Service represents user service
type Service interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, input LoginInput) (string, error)
	Logout(ctx context.Context, input LogoutInput) error
}

type service struct {
	userRepo user.Repository
}

// New instantiates a user service
func New(userRepo user.Repository) Service {
	return service{
		userRepo: userRepo,
	}
}
