package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Update updates a relationship
func (r repository) Update(ctx context.Context, relationship model.Relationship) (model.Relationship, error) {
	rel, err := ormmodel.Relationships(ormmodel.RelationshipWhere.ID.EQ(relationship.ID)).One(ctx, r.db)
	if err != nil {
		return model.Relationship{}, err
	}

	rel.Type = relationship.Type.ToString()
	_, err = rel.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return model.Relationship{}, err
	}

	return model.Relationship{
		ID:          rel.ID,
		RequestorID: rel.RequestorID,
		TargetID:    rel.TargetID,
		Type:        model.RelationshipType(rel.Type),
		CreatedAt:   rel.CreatedAt,
		UpdatedAt:   rel.UpdatedAt,
	}, nil
}
