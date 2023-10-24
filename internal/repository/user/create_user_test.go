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

func TestImpl_CreateUser(t *testing.T) {
	type args struct {
		givenUser   model.User
		expDBFailed bool
		expRs       model.User
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			expRs: model.User{
				Email: "test@example.com",
			},
		},
		"error: db failed": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			expDBFailed: true,
			expErr:      errors.New("ormmodel: unable to insert into users: sql: database is closed"),
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

				// When
				result, err := instance.CreateUser(ctx, tc.givenUser)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotZero(t, result.ID)
					require.NotZero(t, result.CreatedAt)
					require.NotZero(t, result.UpdatedAt)
					if !cmp.Equal(tc.expRs, result,
						cmpopts.IgnoreFields(model.User{}, "ID", "CreatedAt", "UpdatedAt")) {
						t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, result,
							cmp.Diff(tc.expRs, result, cmpopts.IgnoreFields(model.User{}, "ID", "CreatedAt", "UpdatedAt")))
						t.FailNow()
					}
				}
			})
		})
	}
}
