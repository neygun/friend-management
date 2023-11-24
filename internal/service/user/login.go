package user

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/neygun/friend-management/internal/model"
)

// LoginInput is the input from request to login
type LoginInput struct {
	Email    string
	Password string
}

// Login processes user login
func (s service) Login(ctx context.Context, input LoginInput) (string, error) {
	// get user by email
	user, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	// check if the user exists
	if user == (model.User{}) {
		return "", ErrUserNotFound
	}

	// check password
	if !CheckPasswordHash(input.Password, user.Password) {
		return "", ErrWrongPassword
	}

	return initToken(user.ID)
}

func initToken(userID int64) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	claims := map[string]interface{}{"user_id": userID}
	jwtauth.SetExpiryIn(claims, time.Minute*5)

	_, tokenString, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
