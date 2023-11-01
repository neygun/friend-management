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

func TestImpl_GetByCriteria(t *testing.T) {
	type args struct {
		givenFilter model.RelationshipFilter
		isEmpty     bool
		expDBFailed bool
		expRs       []model.Relationship
		expErr      error
	}

	tcs := map[string]args{
		"success": {
			givenFilter: model.RelationshipFilter{
				RequestorID: 1,
				TargetID:    2,
			},
			expRs: []model.Relationship{
				{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
				{
					ID:          2,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
			},
		},
		"empty": {
			givenFilter: model.RelationshipFilter{
				RequestorID: 1,
				TargetID:    2,
			},
			isEmpty: true,
			expRs:   []model.Relationship{},
		},
		"error: db failed": {
			givenFilter: model.RelationshipFilter{
				RequestorID: 1,
				TargetID:    2,
			},
			expDBFailed: true,
			expErr:      errors.New("ormmodel: failed to assign all query results to Relationship slice: bind failed to execute query: sql: database is closed"),
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
				if !tc.isEmpty {
					test.LoadTestSQLFile(t, tx, "testdata/get_by_criteria.sql")
				}

				// When
				result, err := instance.GetByCriteria(ctx, tc.givenFilter)

				// Then
				if tc.expErr != nil {
					require.EqualError(t, err, tc.expErr.Error())
				} else {
					require.NoError(t, err)
					if !cmp.Equal(tc.expRs, result,
						cmpopts.IgnoreFields(model.Relationship{}, "CreatedAt", "UpdatedAt")) {
						t.Errorf("\n order mismatched. \n expected: %+v \n got: %+v \n diff: %+v", tc.expRs, result,
							cmp.Diff(tc.expRs, result, cmpopts.IgnoreFields(model.Relationship{}, "CreatedAt", "UpdatedAt")))
						t.FailNow()
					}
				}
			})
		})
	}
}
