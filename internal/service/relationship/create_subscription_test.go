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

func TestController_CreateSubscription(t *testing.T) {
	type mockGetByCriteriaRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockCreateRepo struct {
		expCall bool
		input   model.Relationship
		output  model.Relationship
		err     error
	}

	type args struct {
		givenCreateSubscriptionInput CreateSubscriptionInput
		mockGetByCriteriaRepo        mockGetByCriteriaRepo
		mockCreateRepo               mockCreateRepo
		expRs                        model.Relationship
		expErr                       error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
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
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
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
		"err - Create": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
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
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe.ToString(),
				},
				err: errors.New("Save error"),
			},
			expErr: errors.New("Save error"),
		},
		"success": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
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
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe.ToString(),
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe.ToString(),
				},
			},
			expRs: model.Relationship{
				ID:          1,
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeSubscribe.ToString(),
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

			if tc.mockCreateRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("Create", ctx, tc.mockCreateRepo.input).Return(tc.mockCreateRepo.output, tc.mockCreateRepo.err),
				)
			}

			instance := New(mockUserRepo, mockRelationshipRepo)
			rs, err := instance.CreateSubscription(ctx, tc.givenCreateSubscriptionInput)

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
