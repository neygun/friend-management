package user

import (
	"context"
	"errors"
	"testing"

	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateUser(t *testing.T) {
	type mockCreateRepo struct {
		expCall bool
		input   model.User
		output  model.User
		err     error
	}
	type args struct {
		givenUser      model.User
		mockCreateRepo mockCreateRepo
		expRs          model.User
		expErr         error
	}

	tcs := map[string]args{
		"err - Create": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.User{
					Email: "test@example.com",
				},
				err: errors.New("Create error"),
			},
			expErr: errors.New("Create error"),
		},
		"success": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.User{
					Email: "test@example.com",
				},
				output: model.User{
					ID:    1,
					Email: "test@example.com",
				},
			},
			expRs: model.User{
				ID:    1,
				Email: "test@example.com",
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			ctx := context.Background()
			mockUserRepo := user.NewMockRepository(t)

			if tc.mockCreateRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("Create", ctx, tc.mockCreateRepo.input).Return(tc.mockCreateRepo.output, tc.mockCreateRepo.err),
				}
			}

			// When
			instance := New(mockUserRepo)
			rs, err := instance.CreateUser(ctx, tc.givenUser)

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
