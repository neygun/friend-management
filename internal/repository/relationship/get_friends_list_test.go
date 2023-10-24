package relationship

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/neygun/friend-management/internal/repository/relationship/testdata"
	"github.com/neygun/friend-management/pkg/db"
	"github.com/neygun/friend-management/pkg/util/test"
	"github.com/stretchr/testify/require"
)

func TestImpl_GetFriendsList(t *testing.T) {
	type args struct {
		givenID     int64
		expDBFailed bool
		expRs       []string
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			givenID: 1,
			expRs:   []string{"test2@example.com", "test3@example.com"},
		},
		"error: db failed": {
			givenID:     1,
			expDBFailed: true,
			expErr:      errors.New("bind failed to execute query: sql: database is closed"),
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
				testdata.LoadTestSQLFile(t, tx, "testdata/get_friends_list.sql")

				// When
				result, err := instance.GetFriendsList(ctx, tc.givenID)

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
