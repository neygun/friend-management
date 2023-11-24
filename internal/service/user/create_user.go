package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateUser creates a user
func (s service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	// get user by email
	theUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return model.User{}, err
	}

	// check if the user exists
	if theUser != (model.User{}) {
		return model.User{}, ErrUserExists
	}

	// hash password
	hashedPassword, err := s.passwordEncoder.HashPassword(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = hashedPassword

	return s.userRepo.Create(ctx, user)
}
