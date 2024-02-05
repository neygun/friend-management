package user

import (
	"context"
	"time"
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
	if user.ID == 0 {
		return ErrUserNotFound
	}

	// Add the token to the blacklist with a 24-hour expiry
	expiration := time.Hour * 24
	if err := s.authRepo.AddTokenToBlacklist(ctx, input.Token, expiration); err != nil {
		return err
	}
	return nil
}
