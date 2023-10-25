package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/queries"
)

// GetCommonFriends returns the common friends list between user1 and user2
func (r repository) GetCommonFriends(ctx context.Context, user1ID, user2ID int64) ([]string, error) {
	type Result struct {
		Email string
	}
	var rs []Result
	err := queries.Raw(`
		SELECT `+ormmodel.UserColumns.Email+
		` FROM `+ormmodel.TableNames.Relationships+` r1 `+
		` INNER JOIN `+ormmodel.TableNames.Relationships+` r2 ON r1.`+ormmodel.RelationshipColumns.TargetID+`=r2.`+ormmodel.RelationshipColumns.TargetID+
		` INNER JOIN `+ormmodel.TableNames.Users+` ON `+ormmodel.TableNames.Users+`.`+ormmodel.UserColumns.ID+`=r1.`+ormmodel.RelationshipColumns.TargetID+
		` WHERE r1.`+ormmodel.RelationshipColumns.RequestorID+`=$1 AND r1.`+ormmodel.RelationshipColumns.Type+`=$3 
			AND r2.`+ormmodel.RelationshipColumns.RequestorID+`=$2 AND r2.`+ormmodel.RelationshipColumns.Type+`=$3 
		UNION
		SELECT `+ormmodel.UserColumns.Email+
		` FROM `+ormmodel.TableNames.Relationships+` r1 `+
		` INNER JOIN `+ormmodel.TableNames.Relationships+` r2 ON r1.`+ormmodel.RelationshipColumns.RequestorID+`=r2.`+ormmodel.RelationshipColumns.RequestorID+
		` INNER JOIN `+ormmodel.TableNames.Users+` ON `+ormmodel.TableNames.Users+`.`+ormmodel.UserColumns.ID+`=r1.`+ormmodel.RelationshipColumns.RequestorID+
		` WHERE r1.`+ormmodel.RelationshipColumns.TargetID+`=$1 AND r1.`+ormmodel.RelationshipColumns.Type+`=$3 
			AND r2.`+ormmodel.RelationshipColumns.TargetID+`=$2 AND r2.`+ormmodel.RelationshipColumns.Type+`=$3 
	`, user1ID, user2ID, model.RelationshipTypeFriend.ToString()).Bind(ctx, r.db, &rs)

	if err != nil {
		return nil, err
	}

	emails := make([]string, len(rs))
	for i, v := range rs {
		emails[i] = v.Email
	}

	return emails, nil
}
