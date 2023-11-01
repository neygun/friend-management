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
	type mockGetByCriteriaUserRepo struct {
		expCall bool
		input   model.UserFilter
		output  []model.User
		err     error
	}

	type mockGetByCriteriaRelationshipRepo struct {
		expCall bool
		input   model.RelationshipFilter
		output  []model.Relationship
		err     error
	}

	type mockUpdateRepo struct {
		expCall bool
		input   model.Relationship
		output  model.Relationship
		err     error
	}

	type mockCreateRepo struct {
		expCall bool
		input   model.Relationship
		output  model.Relationship
		err     error
	}

	type args struct {
		givenCreateSubscriptionInput      CreateSubscriptionInput
		mockGetByCriteriaUserRepo         mockGetByCriteriaUserRepo
		mockGetByCriteriaRelationshipRepo mockGetByCriteriaRelationshipRepo
		mockUpdateRepo                    mockUpdateRepo
		mockCreateRepo                    mockCreateRepo
		expRs                             model.Relationship
		expErr                            error
	}

	tcs := map[string]args{
		"err - user not found": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
		"err - subscription exists": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				output: []model.Relationship{
					{
						ID:          1,
						RequestorID: 1,
						TargetID:    2,
						Type:        model.RelationshipTypeSubscribe,
					},
				},
			},
			expErr: ErrSubscriptionExists,
		},
		"err - user.GetByCriteria": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
				expCall: true,
				input: model.UserFilter{
					Emails: []string{
						"test1@example.com",
						"test2@example.com",
					},
				},
				err: errors.New("user.GetByCriteria error"),
			},
			expErr: errors.New("user.GetByCriteria error"),
		},
		"err - relationship.GetByCriteria": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				err: errors.New("relationship.GetByCriteria error"),
			},
			expErr: errors.New("relationship.GetByCriteria error"),
		},
		"err - Update": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				output: []model.Relationship{
					{
						ID:          1,
						RequestorID: 1,
						TargetID:    2,
						Type:        model.RelationshipTypeBlock,
					},
				},
			},
			mockUpdateRepo: mockUpdateRepo{
				expCall: true,
				input: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
				err: errors.New("Update error"),
			},
			expErr: errors.New("Update error"),
		},
		"err - Create": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				output: []model.Relationship{
					{
						ID:          1,
						RequestorID: 1,
						TargetID:    2,
						Type:        model.RelationshipTypeFriend,
					},
				},
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
				err: errors.New("Create error"),
			},
			expErr: errors.New("Create error"),
		},
		"update success": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				output: []model.Relationship{
					{
						ID:          1,
						RequestorID: 1,
						TargetID:    2,
						Type:        model.RelationshipTypeBlock,
					},
				},
			},
			mockUpdateRepo: mockUpdateRepo{
				expCall: true,
				input: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
			},
			expRs: model.Relationship{
				ID:          1,
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeSubscribe,
			},
		},
		"create success": {
			givenCreateSubscriptionInput: CreateSubscriptionInput{
				Requestor: "test1@example.com",
				Target:    "test2@example.com",
			},
			mockGetByCriteriaUserRepo: mockGetByCriteriaUserRepo{
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
			mockGetByCriteriaRelationshipRepo: mockGetByCriteriaRelationshipRepo{
				expCall: true,
				input: model.RelationshipFilter{
					RequestorID: 1,
					TargetID:    2,
				},
				output: []model.Relationship{
					{
						ID:          1,
						RequestorID: 1,
						TargetID:    2,
						Type:        model.RelationshipTypeFriend,
					},
				},
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.Relationship{
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
				output: model.Relationship{
					ID:          1,
					RequestorID: 1,
					TargetID:    2,
					Type:        model.RelationshipTypeSubscribe,
				},
			},
			expRs: model.Relationship{
				ID:          1,
				RequestorID: 1,
				TargetID:    2,
				Type:        model.RelationshipTypeSubscribe,
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
			if tc.mockGetByCriteriaUserRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaUserRepo.input).Return(tc.mockGetByCriteriaUserRepo.output, tc.mockGetByCriteriaUserRepo.err),
				}
			}

			if tc.mockGetByCriteriaRelationshipRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("GetByCriteria", ctx, tc.mockGetByCriteriaRelationshipRepo.input).Return(tc.mockGetByCriteriaRelationshipRepo.output, tc.mockGetByCriteriaRelationshipRepo.err),
				)
			}

			if tc.mockUpdateRepo.expCall {
				mockRelationshipRepo.ExpectedCalls = append(mockRelationshipRepo.ExpectedCalls,
					mockRelationshipRepo.On("Update", ctx, tc.mockUpdateRepo.input).Return(tc.mockUpdateRepo.output, tc.mockUpdateRepo.err),
				)
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
