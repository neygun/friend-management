package user

import (
	"context"

	ormmodel "github.com/neygun/friend-management/internal/repository/ormmodel"

	"github.com/neygun/friend-management/internal/model"
	"github.com/volatiletech/sqlboiler/types"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserFilter struct {
	Emails []string
}

func (r userRepository) GetUsers(ctx context.Context, userFilter UserFilter) ([]model.User, error) {
	var qms []qm.QueryMod

	if userFilter.Emails != nil {
		qms = append(qms, qm.WhereIn("email in ?", types.Array(userFilter.Emails)))
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
