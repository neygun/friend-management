package relationship

import (
	"context"
	"database/sql"

	"github.com/neygun/friend-management/internal/model"
	"github.com/sony/sonyflake"
)

// Repository represents relationship repository
type Repository interface {
	FriendConnectionExists(ctx context.Context, user1 model.User, user2 model.User, relationshipType model.RelationshipType) (bool, error)

	BlockExists(ctx context.Context, requestor model.User, target model.User, relationshipType model.RelationshipType) (bool, error)

	CreateFriendConnection(ctx context.Context, user1 model.User, user2 model.User, relationshipType model.RelationshipType) (model.Relationship, error)

	Save(ctx context.Context, user1 model.User, user2 model.User, relationshipType model.RelationshipType) (model.Relationship, error)
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
