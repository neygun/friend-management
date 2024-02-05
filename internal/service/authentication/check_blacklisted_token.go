package authentication

import (
	"context"
)

func (s service) CheckBlacklistedToken(ctx context.Context, token string) (bool, error) {
	isBlacklisted, err := s.authRepo.CheckBlacklistedToken(ctx, token)
	if err != nil {
		return false, err
	}

	return isBlacklisted, nil
}
