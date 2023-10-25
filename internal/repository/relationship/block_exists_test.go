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

func TestImpl_BlockExists(t *testing.T) {
	type args struct {
		givenUserIDs []int64
		expDBFailed  bool
		expRs        bool
		expErr       error
	}

	tcs := map[string]args{
		"success": {
			givenUserIDs: []int64{1, 2},
			expRs:        true,
		},
		"error: db failed": {
			givenUserIDs: []int64{1, 2},
			expDBFailed:  true,
			expErr:       errors.New("ormmodel: failed to check if relationships exists: sql: database is closed"),
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
				test.LoadTestSQLFile(t, tx, "testdata/block_exists.sql")

				// When
				result, err := instance.BlockExists(ctx, tc.givenUserIDs)

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
