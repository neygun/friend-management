package user

import (
	"context"

	ormmodel "github.com/neygun/friend-management/internal/repository/ormmodel"

	"github.com/neygun/friend-management/internal/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetByCriteria gets users by criteria
func (r repository) GetByCriteria(ctx context.Context, filter model.UserFilter) ([]model.User, error) {
	var qms []qm.QueryMod

	if filter.Emails != nil {
		qms = append(qms, ormmodel.UserWhere.Email.IN(filter.Emails))
	}

	users, err := ormmodel.Users(qms...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	result := make([]model.User, len(users))
	for i, u := range users {
		result[i] = model.User{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
	}

	return result, nil
}
