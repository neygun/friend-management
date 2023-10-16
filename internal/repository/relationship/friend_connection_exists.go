package relationship

import (
	"context"
	"fmt"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// FriendConnectionExists checks if a friend connection between 2 users exists
// sql query: SELECT * FROM relationship WHERE type="friend" AND ((requestor_id=user1.ID AND target_id=user2.id)
//
//	OR (requestor_id=user2.ID AND target_id=user1.id))
func (r repository) FriendConnectionExists(ctx context.Context, user1 model.User, user2 model.User, relationshipType model.RelationshipType) (bool, error) {
	qms := []qm.QueryMod{
		ormmodel.RelationshipWhere.Type.EQ(relationshipType.ToString()),
		qm.And(fmt.Sprintf("(%s = ? AND %s = ?) OR (%s = ? AND %s = ?)",
			ormmodel.RelationshipColumns.RequestorID,
			ormmodel.RelationshipColumns.TargetID,
			ormmodel.RelationshipColumns.RequestorID,
			ormmodel.RelationshipColumns.TargetID,
		), user1.ID, user2.ID, user2.ID, user1.ID),
	}
	exists, err := ormmodel.Relationships(qms...).Exists(ctx, r.db)

	return exists, err
}
