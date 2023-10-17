package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
)

// GetRelationship gets a relationship by requestorId, targetId and type
func (r repository) GetRelationship(ctx context.Context, requestorId int64, targetId int64, relationshipType model.RelationshipType) (model.Relationship, error) {
	relationship, err := ormmodel.Relationships(
		ormmodel.RelationshipWhere.RequestorID.EQ(requestorId),
		ormmodel.RelationshipWhere.TargetID.EQ(targetId),
		ormmodel.RelationshipWhere.Type.EQ(relationshipType.ToString())).One(ctx, r.db)
	if err != nil {
		return model.Relationship{}, err
	}
	return model.Relationship{
		ID:          relationship.ID,
		RequestorID: relationship.RequestorID,
		TargetID:    relationship.TargetID,
		Type:        relationship.Type,
		CreatedAt:   relationship.CreatedAt,
		UpdatedAt:   relationship.UpdatedAt,
	}, nil
}
