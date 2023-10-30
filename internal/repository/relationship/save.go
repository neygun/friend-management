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
	rel := ormmodel.Relationship{
		ID:          int64(newID),
		RequestorID: relationship.RequestorID,
		TargetID:    relationship.TargetID,
		Type:        relationship.Type,
	}

	if err := rel.Upsert(ctx, r.db, true, []string{ormmodel.RelationshipColumns.RequestorID, ormmodel.RelationshipColumns.TargetID,
		ormmodel.RelationshipColumns.Type}, boil.Infer(), boil.Infer()); err != nil {
		return model.Relationship{}, err
	}

	return model.Relationship{
		ID:          rel.ID,
		RequestorID: rel.RequestorID,
		TargetID:    rel.TargetID,
		Type:        rel.Type,
		CreatedAt:   rel.CreatedAt,
		UpdatedAt:   rel.UpdatedAt,
	}, nil
}
