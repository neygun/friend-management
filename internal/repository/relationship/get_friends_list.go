package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/queries"
)

// GetFriendsList returns a list of emails having friend connection with the user with id
func (r repository) GetFriendsList(ctx context.Context, id int64) ([]string, error) {
	type Result struct {
		Email string
	}
	var rs []Result
	err := queries.Raw(`
		SELECT `+ormmodel.UserColumns.Email+
		` FROM `+ormmodel.TableNames.Relationships+
		` INNER JOIN `+ormmodel.TableNames.Users+` ON `+ormmodel.TableNames.Users+`.`+ormmodel.UserColumns.ID+`=`+ormmodel.RelationshipColumns.TargetID+
		` WHERE `+ormmodel.RelationshipColumns.RequestorID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$2
		UNION
		SELECT `+ormmodel.UserColumns.Email+
		` FROM `+ormmodel.TableNames.Relationships+
		` INNER JOIN `+ormmodel.TableNames.Users+` ON `+ormmodel.TableNames.Users+`.`+ormmodel.UserColumns.ID+`=`+ormmodel.RelationshipColumns.RequestorID+
		` WHERE `+ormmodel.RelationshipColumns.TargetID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$2
	`, id, model.RelationshipTypeFriend.ToString()).Bind(ctx, r.db, &rs)

	if err != nil {
		return nil, err
	}

	emails := make([]string, len(rs))
	for i, v := range rs {
		emails[i] = v.Email
	}

	return emails, nil
}
