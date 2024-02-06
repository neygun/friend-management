package authentication

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (r repository) CheckBlacklistedToken(ctx context.Context, token string) (bool, error) {
	_, err := r.cache.Get(ctx, token).Result()
	if err != nil {
		if err == redis.Nil {
			// Token not found in the blacklist
			return false, nil
		}
		return false, err
	}

	return true, nil
}
