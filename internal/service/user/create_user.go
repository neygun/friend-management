package user

import (
	"context"
	"time"

	"github.com/neygun/friend-management/internal/model"
)

// UserInput represents the input from request to create user
type UserInput struct {
	ID        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUser creates a user
func (s service) CreateUser(ctx context.Context, userInput UserInput) (model.User, error) {
	return s.userRepo.CreateUser(ctx, model.User{
		ID:        userInput.ID,
		Email:     userInput.Email,
		CreatedAt: userInput.CreatedAt,
		UpdatedAt: userInput.UpdatedAt,
	})
}
