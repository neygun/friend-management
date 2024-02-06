package authentication

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache
type Repository interface {
	CheckBlacklistedToken(ctx context.Context, token string) (bool, error)
	AddTokenToBlacklist(ctx context.Context, token string, expiration time.Duration) error
}

type repository struct {
	cache *redis.Client
}

func New(cache *redis.Client) Repository {
	return repository{
		cache: cache,
	}
}
