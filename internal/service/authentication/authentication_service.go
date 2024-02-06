package authentication

import (
	"context"

	"github.com/neygun/friend-management/internal/cache/authentication"
)

type Service interface {
	CheckBlacklistedToken(ctx context.Context, token string) (bool, error)
}

type service struct {
	authRepo authentication.Repository
}

func New(authRepo authentication.Repository) Service {
	return service{
		authRepo: authRepo,
	}
}
