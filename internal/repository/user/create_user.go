package user

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	newID, err := r.idsnf.NextID()
	if err != nil {
		return model.User{}, err
	}
	u := ormmodel.User{
		ID:    int64(newID),
		Email: user.Email,
	}

	if err := u.Insert(ctx, r.db, boil.Infer()); err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
