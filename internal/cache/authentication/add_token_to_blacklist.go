package authentication

import (
	"context"
	"time"
)

func (r repository) AddTokenToBlacklist(ctx context.Context, token string, expiration time.Duration) error {
	err := r.cache.Set(ctx, token, true, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
