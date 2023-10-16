package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreateFriendConnection creates a friend connection between 2 users in db
func (r repository) CreateFriendConnection(ctx context.Context, user1 model.User, user2 model.User, relationshipType model.RelationshipType) (model.Relationship, error) {
	newID, err := r.idsnf.NextID()
	if err != nil {
		return model.Relationship{}, err
	}
	friendConn := ormmodel.Relationship{
		ID:          int64(newID),
		RequestorID: user1.ID,
		TargetID:    user2.ID,
		Type:        relationshipType.ToString(),
	}

	if err := friendConn.Insert(ctx, r.db, boil.Infer()); err != nil {
		return model.Relationship{}, err
	}

	return model.Relationship{
		ID:          friendConn.ID,
		RequestorID: friendConn.RequestorID,
		TargetID:    friendConn.TargetID,
		Type:        friendConn.Type,
		CreatedAt:   friendConn.CreatedAt,
		UpdatedAt:   friendConn.UpdatedAt,
	}, nil
}
