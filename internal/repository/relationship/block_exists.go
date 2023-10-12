package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// BlockExists
// sql query: SELECT * FROM relationship WHERE requestor_id=requestor.ID AND target_id=target.ID AND type="block"
func (r relationshipRepository) BlockExists(ctx context.Context, requestor model.User, target model.User) (bool, error) {
	exists, err := ormmodel.Relationships(qm.Where("requestor_id=? and target_id=? and type=?", requestor.ID, target.ID, "block")).Exists(ctx, r.db)

	return exists, nil
}
