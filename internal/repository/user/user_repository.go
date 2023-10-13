package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neygun/friend-management/internal/model"
	"github.com/sony/sonyflake"
)

type UserRepository interface {
	GetUsers(ctx context.Context, userFilter UserFilter) ([]model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type userRepository struct {
	db    *sql.DB
	idsnf *sonyflake.Sonyflake
}

func New(db *sql.DB) UserRepository {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if flake == nil {
		fmt.Printf("Couldn't generate sonyflake.NewSonyflake. Doesn't work on Go Playground due to fake time.\n")
	}

	return userRepository{
		db:    db,
		idsnf: flake,
	}
}
