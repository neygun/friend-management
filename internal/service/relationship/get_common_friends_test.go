package relationship

import (
	"context"
	"errors"
	"testing"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/relationship"
	"github.com/neygun/friend-management/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_GetCommonFriends(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockGetCommonFriendsRepo struct {
		expCall bool
		user1ID int64
		user2ID int64
		output  []string
		err     error
	}

	type args struct {
		givenInput               GetCommonFriendsInput
		mockGetByCriteriaRepo    mockGetByCriteriaRepo
		mockGetCommonFriendsRepo mockGetCommonFriendsRepo
		expCommonFriends         []string
		expCount                 int
		expErr                   error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenInput: GetCommonFriendsInput{
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
		"err - GetByCriteria": {
			givenInput: GetCommonFriendsInput{
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
		"err - GetCommonFriends": {
			givenInput: GetCommonFriendsInput{
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
			mockGetCommonFriendsRepo: mockGetCommonFriendsRepo{
				expCall: true,
				user1ID: 1,
				user2ID: 2,
				err:     errors.New("GetCommonFriends error"),
			},
			expErr: errors.New("GetCommonFriends error"),
		},
		"success": {
			givenInput: GetCommonFriendsInput{
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
			mockGetCommonFriendsRepo: mockGetCommonFriendsRepo{
				expCall: true,
				user1ID: 1,
				user2ID: 2,
				output: []string{
					"test3@example.com",
					"test4@example.com",
				},
			},
			expCommonFriends: []string{
				"test3@example.com",
				"test4@example.com",
			},
			expCount: 2,
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

			if tc.mockGetCommonFriendsRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = []*mock.Call{
					mockRelationshipRepo.On("GetCommonFriends", ctx, tc.mockGetCommonFriendsRepo.user1ID, tc.mockGetCommonFriendsRepo.user2ID).
						Return(tc.mockGetCommonFriendsRepo.output, tc.mockGetCommonFriendsRepo.err),
				}
			}

			instance := New(mockUserRepo, mockRelationshipRepo)
			commonFriends, count, err := instance.GetCommonFriends(ctx, tc.givenInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expCommonFriends, commonFriends)
				require.Equal(t, tc.expCount, count)
			}
		})
	}
}
