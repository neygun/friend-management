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
	type mockGetByEmailRepo struct {
		expCall bool
		input   string
		output  model.User
		err     error
	}

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
		mockGetByEmailRepo                mockGetByEmailRepo
		mockGetByCriteriaRepo             mockGetByCriteriaRepo
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
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test1@example.com",
				output:  model.User{},
			},
			expErr: ErrUserNotFound,
		},
		"err - GetByEmail": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test1@example.com",
				err:     errors.New("GetByEmail error"),
			},
			expErr: errors.New("GetByEmail error"),
		},
		"err - GetByCriteria": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test1@example.com",
				output: model.User{
					ID:    1,
					Email: "test1@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test2@example.com",
					},
				},
				err: errors.New("GetByCriteria error"),
			},
			expErr: errors.New("GetByCriteria error"),
		},
		"err - GetEmailsReceivingUpdates": {
			givenInput: GetEmailsReceivingUpdatesInput{
				Sender: "test1@example.com",
				Text:   "test test2@example.com",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test1@example.com",
				output: model.User{
					ID:    1,
					Email: "test1@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
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
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test1@example.com",
				output: model.User{
					ID:    1,
					Email: "test1@example.com",
				},
			},
			mockGetByCriteriaRepo: mockGetByCriteriaRepo{
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

			if tc.mockGetByEmailRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByEmail", ctx, tc.mockGetByEmailRepo.input).Return(tc.mockGetByEmailRepo.output, tc.mockGetByEmailRepo.err),
				}
			}

			if tc.mockGetByCriteriaRepo.expCall {
				mockUserRepo.ExpectedCalls = append(mockUserRepo.ExpectedCalls,
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRepo.input).Return(tc.mockGetByCriteriaRepo.output, tc.mockGetByCriteriaRepo.err),
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
