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

func TestController_GetFriendsList(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockGetFriendsListRepo struct {
		expCall bool
		input   int64
		output  []string
		err     error
	}

	type args struct {
		givenGetFriendsInput   GetFriendsInput
		mockGetByCriteriaRepo  mockGetByCriteriaRepo
		mockGetFriendsListRepo mockGetFriendsListRepo
		expFriendsList         []string
		expCount               int
		expErr                 error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenGetFriendsInput: GetFriendsInput{
				Email: "test@example.com",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test@example.com",
					},
				},
				output: []model.User{},
			},
			expErr: ErrUserNotFound,
		},
		"err - GetByCriteria": {
			givenGetFriendsInput: GetFriendsInput{
				Email: "test@example.com",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test@example.com",
					},
				},
				err: errors.New("GetByCriteria error"),
			},
			expErr: errors.New("GetByCriteria error"),
		},
		"err - GetFriendsList": {
			givenGetFriendsInput: GetFriendsInput{
				Email: "test@example.com",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test@example.com",
					},
				},
			},
			mockGetFriendsListRepo: mockGetFriendsListRepo{
				expCall: true,
				input:   1,
				err:     errors.New("GetFriendsList error"),
			},
			expErr: errors.New("GetFriendsList error"),
		},
		"success": {
			givenGetFriendsInput: GetFriendsInput{
				Email: "test@example.com",
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test@example.com",
					},
				},
				output: []model.User{
					{
						ID:    1,
						Email: "test@example.com",
					},
				},
			},
			mockGetFriendsListRepo: mockGetFriendsListRepo{
				expCall: true,
				input:   1,
				output: []string{
					"test2@example.com",
					"test3@example.com",
				},
			},
			expFriendsList: []string{
				"test2@example.com",
				"test3@example.com",
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

			if tc.mockGetFriendsListRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("GetFriendsList", ctx, tc.mockGetFriendsListRepo.input).Return(tc.mockGetFriendsListRepo.output, tc.mockGetFriendsListRepo.err),
				)
			}

			instance := New(mockUserRepo, mockRelationshipRepo)
			friendsList, count, err := instance.GetFriendsList(ctx, tc.givenGetFriendsInput)

			// Then
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expFriendsList, friendsList)
				require.Equal(t, tc.expCount, count)
			}
		})
	}
}
