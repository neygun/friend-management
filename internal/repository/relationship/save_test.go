package relationship

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

func TestImpl_Save(t *testing.T) {
	type args struct {
		givenRelationship model.Relationship
		expDBFailed       bool
		expRs             model.Relationship
		expErr            error
	}

	tcs := map[string]args{
		"success": {
			givenRelationship: model.Relationship{
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeFriend.ToString(),
			},
			expRs: model.Relationship{
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeFriend.ToString(),
			},
		},
		"error: db failed": {
			givenRelationship: model.Relationship{
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeFriend.ToString(),
			},
			expDBFailed: true,
			expErr:      errors.New("ormmodel: unable to upsert relationships: sql: database is closed"),
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
				test.LoadTestSQLFile(t, tx, "testdata/save.sql")

				// When
				result, err := instance.Save(ctx, tc.givenRelationship)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					require.NotZero(t, result.ID)
					require.NotZero(t, result.CreatedAt)
					require.NotZero(t, result.UpdatedAt)
					if !cmp.Equal(tc.expRs, result,
						cmpopts.IgnoreFields(model.Relationship{}, "ID", "CreatedAt", "UpdatedAt")) {
						t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, result,
							cmp.Diff(tc.expRs, result, cmpopts.IgnoreFields(model.Relationship{}, "ID", "CreatedAt", "UpdatedAt")))
						t.FailNow()
					}
				}
			})
		})
	}
}
