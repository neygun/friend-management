package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r repository) GetFriendsList(ctx context.Context, id int64) ([]string, error) {
	boil.DebugMode = true
	// err := ormmodel.NewQuery(
	// 	qm.Select("(CASE WHEN r."+ormmodel.RelationshipColumns.RequestorID+"="+strconv.FormatInt(id, 10)+
	// 		" THEN u_target."+ormmodel.UserColumns.Email+" ELSE u_requestor."+ormmodel.UserColumns.Email+" END) AS email"),
	// 	qm.From(ormmodel.TableNames.Relationship+" r "),
	// 	qm.InnerJoin("\"user\" AS u_requestor ON u_requestor."+ormmodel.UserColumns.ID+"=r."+ormmodel.RelationshipColumns.RequestorID),
	// 	qm.InnerJoin("\"user\" AS u_target ON u_target."+ormmodel.UserColumns.ID+"=r."+ormmodel.RelationshipColumns.TargetID),
	// 	ormmodel.RelationshipWhere.Type.EQ(model.RelationshipTypeFriend.ToString()),
	// 	qm.Where(strconv.FormatInt(id, 10)+" IN (r.requestor_id,r.target_id)"),
	// ).Bind(ctx, r.db, &qs)

	var users []ormmodel.User
	err := ormmodel.Users(
		qm.Select("\"user\".*"),
		qm.InnerJoin(ormmodel.TableNames.Relationship+" ON "+ormmodel.RelationshipColumns.Type+"=? AND ("+
			ormmodel.RelationshipColumns.RequestorID+"=? AND "+ormmodel.RelationshipColumns.TargetID+"=\"user\"."+ormmodel.UserColumns.ID+" OR "+
			ormmodel.RelationshipColumns.TargetID+"=? AND "+ormmodel.RelationshipColumns.RequestorID+"=\"user\"."+ormmodel.UserColumns.ID+")", model.RelationshipTypeFriend, id, id),
	).Bind(context.Background(), r.db, &users)
	if err != nil {
		return nil, err
	}

	var rs []string
	for i := range users {
		rs = append(rs, users[i].Email)
	}

	return rs, nil
}
