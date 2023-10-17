package relationship

import (
	"context"
	"database/sql"

	"github.com/neygun/friend-management/internal/model"
	"github.com/sony/sonyflake"
)

// Repository represents relationship repository
type Repository interface {
	Save(ctx context.Context, relationship model.Relationship) (model.Relationship, error)
	GetByFilter(ctx context.Context, filter Filter) ([]model.Relationship, error)
}

type repository struct {
	db    *sql.DB
	idsnf *sonyflake.Sonyflake
}

// New instantiates a relationship repository
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
