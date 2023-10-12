package relationship

import (
	"context"
	"fmt"
	"strconv"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// FriendConnectionExists
// sql query: SELECT * FROM relationship WHERE type="friend" AND ((requestor_id=user1.ID AND target_id=user2.id)
//
//	OR (requestor_id=user2.ID AND target_id=user1.id))
func (r relationshipRepository) FriendConnectionExists(ctx context.Context, user1 model.User, user2 model.User) (bool, error) {
	qms := []qm.QueryMod{
		ormmodel.RelationshipWhere.Type.EQ("friend"),
		qm.And(fmt.Sprintf("(%s = %s AND %s = %s) OR (%s = %s AND %s = %s)",
			ormmodel.RelationshipColumns.RequestorID, strconv.Itoa(int(user1.ID)), ormmodel.RelationshipColumns.TargetID, strconv.Itoa(int(user2.ID)),
			ormmodel.RelationshipColumns.RequestorID, strconv.Itoa(int(user2.ID)), ormmodel.RelationshipColumns.TargetID, strconv.Itoa(int(user1.ID)),
		)),
	}
	exists, err := ormmodel.Relationships(qms...).Exists(ctx, r.db)
	if err != nil {
		return false, err
	}

	return exists, nil

	// exists, err := ormmodel.Relationships(qm.Where(qm.Expr(
	// 	ormmodel.RelationshipWhere.RequestorID.EQ(user1.ID),
	// 	qm.And(ormmodel.RelationshipWhere.TargetID.EQ(user2.ID)),
	// 	qm.And(ormmodel.RelationshipWhere.Type.EQ("friend")),
	// ),
	// 	qm.Or2(qm.Expr(
	// 		ormmodel.RelationshipWhere.RequestorID.EQ(user2.ID),
	// 		qm.And(ormmodel.RelationshipWhere.TargetID.EQ(user1.ID)),
	// 		qm.And(ormmodel.RelationshipWhere.Type.EQ("friend")),
	// 	)),
	// )).Exists(ctx, db)
}
