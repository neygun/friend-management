package relationship

import (
	"context"
	"strconv"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// SELECT (CASE WHEN r.requestor_id=$id THEN u_target.email ELSE u_requestor.email END) AS email FROM relationship r
// INNER JOIN "user" AS u_requestor ON u_requestor.id=r.requestor_id
// INNER JOIN "user" AS u_target ON u_target.id=r.target_id
// WHERE (r.type = 'FRIEND') AND ($id IN (r.requestor_id,r.target_id));
// GetFriendsList returns a list of emails having friend connection with the user with id
func (r repository) GetFriendsList(ctx context.Context, id int64) ([]string, error) {
	type Result struct {
		Email string
	}
	var rs []Result
	err := ormmodel.NewQuery(
		qm.Select("(CASE WHEN r."+ormmodel.RelationshipColumns.RequestorID+"="+strconv.FormatInt(id, 10)+
			" THEN u_target.email ELSE u_requestor.email END) AS email"),
		qm.From(ormmodel.TableNames.Relationship+" r "),
		qm.InnerJoin("\"user\" AS u_requestor ON u_requestor.id=r."+ormmodel.RelationshipColumns.RequestorID),
		qm.InnerJoin("\"user\" AS u_target ON u_target.id=r."+ormmodel.RelationshipColumns.TargetID),
		qm.Where("r.type=?", model.RelationshipTypeFriend.ToString()),
		qm.Where(strconv.FormatInt(id, 10)+" IN (r.requestor_id,r.target_id)"),
	).Bind(ctx, r.db, &rs)

	if err != nil {
		return nil, err
	}

	var emails []string
	for _, v := range rs {
		emails = append(emails, v.Email)
	}

	return emails, nil
}
