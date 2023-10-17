package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetByFilter gets a relationship by filter
func (r repository) GetByFilter(ctx context.Context, filter Filter) ([]model.Relationship, error) {

	var qms []qm.QueryMod

	if filter.RequestorID != 0 {
		qms = append(qms, ormmodel.RelationshipWhere.RequestorID.EQ(filter.RequestorID))
	}

	if filter.TargetID != 0 {
		qms = append(qms, ormmodel.RelationshipWhere.TargetID.EQ(filter.TargetID))
	}

	if filter.Type != "" {
		qms = append(qms, ormmodel.RelationshipWhere.Type.EQ(filter.Type))
	}

	relationships, err := ormmodel.Relationships(qms...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	result := make([]model.Relationship, len(relationships))
	for i, r := range relationships {
		result[i] = model.Relationship{
			ID:          r.ID,
			RequestorID: r.RequestorID,
			TargetID:    r.TargetID,
			Type:        r.Type,
			CreatedAt:   r.CreatedAt,
			UpdatedAt:   r.UpdatedAt,
		}
	}

	return result, nil
}
