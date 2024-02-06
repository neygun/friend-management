package user

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth/v5"
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
	if user.ID == 0 {
		return "", ErrUserNotFound
	}

	// check password
	matched, err := checkPasswordHash(input.Password, user.Password)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", ErrWrongPassword
	}

	return initToken(user.Email)
}

const algorithm = "HS256"
const secretKey = "secret"
const expiryTime = time.Minute * 5

func initToken(email string) (string, error) {
	tokenAuth := jwtauth.New(algorithm, []byte(secretKey), nil)

	claims := map[string]interface{}{"email": email}
	jwtauth.SetExpiryIn(claims, expiryTime)

	_, tokenString, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
