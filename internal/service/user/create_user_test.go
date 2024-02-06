package user

import (
	"context"
	"errors"
	"testing"

	"github.com/neygun/friend-management/internal/cache/authentication"
	"github.com/neygun/friend-management/internal/model"
	"github.com/neygun/friend-management/internal/repository/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_CreateUser(t *testing.T) {
	type mockGetByEmailRepo struct {
		expCall bool
		input   string
		output  model.User
		err     error
	}

	type mockCreateRepo struct {
		expCall bool
		input   model.User
		output  model.User
		err     error
	}

	type mockHashPassword struct {
		expCall bool
		input   string
		output  string
		err     error
	}

	type args struct {
		givenUser          model.User
		mockGetByEmailRepo mockGetByEmailRepo
		mockCreateRepo     mockCreateRepo
		mockHashPassword   mockHashPassword
		expRs              model.User
		expErr             error
	}

	tcs := map[string]args{
		"err - user exists": {
			givenUser: model.User{
				Email:    "test@example.com",
				Password: "abc",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test@example.com",
				output: model.User{
					ID:       1,
					Email:    "test@example.com",
					Password: "abctest",
				},
			},
			expErr: ErrUserExists,
		},
		"err - GetByEmail": {
			givenUser: model.User{
				Email:    "test@example.com",
				Password: "abc",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test@example.com",
				err:     errors.New("GetByEmail error"),
			},
			expErr: errors.New("GetByEmail error"),
		},
		"err - HashPassword": {
			givenUser: model.User{
				Email:    "test@example.com",
				Password: "abc",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test@example.com",
				output:  model.User{},
			},
			mockHashPassword: mockHashPassword{
				expCall: true,
				input:   "abc",
				err:     errors.New("HashPassword error"),
			},
			expErr: errors.New("HashPassword error"),
		},
		"err - Create": {
			givenUser: model.User{
				Email:    "test@example.com",
				Password: "abc",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test@example.com",
				output:  model.User{},
			},
			mockHashPassword: mockHashPassword{
				expCall: true,
				input:   "abc",
				output:  "abctest",
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.User{
					Email:    "test@example.com",
					Password: "abctest",
				},
				err: errors.New("Create error"),
			},
			expErr: errors.New("Create error"),
		},
		"success": {
			givenUser: model.User{
				Email:    "test@example.com",
				Password: "abc",
			},
			mockGetByEmailRepo: mockGetByEmailRepo{
				expCall: true,
				input:   "test@example.com",
				output:  model.User{},
			},
			mockHashPassword: mockHashPassword{
				expCall: true,
				input:   "abc",
				output:  "abctest",
			},
			mockCreateRepo: mockCreateRepo{
				expCall: true,
				input: model.User{
					Email:    "test@example.com",
					Password: "abctest",
				},
				output: model.User{
					ID:       1,
					Email:    "test@example.com",
					Password: "abctest",
				},
			},
			expRs: model.User{
				ID:       1,
				Email:    "test@example.com",
				Password: "abctest",
			},
		},
	}

	for scenario, tc := range tcs {
		t.Run(scenario, func(t *testing.T) {
			// Given
			hashPasswordWrapperFn = func(password string) (string, error) {
				if tc.mockHashPassword.input != password {
					require.Error(t, errors.New("unexpected hash password input"))
				}
				return tc.mockHashPassword.output, tc.mockHashPassword.err
			}
			defer func() {
				hashPasswordWrapperFn = hashPassword
			}()

			ctx := context.Background()
			mockUserRepo := user.NewMockRepository(t)

			if tc.mockGetByEmailRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByEmail", ctx, tc.mockGetByEmailRepo.input).Return(tc.mockGetByEmailRepo.output, tc.mockGetByEmailRepo.err),
				}
			}

			if tc.mockCreateRepo.expCall {
				mockUserRepo.ExpectedCalls = append(mockUserRepo.ExpectedCalls,
					mockUserRepo.On("Create", ctx, tc.mockCreateRepo.input).Return(tc.mockCreateRepo.output, tc.mockCreateRepo.err),
				)
			}

			mockAuthRepo := authentication.NewMockRepository(t)

			// When
			instance := New(mockUserRepo, mockAuthRepo)
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
