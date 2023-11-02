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
	type mockCreateUserRepo struct {
		expCall bool
		input   model.User
		output  model.User
		err     error
	}
	type args struct {
		givenUser          model.User
		mockCreateUserRepo mockCreateUserRepo
		expRs              model.User
		expErr             error
	}

	tcs := map[string]args{
		"err - CreateUser": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			mockCreateUserRepo: mockCreateUserRepo{
				expCall: true,
				input: model.User{
					Email: "test@example.com",
				},
				err: errors.New("CreateUser error"),
			},
			expErr: errors.New("CreateUser error"),
		},
		"success": {
			givenUser: model.User{
				Email: "test@example.com",
			},
			mockCreateUserRepo: mockCreateUserRepo{
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

			// When
			if tc.mockCreateUserRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("CreateUser", ctx, tc.mockCreateUserRepo.input).Return(tc.mockCreateUserRepo.output, tc.mockCreateUserRepo.err),
				}
			}

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
