package user

import (
	"context"
	"database/sql"

	ormmodel "github.com/neygun/friend-management/internal/repository/ormmodel"

	"github.com/neygun/friend-management/internal/model"
)

// GetByEmail gets a user by email
func (r repository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := ormmodel.Users(ormmodel.UserWhere.Email.EQ(email)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return model.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
