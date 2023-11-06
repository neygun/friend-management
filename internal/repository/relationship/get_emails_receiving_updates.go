package relationship

import (
	"context"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/ormmodel"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/types"
)

// GetEmailsReceivingUpdates returns emails that can receive updates from the sender
func (r repository) GetEmailsReceivingUpdates(ctx context.Context, senderID int64, mentionedUserIDs []int64) ([]string, error) {
	ids := types.Array(mentionedUserIDs)
	type Result struct {
		Email string
	}
	var rs []Result
	err := queries.Raw(`
		SELECT `+ormmodel.UserColumns.Email+`
		FROM (
			SELECT `+ormmodel.UserColumns.ID+`, `+ormmodel.UserColumns.Email+`
			FROM `+ormmodel.TableNames.Users+`
			WHERE `+ormmodel.UserColumns.ID+` IN (
				SELECT `+ormmodel.RelationshipColumns.TargetID+`
				FROM `+ormmodel.TableNames.Relationships+`
				WHERE `+ormmodel.RelationshipColumns.RequestorID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$2
				UNION
				SELECT `+ormmodel.RelationshipColumns.RequestorID+`
				FROM `+ormmodel.TableNames.Relationships+`
				WHERE `+ormmodel.RelationshipColumns.TargetID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$2
			) OR `+ormmodel.UserColumns.ID+` IN (
				SELECT `+ormmodel.RelationshipColumns.RequestorID+`
				FROM `+ormmodel.TableNames.Relationships+`
				WHERE `+ormmodel.RelationshipColumns.TargetID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$3
			) OR `+ormmodel.UserColumns.ID+` = ANY($5)
		)
		WHERE `+ormmodel.UserColumns.ID+` NOT IN (
			SELECT `+ormmodel.RelationshipColumns.RequestorID+`
			FROM `+ormmodel.TableNames.Relationships+`
			WHERE `+ormmodel.RelationshipColumns.TargetID+`=$1 AND `+ormmodel.RelationshipColumns.Type+`=$4
		) AND `+ormmodel.UserColumns.ID+`!=$1
	`, senderID, model.RelationshipTypeFriend.ToString(), model.RelationshipTypeSubscribe.ToString(), model.RelationshipTypeBlock.ToString(), ids).Bind(ctx, r.db, &rs)

	if err != nil {
		return nil, err
	}

	emails := make([]string, len(rs))
	for i, v := range rs {
		emails[i] = v.Email
	}

	return emails, nil
}
