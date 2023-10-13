package relationship

import (
	"context"
	"fmt"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// BlockExists
// sql query: SELECT * FROM relationship WHERE requestor_id=requestor.ID AND target_id=target.ID AND type="block"
func (r relationshipRepository) BlockExists(ctx context.Context, requestor model.User, target model.User) (bool, error) {
	// Check if requestor blocks target when creating blocking relationship
	exists, err := ormmodel.Relationships(
		qm.Where(fmt.Sprintf("%s = ? AND %s = ? AND %s = ?",
			ormmodel.RelationshipColumns.RequestorID,
			ormmodel.RelationshipColumns.TargetID,
			ormmodel.RelationshipColumns.Type), requestor.ID, target.ID, "block"),
	).Exists(ctx, r.db)

	return exists, err
}
