package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Save upserts relationship
func (r repository) Save(ctx context.Context, relationship model.Relationship) (model.Relationship, error) {
	newID, err := r.idsnf.NextID()
	if err != nil {
		return model.Relationship{}, err
	}
	friendConn := ormmodel.Relationship{
		ID:          int64(newID),
		RequestorID: relationship.RequestorID,
		TargetID:    relationship.TargetID,
		Type:        relationship.Type,
	}

	if err := friendConn.Upsert(ctx, r.db, true, []string{ormmodel.RelationshipColumns.RequestorID, ormmodel.RelationshipColumns.TargetID,
		ormmodel.RelationshipColumns.Type}, boil.Whitelist(ormmodel.RelationshipColumns.Type), boil.Infer()); err != nil {
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
