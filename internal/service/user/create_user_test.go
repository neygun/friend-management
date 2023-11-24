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
	type mockHashPassword struct {
		expCall bool
		input   string
		output  string
		err     error
	}

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
	type args struct {
		givenUser          model.User
		mockHashPassword   mockHashPassword
		mockGetByEmailRepo mockGetByEmailRepo
		mockCreateRepo     mockCreateRepo
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
			ctx := context.Background()
			mockUserRepo := user.NewMockRepository(t)
			mockPasswordEncoder := NewMockPasswordEncoder(t)

			if tc.mockGetByEmailRepo.expCall {
				mockUserRepo.ExpectedCalls = []*mock.Call{
					mockUserRepo.On("GetByEmail", ctx, tc.mockGetByEmailRepo.input).Return(tc.mockGetByEmailRepo.output, tc.mockGetByEmailRepo.err),
				}
			}

			if tc.mockHashPassword.expCall {
				mockPasswordEncoder.ExpectedCalls = []*mock.Call{
					mockPasswordEncoder.On("HashPassword", tc.mockHashPassword.input).Return(tc.mockHashPassword.output, tc.mockHashPassword.err),
				}
			}

			if tc.mockCreateRepo.expCall {
				mockUserRepo.ExpectedCalls = append(mockUserRepo.ExpectedCalls,
					mockUserRepo.On("Create", ctx, tc.mockCreateRepo.input).Return(tc.mockCreateRepo.output, tc.mockCreateRepo.err),
				)
			}

			// When
			instance := New(mockUserRepo, mockPasswordEncoder)
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
