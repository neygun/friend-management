package relationship

import (
	"context"
	"errors"
	"testing"

	friendErr "github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/relationship"
	"github.com/neygun/friend-management/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_CreateFriendConnection(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockBlockExistsRepo struct {
		expCall bool
		input   []int64
		output  bool
		err     error
	}

	type mockCreateRepo struct {
		expCall bool
		input   model.Relationship
		output  model.Relationship
		err     error
	}

	type args struct {
		givenFriendConnInput  FriendConnectionInput
		mockGetByCriteriaRepo mockGetByCriteriaRepo
		mockBlockExistsRepo   mockBlockExistsRepo
		mockCreateRepo        mockCreateRepo
		expRs                 model.Relationship
		expErr                error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
				},
			},
			expErr: ErrUserNotFound,
		},
		"err - block exists": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
					{
						ID:    2,
						Email: "test2@example.com",
					},
				},
			},
			mockBlockExistsRepo: mockBlockExistsRepo{
				expCall: true,
				input:   []int64{1, 2},
				output:  true,
			},
			expErr: ErrBlockExists,
		},
		"err - friend connection exists": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
					{
						ID:    2,
						Email: "test2@example.com",
					},
				},
			},
			mockBlockExistsRepo: mockBlockExistsRepo{
				expCall: true,
				input:   []int64{1, 2},
				output:  false,
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
				err: friendErr.Wrap(&pgconn.PgError{
					Code: "23505",
				}, "ormmodel: unable to insert into relationships"),
			},
			expErr: ErrFriendConnectionExists,
		},
		"err - GetByCriteria": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				err: errors.New("GetByCriteria error"),
			},
			expErr: errors.New("GetByCriteria error"),
		},
		"err - BlockExists": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
					{
						ID:    2,
						Email: "test2@example.com",
					},
				},
			},
			mockBlockExistsRepo: mockBlockExistsRepo{
				expCall: true,
				input:   []int64{1, 2},
				err:     errors.New("BlockExists error"),
			},
			expErr: errors.New("BlockExists error"),
		},
		"err - Create": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
					{
						ID:    2,
						Email: "test2@example.com",
					},
				},
			},
			mockBlockExistsRepo: mockBlockExistsRepo{
				expCall: true,
				input:   []int64{1, 2},
				output:  false,
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
				err: errors.New("Create error"),
			},
			expErr: errors.New("Create error"),
		},
		"success": {
			givenFriendConnInput: FriendConnectionInput{
				Friends: []string{
					"test1@example.com",
					"test2@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test1@example.com",
					},
					{
						ID:    2,
						Email: "test2@example.com",
					},
				},
			},
			mockBlockExistsRepo: mockBlockExistsRepo{
				expCall: true,
				input:   []int64{1, 2},
				output:  false,
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeFriend,
				},
			},
			expRs: model.Relationship{
				ID:          1,
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeFriend,
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()
			mockUserRepo := user.NewMockRepository(t)
			mockRelationshipRepo := relationship.NewMockRepository(t)

			// When
			if tc.mockGetByCriteriaRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRepo.input).Return(tc.mockGetByCriteriaRepo.output, tc.mockGetByCriteriaRepo.err),
				}
			}

			if tc.mockBlockExistsRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("BlockExists", ctx, tc.mockBlockExistsRepo.input).Return(tc.mockBlockExistsRepo.output, tc.mockBlockExistsRepo.err),
				)
			}

			if tc.mockCreateRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("Create", ctx, tc.mockCreateRepo.input).Return(tc.mockCreateRepo.output, tc.mockCreateRepo.err),
				)
			}

			instance := New(mockUserRepo, mockRelationshipRepo)
			rs, err := instance.CreateFriendConnection(ctx, tc.givenFriendConnInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
			}
		})
	}
}
