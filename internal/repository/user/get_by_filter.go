package user

import (
	"context"

	ormmodel "github.com/neygun/friend-management/internal/repository/ormmodel"

	"github.com/neygun/friend-management/internal/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetByFilter gets users by filter
func (r repository) GetByFilter(ctx context.Context, filter Filter) ([]model.User, error) {
	var qms []qm.QueryMod

	if filter.Emails != nil {
		emails := make([]interface{}, len(filter.Emails))
		for i, v := range filter.Emails {
			emails[i] = v
		}
		qms = append(qms, qm.WhereIn("email IN ?", emails...))
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
