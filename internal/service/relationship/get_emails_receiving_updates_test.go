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

func TestService_GetEmailsReceivingUpdates(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockGetEmailsReceivingUpdatesRepo struct {
		expCall          bool
		senderID         int64
		mentionedUserIDs []int64
		output           []string
		err              error
	}

	type args struct {
		givenInput                        GetEmailsReceivingUpdatesInput
		mockGetByCriteriaRepos            []mockGetByCriteriaRepo
		mockGetEmailsReceivingUpdatesRepo mockGetEmailsReceivingUpdatesRepo
		expRs                             []string
		expErr                            error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByCriteriaRepos: []mockGetByCriteriaRepo{
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test1@example.com",
						},
					},
					output: []model.User{},
				},
				{
					expCall: false,
				},
			},
			expErr: ErrUserNotFound,
		},
		"err - 1st GetByCriteria": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByCriteriaRepos: []mockGetByCriteriaRepo{
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test1@example.com",
						},
					},
					err: errors.New("GetByCriteria error"),
				},
				{
					expCall: false,
				},
			},
			expErr: errors.New("GetByCriteria error"),
		},
		"err - 2nd GetByCriteria": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByCriteriaRepos: []mockGetByCriteriaRepo{
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test1@example.com",
						},
					},
					output: []model.User{
						{
							ID:    1,
							Email: "test1@example.com",
						},
					},
				},
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test2@example.com",
						},
					},
					err: errors.New("GetByCriteria error"),
				},
			},
			expErr: errors.New("GetByCriteria error"),
		},
		"err - GetEmailsReceivingUpdates": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByCriteriaRepos: []mockGetByCriteriaRepo{
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test1@example.com",
						},
					},
					output: []model.User{
						{
							ID:    1,
							Email: "test1@example.com",
						},
					},
				},
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test2@example.com",
						},
					},
					output: []model.User{
						{
							ID:    2,
							Email: "test2@example.com",
						},
					},
				},
			},
			mockGetEmailsReceivingUpdatesRepo: mockGetEmailsReceivingUpdatesRepo{
				expCall:          true,
				senderID:         1,
				mentionedUserIDs: []int64{2},
				err:              errors.New("GetEmailsReceivingUpdates error"),
			},
			expErr: errors.New("GetEmailsReceivingUpdates error"),
		},
		"success": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByCriteriaRepos: []mockGetByCriteriaRepo{
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test1@example.com",
						},
					},
					output: []model.User{
						{
							ID:    1,
							Email: "test1@example.com",
						},
					},
				},
				{
					expCall: true,
					input: model.UserFilter{
						Emails: []string{
							"test2@example.com",
						},
					},
					output: []model.User{
						{
							ID:    2,
							Email: "test2@example.com",
						},
					},
				},
			},
			mockGetEmailsReceivingUpdatesRepo: mockGetEmailsReceivingUpdatesRepo{
				expCall:          true,
				senderID:         1,
				mentionedUserIDs: []int64{2},
				output: []string{
					"test2@example.com",
					"test4@example.com",
					"test5@example.com",
				},
			},
			expRs: []string{
				"test2@example.com",
				"test4@example.com",
				"test5@example.com",
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()
			mockUserRepo := user.NewMockRepository(t)
			mockRelationshipRepo := relationship.NewMockRepository(t)

			if tc.mockGetByCriteriaRepos[0].expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRepos[0].input).Return(tc.mockGetByCriteriaRepos[0].output, tc.mockGetByCriteriaRepos[0].err),
				}
			}

			if tc.mockGetByCriteriaRepos[1].expCall {
				mockUserRepo.ExpectedCalls = append(mockUserRepo.ExpectedCalls,
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRepos[1].input).Return(tc.mockGetByCriteriaRepos[1].output, tc.mockGetByCriteriaRepos[1].err),
				)
			}

			if tc.mockGetEmailsReceivingUpdatesRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = []*mock.Call{
					mockRelationshipRepo.On("GetEmailsReceivingUpdates", ctx, tc.mockGetEmailsReceivingUpdatesRepo.senderID, tc.mockGetEmailsReceivingUpdatesRepo.mentionedUserIDs).
						Return(tc.mockGetEmailsReceivingUpdatesRepo.output, tc.mockGetEmailsReceivingUpdatesRepo.err),
				}
			}

			// When
			instance := New(mockUserRepo, mockRelationshipRepo)
			rs, err := instance.GetEmailsReceivingUpdates(ctx, tc.givenInput)

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
