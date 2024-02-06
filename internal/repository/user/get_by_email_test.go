package user

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/pkg/db"
	"github.com/neygun/friend-management/pkg/util/test"
	"github.com/stretchr/testify/require"
)

func TestImpl_GetByEmail(t *testing.T) {
	type args struct {
		givenEmail  string
		isNotFound  bool
		expDBFailed bool
		expRs       model.User
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			givenEmail: "test@example.com",
			expRs: model.User{
				ID:       1,
				Email:    "test@example.com",
				Password: "123456",
			},
		},
		"not found": {
			givenEmail: "test@example.com",
			isNotFound: true,
			expRs:      model.User{},
		},
		"error: db failed": {
			givenEmail:  "test@example.com",
			expDBFailed: true,
			expErr:      errors.New("ormmodel: failed to execute a one query for users: bind failed to execute query: sql: database is closed"),
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
				if !tc.isNotFound {
					test.LoadTestSQLFile(t, tx, "testdata/get_by_email.sql")
				}

				// When
				result, err := instance.GetByEmail(ctx, tc.givenEmail)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					if !cmp.Equal(tc.expRs, result,
						cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt")) {
						t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, result,
							cmp.Diff(tc.expRs, result, cmpopts.IgnoreFields(model.User{}, "CreatedAt", "UpdatedAt")))
						t.FailNow()
					}
				}
			})
		})
	}
}
