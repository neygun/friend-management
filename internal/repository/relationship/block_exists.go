package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r repository) BlockExists(ctx context.Context, userIds []int64) (bool, error) {
	var qms []qm.QueryMod

	qms = append(qms, ormmodel.RelationshipWhere.RequestorID.IN(userIds))
	qms = append(qms, ormmodel.RelationshipWhere.TargetID.IN(userIds))
	qms = append(qms, ormmodel.RelationshipWhere.Type.EQ(model.RelationshipTypeBlock.ToString()))

	blocks, err := ormmodel.Relationships(qms...).All(ctx, r.db)
	if err != nil {
		return false, err
	}

	return len(blocks) != 0, nil
}
