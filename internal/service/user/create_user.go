package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
)

// CreateUser creates a user
func (s service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	// get user by email
	u, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return model.User{}, err
	}

	// check if the user exists
	if u.ID != 0 {
		return model.User{}, ErrUserExists
	}

	// hash password
	hashedPassword, err := hashPasswordWrapperFn(user.Password)
	if err != nil {
		return model.User{}, err
	}
	user.Password = hashedPassword

	return s.userRepo.Create(ctx, user)
}
