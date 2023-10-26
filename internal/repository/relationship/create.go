package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create creates a relationship
func (r repository) Create(ctx context.Context, relationship model.Relationship) (model.Relationship, error) {
	newID, err := r.idsnf.NextID()
	if err != nil {
		return model.Relationship{}, err
	}
	subscription := ormmodel.Relationship{
		ID:          int64(newID),
		RequestorID: relationship.RequestorID,
		TargetID:    relationship.TargetID,
		Type:        relationship.Type,
	}

	if err := subscription.Upsert(ctx, r.db, true, []string{ormmodel.RelationshipColumns.RequestorID, ormmodel.RelationshipColumns.TargetID,
		ormmodel.RelationshipColumns.Type}, boil.Infer(), boil.Infer()); err != nil {
		return model.Relationship{}, err
	}

	return model.Relationship{
		ID:          subscription.ID,
		RequestorID: subscription.RequestorID,
		TargetID:    subscription.TargetID,
		Type:        subscription.Type,
		CreatedAt:   subscription.CreatedAt,
		UpdatedAt:   subscription.UpdatedAt,
	}, nil
}
