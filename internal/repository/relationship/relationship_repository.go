package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/pkg/db"
	"github.com/sony/sonyflake"
)

// Repository represents relationship repository
type Repository interface {
	GetByCriteria(ctx context.Context, filter model.RelationshipFilter) ([]model.Relationship, error)
	IsExistBlock(ctx context.Context, userIDs []int64) (bool, error)
	GetFriendsList(ctx context.Context, id int64) ([]string, error)
	GetCommonFriends(ctx context.Context, user1ID, user2ID int64) ([]string, error)
	Create(ctx context.Context, relationship model.Relationship) (model.Relationship, error)
	Update(ctx context.Context, relationship model.Relationship) (model.Relationship, error)
}

type repository struct {
	db    db.ContextExecutor
	idsnf *sonyflake.Sonyflake
}

// New instantiates a relationship repository
func New(db db.ContextExecutor) Repository {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if flake == nil {
		panic("Couldn't generate sonyflake.NewSonyflake")
	}

	return repository{
		db:    db,
		idsnf: flake,
	}
}
