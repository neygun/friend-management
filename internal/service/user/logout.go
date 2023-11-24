package user

import (
	"context"
	"time"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/pkg/util"
)

// LogoutInput is the input from request to logout
type LogoutInput struct {
	Email string
	Token string
}

// Logout processes user logout
func (s service) Logout(ctx context.Context, input LogoutInput) error {
	// get user by email
	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return err
	}

	// check if the user exists
	if user == (model.User{}) {
		return ErrUserNotFound
	}

	// Add the token to the blacklist with a 24-hour expiry
	if err := addToBlacklist(input.Token, time.Hour*24); err != nil {
		return err
	}
	return nil
}

func addToBlacklist(token string, expiry time.Duration) error {
	ctx := context.Background()
	client := util.NewRedisClient()
	err := client.Set(ctx, token, "revoked", expiry).Err()
	if err != nil {
		return err
	}
	return nil
}
