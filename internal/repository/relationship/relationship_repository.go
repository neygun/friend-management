package relationship

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/neygun/friend-management/internal/model"
	"github.com/sony/sonyflake"
)

type RelationshipRepository interface {
	FriendConnectionExists(ctx context.Context, user1 model.User, user2 model.User) (bool, error)
	BlockExists(ctx context.Context, requestor model.User, target model.User) (bool, error)
}

type relationshipRepository struct {
	db    *sql.DB
	idsnf *sonyflake.Sonyflake
}

func New(db *sql.DB) RelationshipRepository {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if flake == nil {
		fmt.Printf("Couldn't generate sonyflake.NewSonyflake. Doesn't work on Go Playground due to fake time.\n")
	}

	return relationshipRepository{
		db:    db,
		idsnf: flake,
	}
}