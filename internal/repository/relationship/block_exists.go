package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// BlockExists checks if block between 2 user exists
func (r repository) BlockExists(ctx context.Context, userIds []int64) (bool, error) {
	var qms []qm.QueryMod

	qms = append(qms, ormmodel.RelationshipWhere.RequestorID.IN(userIds))
	qms = append(qms, ormmodel.RelationshipWhere.TargetID.IN(userIds))
	qms = append(qms, ormmodel.RelationshipWhere.Type.EQ(model.RelationshipTypeBlock.ToString()))

	exists, err := ormmodel.Relationships(qms...).Exists(ctx, r.db)
	return exists, err
}
