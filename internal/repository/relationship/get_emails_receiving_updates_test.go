package relationship

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/neygun/friend-management/pkg/db"
	"github.com/neygun/friend-management/pkg/util/test"
	"github.com/stretchr/testify/require"
)

func TestImpl_GetEmailsReceivingUpdates(t *testing.T) {
	type args struct {
		givenSenderID         int64
		givenMentionedUserIDs []int64
		expDBFailed           bool
		expRs                 []string
		expErr                error
	}

	tcs := map[string]args{
		"success": {
			givenSenderID:         1,
			givenMentionedUserIDs: []int64{2, 3},
			expRs:                 []string{"test2@example.com", "test4@example.com", "test5@example.com"},
		},
		"error: db failed": {
			givenSenderID:         1,
			givenMentionedUserIDs: []int64{2, 3},
			expDBFailed:           true,
			expErr:                errors.New("bind failed to execute query: sql: database is closed"),
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			test.WithTxDB(t, func(tx db.ContextExecutor) {
				// Given
				instance := New(tx)
				if tc.expDBFailed {
					dbMock, _, err := sqlmock.New()
					require.NoError(t, err)
					dbMock.Close()
					instance = New(dbMock)
				}
				test.LoadTestSQLFile(t, tx, "testdata/get_emails_receiving_updates.sql")

				// When
				result, err := instance.GetEmailsReceivingUpdates(ctx, tc.givenSenderID, tc.givenMentionedUserIDs)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expRs, result)
				}
			})
		})
	}
}
