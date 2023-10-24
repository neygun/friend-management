package relationship

import (
	"context"

	"github.com/volatiletech/sqlboiler/queries"
)

// GetFriendsList returns a list of emails having friend connection with the user with id
func (r repository) GetFriendsList(ctx context.Context, id int64) ([]string, error) {
	type Result struct {
		Email string
	}
	var rs []Result
	err := queries.Raw(`
		SELECT u.email
		FROM relationships r INNER JOIN users u ON u.id=r.target_id
		WHERE requestor_id=$1 AND "type"='FRIEND'
		UNION
		SELECT u.email
		FROM relationships r INNER JOIN users u ON u.id=r.requestor_id
		WHERE target_id=$2 AND "type"='FRIEND'
	`, id, id).Bind(ctx, r.db, &rs)

	if err != nil {
		return nil, err
	}

	var emails []string
	for _, v := range rs {
		emails = append(emails, v.Email)
	}

	return emails, nil
}
