package user

import (
	"context"
	"database/sql"

	"github.com/neygun/friend-management/internal/model"
	"github.com/sony/sonyflake"
)

// Repository represents user repository
type Repository interface {
	GetByFilter(ctx context.Context, filter Filter) ([]model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type repository struct {
	db    *sql.DB
	idsnf *sonyflake.Sonyflake
}

// New instantiates a user repository
func New(db *sql.DB) Repository {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if flake == nil {
		panic("Couldn't generate sonyflake.NewSonyflake")
	}

	return repository{
		db:    db,
		idsnf: flake,
	}
}
