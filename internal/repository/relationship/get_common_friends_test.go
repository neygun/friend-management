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

func TestImpl_GetCommonFriends(t *testing.T) {
	type args struct {
		givenUser1ID int64
		givenUser2ID int64
		expDBFailed  bool
		expRs        []string
		expErr       error
	}

	tcs := map[string]args{
		"success": {
			givenUser1ID: 1,
			givenUser2ID: 2,
			expRs:        []string{"test3@example.com", "test4@example.com"},
		},
		"error: db failed": {
			givenUser1ID: 1,
			givenUser2ID: 2,
			expDBFailed:  true,
			expErr:       errors.New("bind failed to execute query: sql: database is closed"),
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
				test.LoadTestSQLFile(t, tx, "testdata/get_common_friends.sql")

				// When
				result, err := instance.GetCommonFriends(ctx, tc.givenUser1ID, tc.givenUser2ID)

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
